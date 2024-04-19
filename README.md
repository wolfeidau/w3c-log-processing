# w3c-log-processing

I am working on processing a years w3c logs and wanted to do some experiments before working on the real data, this project holds the results of this.

# Overview

This project breaks the problem into a couple of parts, and includes pprof flags to enable profiling. These parts are:

1. Processing of large files, in the case of my test dataset 3gb of logs
2. Parsing of w3c web logs
3. Parsing of user agent to enable classification of the client

# Usage

```
Usage: main <file> [flags]

Arguments:
  <file>    W3C log file to read.

Flags:
  -h, --help                  Show context-sensitive help.
      --cpu-profile=STRING    Enable CPU profiling.
```

# references

* https://www.kaggle.com/datasets/eliasdabbas/web-server-access-logs/data
