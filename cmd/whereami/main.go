package main

import (
	"github.com/s-yakubovskiy/whereami/cmd/cmd"
	"github.com/s-yakubovskiy/whereami/config"
)

var Version string

func init() {
	config.Init()
}

// TODO:
// 1.

func main() {
	cmd.Execute()
}
