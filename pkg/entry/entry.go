package entry

import "keylogger/utils"

func AppEntry() {
	defer utils.Recover(false)

	runCLI()
}
