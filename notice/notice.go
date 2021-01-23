package notice

import (
	"fmt"
	"github.com/martinlindhe/notify"
)

func Pop(appName string, title string, text string) {
	fmt.Println(title, text)
	notify.Notify(appName, title, text, "")
	return
}
