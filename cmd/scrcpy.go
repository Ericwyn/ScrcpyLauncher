package cmd

import (
	"github.com/Ericwyn/GoTools/shell"
	"github.com/Ericwyn/ScrcpyLauncher/conf"
	"github.com/Ericwyn/ScrcpyLauncher/log"
	"github.com/spf13/viper"
	"strings"
)

func GetScrcpyVersion() string {
	res := shell.RunShellRes(viper.GetString(conf.ConfigScrcpyPath) + " -v")
	line1 := strings.Split(res, "\n")[0]
	line1 = strings.Trim(line1, " ")
	line1 = strings.Trim(line1, "\n")
	if line1 != "" {
		return strings.Replace(line1, "scrcpy ", "", -1)
	}

	return "UNKNOWN"
}

func GetStartScrcpyCommand(selectDeviceCmd string) string {
	command := viper.GetString(conf.ConfigScrcpyPath)
	if selectDeviceCmd != "" {
		command += " -s " + selectDeviceCmd
	}
	if viper.GetBool(conf.ConfigDisableScreenSaver) {
		command += " --disable-screensaver"
	}
	if viper.GetBool(conf.ConfigTurnScreenOff) {
		command += " --turn-screen-off"
	}
	if viper.GetBool(conf.ConfigStayAwake) {
		command += " --stay-awake"
	}
	if viper.GetBool(conf.ConfigHidKeyboard) {
		command += " --hid-keyboard"
	}
	if viper.GetBool(conf.ConfigShowTouches) {
		command += " --show-touches"
	}

	if viper.GetBool(conf.ConfigPowerOffOnClose) {
		command += " --power-off-on-close"
	}
	if viper.GetBool(conf.ConfigFullScreen) {
		command += " --fullscreen"
	}
	if viper.GetBool(conf.ConfigAlwaysOnTop) {
		command += " --always-on-top"
	}
	if viper.GetBool(conf.ConfigWindowBorderless) {
		command += " --window-borderless"
	}

	return command
}

func StartScrcpy(selectDeviceId string) {
	shell.RunOtherShell(GetStartScrcpyCommand(selectDeviceId), []string{}, func(resLine string) {
		log.D(strings.ReplaceAll(resLine, "\n", ""))
	})
}
