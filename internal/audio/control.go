package audio

import (
	"audiop/internal/util"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func ControlAudio(action string, volume float64) error {
	pidData, err := os.ReadFile(util.PIDFilePath)
	if err != nil {
		return fmt.Errorf("error reading PID file: %w", err)
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(pidData)))
	if err != nil {
		return fmt.Errorf("invalid PID: %w", err)
	}

	var signal syscall.Signal

	switch action {
	case "pause":
		signal = syscall.SIGSTOP
	case "resume":
		signal = syscall.SIGCONT
	case "stop":
		signal = syscall.SIGTERM
		if err := os.Remove(util.PIDFilePath); err != nil {
			return fmt.Errorf("error removing PID file: %w", err)
		}

		if err := os.Remove(util.VolumeFilePath); err != nil {
			return fmt.Errorf("error removing volume file: %w", err)
		}

		if err := os.Remove(util.YoutubeURLPath); err != nil {
			return fmt.Errorf("error removing YouTube URL file: %w", err)
		}
	case "volume":
		currentVolumeData, err := os.ReadFile(util.VolumeFilePath)
		if err != nil {
			return fmt.Errorf("error reading volume file: %w", err)
		}

		currentVolume, err := strconv.ParseFloat(strings.TrimSpace(string(currentVolumeData)), 64)
		if err != nil {
			return fmt.Errorf("invalid volume: %w", err)
		}

		if currentVolume == volume {
			return nil // No change in volume
		}

		if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
			return fmt.Errorf("error stopping ffmpeg: %w", err)
		}

		youtubeURLData, err := os.ReadFile(util.YoutubeURLPath)
		if err != nil {
			return fmt.Errorf("error reading YouTube URL file: %w", err)
		}

		youtubeURL := strings.TrimSpace(string(youtubeURLData))

		if err := StartAudio(youtubeURL, volume); err != nil {
			return fmt.Errorf("error starting ffmpeg with new volume: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("invalid action: %s", action)
	}

	if err := syscall.Kill(pid, signal); err != nil {
		return fmt.Errorf("error sending signal to process: %w", err)
	}

	fmt.Printf("%sed process with PID %d\n", strings.Title(action), pid)
	return nil
}
