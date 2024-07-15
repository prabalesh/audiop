package audio

import (
	"audiop/internal/util"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func StartAudio(youtubeURL string, volume float64) error {
	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-g", youtubeURL)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error fetching audio stream URL: %w", err)
	}
	util.Colorize(util.ColorGreen, "Audio stream fetched successfully...\n")

	audioURL := strings.TrimSpace(string(output))
	volumeArg := fmt.Sprintf("volume=%f", volume)
	playCmd := exec.Command("nohup", "ffmpeg", "-i", audioURL, "-vn", "-af", volumeArg, "-f", "alsa", "default")
	playCmd.Stdout = os.Stdout
	playCmd.Stderr = os.Stderr
	playCmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := playCmd.Start(); err != nil {
		return fmt.Errorf("error starting ffmpeg: %w", err)
	}

	pid := playCmd.Process.Pid
	fmt.Printf("Started audio stream with PID %d\n", pid)

	// Write the PID, volume, and YouTube URL to files
	if err := os.WriteFile(util.PIDFilePath, []byte(fmt.Sprintf("%d", pid)), 0644); err != nil {
		return fmt.Errorf("error writing PID file: %w", err)
	}

	if err := os.WriteFile(util.VolumeFilePath, []byte(fmt.Sprintf("%f", volume)), 0644); err != nil {
		return fmt.Errorf("error writing volume file: %w", err)
	}

	if err := os.WriteFile(util.YoutubeURLPath, []byte(youtubeURL), 0644); err != nil {
		return fmt.Errorf("error writing YouTube URL file: %w", err)
	}

	return nil
}
