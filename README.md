# log_viewer
Go log viewer

![Go](https://github.com/ermanimer/log_viewer/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ermanimer/log_viewer)](https://goreportcard.com/report/github.com/ermanimer/log_viewer)

## Features
go log_viewer views log files which are created with [logger](https://github.com/ermanimer/logger) on Unix based systems.

## Installation
```bash
go get -u github.com/ermanimer/log-viewer
```

## Prefixes
| Constant      | Value   | Description                 |
| :------------ | :------ | :-------------------------- |
| debugPrefix   | Debug   | Prefix for debug messages   |
| infoPrefix    | Info    | Prefix for info messages    |
| warningPrefix | Warning | Prefix for warning messages |
| errorPrefix   | Error   | Prefix for error messages   |
| fatalPrefix   | Fatal   | Prefix for fatal messages   |

## Default Filename and Prefixes
| Constant | Value       | Description       |
| :--------| :---------- | :---------------- |
| filename | default.log | Default file name |
| prefixes | diwef       | Default prefixes  |

## Usage
#### With default filename and prefixes
```bash
log-viewer
```

#### With filename and prefixes
```bash
log-viewer -f=filename -p=prefixes
```

#### Help
```bash
log-viewer -h
```

For prefixes, use combination of **d** for debug messages, **i** for information messages, **w** for warning messages, **e** for error messages, **f** for fatal messages.

## Terminal Output
![Terminal Output](/images/terminal_output.png)

