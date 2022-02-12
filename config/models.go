package config

import (
	"fmt"
	"github.com/dhaanpaa-lab0/scr-webhookd/utils"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func GetTempFile() string {
	file, err := ioutil.TempFile("", "webhookd")
	if err != nil {
		return ""
	}

	return file.Name()
}

func GetUserHome() string {
	var homeDir, _ = os.UserHomeDir()
	return homeDir
}

func GetCurrentDir() string {
	getwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return getwd
}

func GetSystemRoot() string {
	absPath, err := filepath.Abs(viper.GetString("system_root"))
	if err != nil {
		return ""
	} else {
		return utils.NewDirIfNotExists(absPath)
	}

}

func GetSystemRootLogsPath() string {
	return utils.NewDirIfNotExists(path.Join(GetSystemRoot(), "logs"))
}

func GetSystemRootLogFile() string {
	return path.Join(GetSystemRootLogsPath(), "logfile.txt")
}

func GetSystemRootScriptsPath() string {
	return utils.NewDirIfNotExists(path.Join(GetSystemRoot(), "scripts"))
}

func GetListenAddress() string {
	return viper.GetString("listen_address")
}

func GetServerHeader() string {
	return viper.GetString("server_header")
}

func GetScripts() map[string]string {
	return viper.GetStringMapString("scripts")
}

func GetScriptFileName(key string) string {
	scripts := GetScripts()
	if scripts[key] == "" {
		return ""
	}
	scriptName := path.Join(GetSystemRootScriptsPath(), scripts[key])
	return scriptName

}

func ExecScript(scriptkey string, postedFileName string) string {
	execReturn := ""
	scriptFileName := GetScriptFileName(strings.ToLower(scriptkey))
	if scriptFileName == "" {
		return execReturn
	}

	if utils.FileExists(scriptFileName) {
		cmdExecScript := &exec.Cmd{
			Path: scriptFileName,
			Args: []string{
				scriptFileName,
				postedFileName,
			},
		}

		cmdOutput, errCmdExecScriptOutput := cmdExecScript.Output()
		if errCmdExecScriptOutput != nil {
			log.Fatal("ExecError:" + errCmdExecScriptOutput.Error())
			return ""
		} else {
			cmdOutputString := string(cmdOutput)
			log.Println(cmdOutputString)
			return cmdOutputString
		}

	} else {
		execReturn = "Script not found"
	}
	return execReturn
}

func init() {
	viper.AddConfigPath("/etc/webhookd")
	viper.AddConfigPath(path.Join(GetUserHome(), ".webhookd"))
	viper.AddConfigPath(GetCurrentDir())

	viper.SetConfigName("webhookd_config")
	viper.SetDefault("system_root", ".")
	viper.SetDefault("listen_address", "localhost:3002")
	viper.SetDefault("server_header", "Server")
	viper.SetDefault("scripts", map[string]string{})
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	logFile, errOpenFile := os.OpenFile(GetSystemRootLogFile(), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if errOpenFile != nil {
		panic(errOpenFile)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
