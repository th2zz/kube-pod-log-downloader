# kube-pod-log-downloader
A simple Go program for downloading pod log from Kubernetes

## Quick Start

This is a simple Go program for downloading pod log from kubernetes by pod name regex.
It has no external dependencies so you can simplely `go run main.go` or `go build main.go` to get the binary.

## Usage

```
./<executable name> <pod_name_regex>	
```

Example command that fetches all pod log with pod name matching regex "api-1.0.0.*":


```
./get-pod-log "api-1.0.0.*"
```

