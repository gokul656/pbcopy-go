# pbcopy-go

A simple Go CLI tool to copy and paste text using the X11 clipboard on Linux/Unix systems.

> ⚠️ **Paused development**: The project is currently on hold because X11 clipboard requires the server process to remain alive at all times to serve clipboard data. This makes fully automated CLI usage challenging.

## Features

- Read from the X11 clipboard (`pbpaste` style)
- Write to the X11 clipboard (`pbcopy` style)
- Pure Go + CGO, no external dependencies like `xclip` required
