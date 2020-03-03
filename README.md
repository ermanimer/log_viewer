# go-log-viewer
Go log viewer

[![Build Status](https://travis-ci.org/ermanimer/log-viewer.svg?branch=master)](https://travis-ci.org/ermanimer/log-viewer)

## Features
go-log-viewer views log files which are created by [logger](https://github.com/ermanimer/logger).

## Installation
```bash
go get -v github.com/ermanimer/log-viewer
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
| Constant        | Value | Description       |
| :-------------- | :---- | :---------------- |
| defaultFilename | logs  | Default file name |
| defaultPrefixes | diwef | Default prefixes  |

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

#### Help
```bash
go-log-viewer --h
```

## Terminal Output
![Terminal Output](/images/terminal_output.png)

