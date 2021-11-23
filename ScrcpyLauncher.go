package main

import (
	"github.com/Ericwyn/GoTools/shell"
	"github.com/Ericwyn/ScrcpyLauncher/conf"
	"github.com/Ericwyn/ScrcpyLauncher/ui"
)

func main() {
	shell.Debug(true)
	conf.InitConfig()

	ui.StartApp()
}
