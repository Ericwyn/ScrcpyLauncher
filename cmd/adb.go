package cmd

import (
	"github.com/Ericwyn/GoTools/shell"
	"github.com/Ericwyn/ScrcpyLauncher/conf"
	"github.com/Ericwyn/ScrcpyLauncher/log"
	"github.com/spf13/viper"
	"strings"
)

type AdbDevices struct {
	Name   string
	Status string
}

func GetAdbVersion() string {
	res := shell.RunShellRes(viper.GetString(conf.ConfigAdbPath), "version")
	line1 := strings.Split(res, "\n")[0]
	line1 = strings.Trim(line1, " ")
	line1 = strings.Trim(line1, "\n")
	if line1 != "" {
		return strings.Replace(line1, "Android Debug Bridge version ", "", -1)
	}

	return "UNKNOWN"
}

func ListAdbDevices() []AdbDevices {
	res := make([]AdbDevices, 0)
	shellRes := shell.RunShellRes(viper.GetString(conf.ConfigAdbPath), "devices")

	//shellRes := shell.RunShellRes("/opt/platform-tools/adb", "devices")

	shellRes = strings.ReplaceAll(shellRes, "\r", "")
	log.D("adb devices now:")
	log.D(shellRes)
	for _, line := range strings.Split(shellRes, "\n") {
		if strings.HasPrefix(line, "List of devices") {
			continue
		}
		line = strings.Trim(line, " ")
		for i := 0; i < 3; i++ {
			line = strings.ReplaceAll(line, "  ", " ")
			line = strings.ReplaceAll(line, "\t\t", "\t")
		}
		lineSplit := strings.Split(line, "\t")
		if len(lineSplit) == 2 {
			res = append(res, AdbDevices{
				Name:   lineSplit[0],
				Status: lineSplit[1],
			})
		}
	}
	return res
}

func ListAdbDeviceId() []string {
	devices := ListAdbDevices()
	res := make([]string, 0)
	for _, device := range devices {
		res = append(res, device.Name)
	}
	return res
}
