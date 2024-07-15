# audiop

audiop is a command-line tool that allows you to stream audio from a YouTube URL using `yt-dlp` and `ffmpeg`. It provides functionalities to start, pause, resume, stop the audio stream, adjust the volume, and get usage information.

## Installation

### Prerequisites

Before using audiop, make sure you have installed the following dependencies:

-   `yt-dlp`: A tool to download video and audio from YouTube.
-   `ffmpeg`: A multimedia framework for handling audio, video, and other multimedia files and streams.

#### Installing Dependencies

#### Ubuntu

```sh
sudo apt-get install yt-dlp ffmpeg
```

#### macOS

```sh
brew install yt-dlp ffmpeg
```

#### Installing yt-dlp Nightly

To update to the nightly version of yt-dlp from stable executable/binary:

```sh
yt-dlp --update-to nightly
```

#### Installing Nightly with pip

Alternatively, you can install the nightly version of yt-dlp with pip:

```sh
python3 -m pip install -U --pre "yt-dlp[default]"
```

### Build and Install audiop

#### Build and Install

To build and install audiop on your system, use the provided Makefile:

```sh
make install
```

#### Uninstall

To uninstall audiop, use:

```sh
make uninstall
```

#### Usage

To use audiop, run the following command with appropriate options:

```sh
audiop --start --url <YouTube URL>
```

Replace <YouTube URL> with the actual URL of the YouTube video or audio stream you want to play.

#### Command Line Options

    --start: Start streaming audio from the specified YouTube URL.
    --pause: Pause the currently playing audio stream.
    --resume: Resume a paused audio stream.
    --stop: Stop the currently playing audio stream.
    --vol <volume>: Adjust the volume of the audio stream (0.0 to 1.0).
    --help: Display usage information and command-line options.

### License

This project is licensed under the MIT License - see the LICENSE file for details.

### Explanation

This README.md file includes sections on installation instructions, how to build and install `audiop`, usage examples, command-line options, license information, and acknowledgments. It provides comprehensive guidance for users on installing dependencies, using the application, and understanding its capabilities.
