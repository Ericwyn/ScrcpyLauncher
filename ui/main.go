package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Ericwyn/ScrcpyLauncher/cmd"
	"github.com/Ericwyn/ScrcpyLauncher/conf"
	"github.com/Ericwyn/ScrcpyLauncher/log"
	"github.com/Ericwyn/ScrcpyLauncher/ui/resource/cusWidget"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var mainWindows fyne.Window

var devicesSelect *widget.Select
var selectDeviceId = "null"

var launcherBtn *widget.Button

func showMainUi() {

	mainWindows = mainApp.NewWindow("ScrcpyLauncher")
	mainWindows.Resize(fyne.Size{
		//Width: 500,
		//Height: 600,
	})
	mainWindows.CenterOnScreen()
	mainWindows.SetMainMenu(createAppMenu())

	sysInfoPanel := container.NewVBox(
		widget.NewLabel("adb version:"+cmd.GetAdbVersion()),
		widget.NewLabel("scrcpy version:"+cmd.GetScrcpyVersion()),
		widget.NewLabel("launcher version:"+conf.Version),
		widget.NewLabel("---------------------------------------------------------------"),
	)

	handleChecked := func(label string, checked bool) {
		if label == "禁止屏保" {
			viper.Set(conf.ConfigDisableScreenSaver, checked)
		} else if label == "禁止锁定" {
			viper.Set(conf.ConfigStayAwake, checked)
		} else if label == "熄灭屏幕" {
			viper.Set(conf.ConfigTurnScreenOff, checked)
		} else if label == "物理键盘模拟" {
			viper.Set(conf.ConfigHidKeyboard, checked)
		} else if label == "显示触摸" {
			viper.Set(conf.ConfigShowTouches, checked)
		} else if label == "退出时关闭设备屏幕" {
			viper.Set(conf.ConfigPowerOffOnClose, checked)
		} else if label == "窗口全屏" {
			viper.Set(conf.ConfigFullScreen, checked)
		} else if label == "窗口保持最前" {
			viper.Set(conf.ConfigAlwaysOnTop, checked)
		} else if label == "窗口无边框" {
			viper.Set(conf.ConfigWindowBorderless, checked)
		}

		conf.SaveConfig()
	}

	scrcpyConfigPanel := container.NewVBox(
		widget.NewLabel("参数配置    "),
		cusWidget.CreateCheckGroup(
			[]cusWidget.LabelAndInit{
				{"禁止电脑屏保", viper.GetBool(conf.ConfigDisableScreenSaver)},
				{"保持常亮", viper.GetBool(conf.ConfigStayAwake)},
				{"关闭设备屏幕", viper.GetBool(conf.ConfigTurnScreenOff)},
			},
			true,  // 横向
			false, // 单选
			func(label string, checked bool) {
				handleChecked(label, checked)
			},
		),
		cusWidget.CreateCheckGroup(
			[]cusWidget.LabelAndInit{
				{"物理键盘模拟", viper.GetBool(conf.ConfigHidKeyboard)},
				{"显示触摸", viper.GetBool(conf.ConfigShowTouches)},
				{"窗口全屏", viper.GetBool(conf.ConfigFullScreen)},
			},
			true,  // 横向
			false, // 单选
			func(label string, checked bool) {
				handleChecked(label, checked)
			},
		),
		cusWidget.CreateCheckGroup(
			[]cusWidget.LabelAndInit{
				{"退出时关闭设备屏幕", viper.GetBool(conf.ConfigPowerOffOnClose)},
				{"保持最前", viper.GetBool(conf.ConfigAlwaysOnTop)},
				{"无边框", viper.GetBool(conf.ConfigWindowBorderless)},
			},
			true,  // 横向
			false, // 单选
			func(label string, checked bool) {
				handleChecked(label, checked)
			},
		),
		widget.NewLabel("---------------------------------------------------------------"),
	)

	ids := cmd.ListAdbDeviceId()
	devicesSelect = widget.NewSelect(
		ids,
		func(s string) {
			device := strings.Split(s, ",")
			fmt.Println("选择启动 ", device[0])
			selectDeviceId = device[0]
		},
	)
	if len(ids) != 0 {
		devicesSelect.SetSelected(ids[0])
	}

	devicePanel := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("选择设备            "),
			widget.NewButton("刷新设备列表", func() {
				devicesSelect.ClearSelected()

				ids := cmd.ListAdbDeviceId()
				devicesSelect.Options = ids
				if len(ids) != 0 {
					devicesSelect.SetSelected(ids[0])

					if launcherBtn.Text != "启动" {
						launcherBtn.Text = "启动"
						launcherBtn.Enable()
						launcherBtn.Refresh()
					}

				}

				devicesSelect.Refresh()
			}),
		),
		devicesSelect,
		widget.NewLabel(""),
	)

	launcherBtn = widget.NewButton("启动", func() {
		if selectDeviceId == "null" {
			selectDeviceId = ""
		}
		mainWindows.SetOnClosed(func() {
			log.D("main windows close")
		})
		mainWindows.Close()
	})
	btnPanel := container.NewVBox(
		launcherBtn,
	)

	if len(ids) == 0 {
		launcherBtn.Text = "未找到 ADB 设备，请刷新"
		launcherBtn.Disable()
		launcherBtn.Refresh()
	}

	c := container.NewVBox(sysInfoPanel, scrcpyConfigPanel, devicePanel, btnPanel)

	mainWindows.SetContent(c)

	mainWindows.SetOnClosed(func() {
		log.D("main windows close")
		os.Exit(0)
	})

	mainWindows.ShowAndRun()
}
