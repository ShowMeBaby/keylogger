package constants

import (
	"time"
)

var (
	AppVersion   = "v0.0.1"
	LastfmKey    = ""
	LastfmSecret = ""
)

const AppName = "按键统计器"
const GroupID = "io.retrocode.keylogger"
const AppDescription = "<cyan>Keylogger - 每天按键统计器</>"
const AppGithubUrl = "https://github.com/ShowMeBaby/keylogger"
const AppLatestReleases = "https://github.com/ShowMeBaby/keylogger/releases/latest"
const AppCheckUpdateUrl = "https://api.github.com/repos/ShowMeBaby/keylogger/releases/latest"
const ProgressFullChar = "#"
const ProgressEmptyChar = " "
const StartupLoadingSeconds = 2
const StartupTickDuration = time.Millisecond * 16

const AppLocalDataDir = ".io_retrocode_keylogger"
