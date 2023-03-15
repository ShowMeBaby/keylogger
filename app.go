package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

type KeyEvent struct {
	KeyChar string
	KeyCode uint16
}

func (t *KeyEvent) String() string {
	return fmt.Sprintf("{KeyChar:%s,KeyCode:%d}", t.KeyChar, t.KeyCode)
}

func main() {
	fmt.Println("键盘监听已启动!!")
	hooks := hook.Start()
	defer hook.End()
	for ev := range hooks {
		if ev.Kind == hook.MouseUp || ev.Kind == hook.KeyUp {
			if ev.Button == hook.MouseMap["right"] {
				fmt.Printf("mouseRight\n")
			} else if ev.Button == hook.MouseMap["left"] {
				fmt.Printf("mouseLeft\n")
			} else {
				KeyJson := KeyEvent{KeyChar: RawcodetoKeychar(ev.Rawcode), KeyCode: ev.Rawcode}
				fmt.Printf("=>%v<= \n", KeyJson.String())
			}
		}
	}
}
