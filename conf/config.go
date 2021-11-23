package conf

import (
	"github.com/Ericwyn/ScrcpyLauncher/log"
	"github.com/Ericwyn/ScrcpyLauncher/ui/resource"
	"github.com/spf13/viper"
)

const Version = "V1.0"
const ReleaseDate = "2021.11.22"

// bin 配置
const ConfigAdbPath = "adb-path"
const ConfigScrcpyPath = "scrcpy-path"

// scrcpy 配置
// 关闭电脑屏保
const ConfigDisableScreenSaver = "disable-screensaver"

// 关闭设备屏幕
const ConfigTurnScreenOff = "turn-screen-off"

// 保持常亮
const ConfigStayAwake = "stay-awake"

// 物理键盘模拟, scrcpy 1.2 之后支持, 请参看官方文档
const ConfigHidKeyboard = "hid-keyboard"

// 显示触摸
const ConfigShowTouches = "show-touches"

// scrcpy 退出时关闭设备屏幕
const ConfigPowerOffOnClose = "power-off-on-close"

// 窗口全屏
const ConfigFullScreen = "fullscreen"

//// 旋转
//const ConfigRotation = "rotation"
// 窗口保持最前
const ConfigAlwaysOnTop = "always-on-top"

// 窗口无边框
const ConfigWindowBorderless = "window-borderless"

func InitConfig() {
	viper.SetDefault(ConfigAdbPath, "adb")
	viper.SetDefault(ConfigScrcpyPath, "scrcpy")

	viper.SetDefault(ConfigDisableScreenSaver, false)
	viper.SetDefault(ConfigTurnScreenOff, false)
	viper.SetDefault(ConfigStayAwake, false)
	viper.SetDefault(ConfigHidKeyboard, false)

	viper.SetDefault(ConfigShowTouches, false)
	viper.SetDefault(ConfigPowerOffOnClose, false)
	viper.SetDefault(ConfigFullScreen, false)
	viper.SetDefault(ConfigAlwaysOnTop, false)
	viper.SetDefault(ConfigWindowBorderless, false)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(resource.GetRunnerPath() + "/.conf")
	viper.AddConfigPath(resource.GetRunnerPath())
	err := viper.ReadInConfig()

	if err != nil {
		log.E("载入配置时候出错")
		panic(err)
	}
	printConfigs()
}

func printConfigs() {
	configList := []string{
		ConfigDisableScreenSaver,
		ConfigTurnScreenOff,
		ConfigStayAwake,
		ConfigAdbPath,
		ConfigScrcpyPath,
	}
	for _, key := range configList {
		log.D("config " + key + "  :  " + viper.GetString(key))
	}
}

func SaveConfig() {
	e := viper.WriteConfig()
	if e != nil {
		log.E("配置文件保存失败")
		log.E(e)
	}
}
