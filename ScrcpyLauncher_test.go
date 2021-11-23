package main

import (
	"fmt"
	"github.com/Ericwyn/GoTools/shell"
	"github.com/Ericwyn/ScrcpyLauncher/cmd"
	"github.com/Ericwyn/ScrcpyLauncher/conf"
	"testing"
)

func TestCallScrcpy(t *testing.T) {

	conf.InitConfig()

	shell.RunShellCb(func(resLine string) {
		fmt.Println(resLine)
	}, "/usr/local/bin/scrcpy", "-v")
	//fmt.Println(res)
}

func TestPrintBinVersion(t *testing.T) {
	conf.InitConfig()

	fmt.Println("adb", cmd.GetAdbVersion())
	fmt.Println("scrcpy", cmd.GetScrcpyVersion())
}

func TestGetScrcpyVersion(t *testing.T) {
	//go shell.RunOtherShell("scrcpy -v", []string{}, func(resLine string) {
	//    fmt.Println(resLine)
	//})
	//time.Sleep(time.Second)
	//res := shell.RunShellRes("scrcpy", "-h")
	//fmt.Println(res)
	shell.RunOtherShell("scrcpy -h", []string{}, func(resLine string) {
		fmt.Println(resLine)
	})
}

func TestStartScrcpy(t *testing.T) {
	//shell.RunOtherShell(
	//    "scrcpy",
	//    []string{
	//        "-s",
	//        "10.239.150.141:39371",
	//        "--disable-screensaver",
	//        "--turn-screen-off",
	//        "--stay-awake",
	//    },
	//    func(resLine string) {
	//      fmt.Println(resLine)
	//    },
	//)

	shell.Debug(true)
	conf.InitConfig()

	cmd.StartScrcpy("10.239.150.141:39371")
}
