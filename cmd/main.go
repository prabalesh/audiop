package main

import (
	"audiop/internal/audio"
	"audiop/internal/util"
	"flag"
	"fmt"
)

func main() {
	flags := util.ParseFlags()

	if flags.Help {
		printUsage()
		return
	}

	if err := util.ValidateFlags(flags); err != nil {
		util.PrintError(err)
		fmt.Println("Usage: audiop --start|--pause|--resume|--stop|--vol [--url <YouTube URL>] [--vol <Volume>]")
		return
	}

	switch {
	case flags.Start:
		if err := audio.StartAudio(flags.URL, flags.Volume); err != nil {
			util.PrintError(err)
		}
	case flags.Pause:
		if err := audio.ControlAudio("pause", 0); err != nil {
			util.PrintError(err)
		}
	case flags.Resume:
		if err := audio.ControlAudio("resume", 0); err != nil {
			util.PrintError(err)
		}
	case flags.Stop:
		if err := audio.ControlAudio("stop", 0); err != nil {
			util.PrintError(err)
		}
	case flags.Volume >= 0 && flags.Volume <= 1:
		if err := audio.ControlAudio("volume", flags.Volume); err != nil {
			util.PrintError(err)
		}
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: audiop --start|--pause|--resume|--stop|--vol [--url <YouTube URL>] [--vol <Volume>]")
	flag.PrintDefaults()
}
