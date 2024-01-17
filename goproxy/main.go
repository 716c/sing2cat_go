package main

import (
	"goproxy/initial"
	"goproxy/merge"
)

func main() {
    initial.GetWorkingDir()
    initial.InitialLogger()
    initial.InitializeLoad()
    merge.GenerateConfigJson()
    defer initial.Logger.Sync()

}