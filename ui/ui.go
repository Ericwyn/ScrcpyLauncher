package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Ericwyn/ScrcpyLauncher/cmd"
	"github.com/Ericwyn/ScrcpyLauncher/ui/resource"
)

var mainApp fyne.App

func StartApp() {
	// 设置整个 app 的信息
	mainApp = app.New()
	mainApp.SetIcon(resource.ResourceIcon)
	mainApp.Settings().SetTheme(&resource.CustomerTheme{})

	showMainUi() // 此处有阻塞，除非把窗口关闭

	if selectDeviceId != "null" {
		cmd.StartScrcpy(selectDeviceId)
	}

}
