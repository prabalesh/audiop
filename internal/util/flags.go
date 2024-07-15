package util

import (
	"flag"
	"fmt"
	"os"
)

// Flags struct holds the command line arguments
type Flags struct {
	Start  bool
	Pause  bool
	Resume bool
	Stop   bool
	URL    string
	Volume float64
	Help   bool // Add Help field for --help flag
}

// ParseFlags parses command line arguments
func ParseFlags() Flags {
	var f Flags
	flag.BoolVar(&f.Start, "start", false, "Start the audio stream")
	flag.BoolVar(&f.Pause, "pause", false, "Pause the audio stream")
	flag.BoolVar(&f.Resume, "resume", false, "Resume the audio stream")
	flag.BoolVar(&f.Stop, "stop", false, "Stop the audio stream")
	flag.StringVar(&f.URL, "url", "", "YouTube URL for the audio stream")
	flag.Float64Var(&f.Volume, "vol", -1, "Volume level (0.0 to 1.0)")
	flag.BoolVar(&f.Help, "help", false, "Show usage information") // Add help flag
	flag.Parse()
	return f
}

// ValidateFlags validates the provided command line arguments
func ValidateFlags(flags Flags) error {
	if flags.Help {
		printUsage()
		os.Exit(0)
	}

	if flags.Start && flags.URL == "" {
		return fmt.Errorf("URL must be specified when starting the audio stream")
	}
	if flags.Volume < -1 || flags.Volume > 1 {
		return fmt.Errorf(`volume must be between 0.0 and 1.0`)
	}
	return nil
}

// printUsage prints usage information
func printUsage() {
	fmt.Println("Usage: audiop --start|--pause|--resume|--stop|--vol [--url <YouTube URL>] [--vol <Volume>]")
	flag.PrintDefaults()
}
