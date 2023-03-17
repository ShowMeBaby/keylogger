package utils

import (
	"bytes"
	"errors"
	"keylogger/pkg/constants"
	"os"
	"os/exec"
	"os/user"
	"path"
	"runtime"
	"strings"
)

// GetLocalDataDir 获取本地数据存储目录
func GetLocalDataDir() string {
	var projectDir string
	if root := os.Getenv("KEYLOGGER_ROOT"); root != "" {
		projectDir = root
	} else {
		// Home目录
		homeDir, err := Home()
		if nil != err {
			panic("未获取到用户Home目录: " + err.Error())
		}
		projectDir = path.Join(homeDir, constants.AppLocalDataDir)
	}

	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		_ = os.MkdirAll(projectDir, os.ModePerm)
	}
	return projectDir
}

// Home 获取当前用户的Home目录
func Home() (string, error) {
	curUser, err := user.Current()
	if nil == err {
		return curUser.HomeDir, nil
	}

	// cross compile support
	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
