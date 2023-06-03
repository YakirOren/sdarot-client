# Sdarot Client

Sdarot Client is a command-line interface (CLI) application written in Go that allows you to easily download TV shows
from the Sdarot streaming platform. With Sdarot Client, you can conveniently access and save your favorite TV shows for
offline viewing.

## Features

- Simple and intuitive interface for interacting with the application.
- Ability to select specific episodes, seasons, or download the entire series.
- Efficient and fast downloading of TV shows from Sdarot.
- Seamless integration with the Sdarot platform for a seamless user experience.

## Installation

### Prebuilt Binaries

1. Go to the [Releases](https://github.com/YakirOren/sdarot-client/releases) page of the GitHub repository.
2. Download the appropriate prebuilt binary for your operating system (e.g., `sdarot-client-linux` for
   Linux, `sdarot-client-darwin` for macOS, or `sdarot-client-windows.exe` for Windows).
3. Make the binary executable (if necessary):
    - Linux/macOS: Run `chmod +x sdarot-client-linux` or `chmod +x sdarot-client-darwin`.
    - Windows: No additional steps needed.
4. Move the binary to a directory included in your system's PATH.
5. Verify the installation by running `sdarot-client` in the terminal. If installed successfully, you should see the
   application's help message.

### Build from Source

1. Clone the repository: `git clone https://github.com/YakirOren/sdarot-client.git`
2. Navigate to the project directory: `cd sdarot-client`
3. Build the application: `go build`

## Usage

1. Run the application: `sdarot-client`
2. Enter the series ID of the TV show you want to download.
3. Choose the desired download mode: specific episodes, specific seasons, or everything.
4. Follow the on-screen prompts to make your selections and initiate the download.
5. Sit back and let Sdarot Client handle the downloading process.

## Demo

![](demo.gif)