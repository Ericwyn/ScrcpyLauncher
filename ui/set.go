package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Ericwyn/ScrcpyLauncher/conf"
	"github.com/Ericwyn/ScrcpyLauncher/log"
	"github.com/spf13/viper"
)

var setWindow fyne.Window
var setWindowsOpening = false

func showSetUi() {
	if setWindowsOpening {
		setWindow.RequestFocus()
		return
	}

	setWindow = mainApp.NewWindow("程序设置")
	setWindow.Resize(fyne.Size{
		Width: 500,
		//Height: 600,
	})
	setWindow.CenterOnScreen()

	adbPathEntry := widget.NewEntry()
	adbPathEntry.SetPlaceHolder("adb 路径")
	adbPathEntry.SetText(viper.GetString(conf.ConfigAdbPath))

	scrcpyPathEntry := widget.NewEntry()
	scrcpyPathEntry.SetPlaceHolder("Scrcpy 路径")
	scrcpyPathEntry.SetText(viper.GetString(conf.ConfigScrcpyPath))

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "adb    路径", Widget: adbPathEntry, HintText: "请填写 adb 的程序路径"},
			{Text: "scrcpy 路径", Widget: scrcpyPathEntry, HintText: "请填写 scrcpy 的程序路径"},
			{Text: "", Widget: widget.NewLabel(""), HintText: ""},
		},

		SubmitText: "保存设置",
		OnSubmit: func() {
			log.D("保存设置")
			viper.Set(conf.ConfigAdbPath, adbPathEntry.Text)
			viper.Set(conf.ConfigScrcpyPath, scrcpyPathEntry.Text)

			conf.SaveConfig()
			// 点击保存之后返回
			setWindowsOpening = false
			setWindow.Close()
		},
	}

	setWindow.SetContent(form)

	setWindow.SetOnClosed(func() {
		setWindowsOpening = false
	})

	setWindowsOpening = true
	setWindow.Show()
}
