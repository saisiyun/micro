package main

import (
    "github.com/micro/micro/plugin"
    "github.com/micro/go-plugins/micro/gzip"
)

func init() {
    plugin.Register(gzip.New())
}