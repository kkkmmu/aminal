Raft is a terminal emulator utilising OpenGL v4.1.

The project is purely a learning exercise right now.

Ensure you have your latest graphics card drivers installed before use.

## Build Dependencies

- Go 1.10.3+
- On macOS, you need Xcode or Command Line Tools for Xcode (`xcode-select --install`) for required headers and libraries.
- On Ubuntu/Debian-like Linux distributions, you need `libgl1-mesa-dev xorg-dev`.
- On CentOS/Fedora-like Linux distributions, you need `libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel`.

## Planned Features

| Feature           | Done | Notes |
|-------------------|------|-------|
| Pty allocation    | ✔    | Linux only so far
| OpenGL rendering  | ✔    |
| Resizing/content reordering | ✔ | 
| ANSI escape codes | 50%  |
| UTF-8 input       | 90%  | No copy + paste as yet 
| UTF-8 output      | ✘    |
| Copy/paste        | ✘    |
| Customisable colour schemes | ✔ | Complete, but the config file has no entry for this yet 
| Config file       | 5%   |
| Scrolling         | 50%  | Infinite buffer implemented, need GUI scrollbar & render updates
| Sweet render effects | ✘ | 
||||

## Platform Support

| Platform | Supported
|----------|------------
| Linux    | ✔
| MacOSX   | ✘ (nearly)
| Windows  | ✘

## Configuration

Raft looks for a config file in the following places: `~/.raft.yml`, `~/.raft/config.yml`, `~/.config/raft/config.yml` (earlier in the list prioritised).

Example config:
```
debug: False
```

The following options are available:
| Name          | Type    | Description
| debug         | bool    | Enables debug logging
