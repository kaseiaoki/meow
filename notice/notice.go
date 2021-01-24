package notice

import (
	"github.com/martinlindhe/notify"
)

func Pop(appName string, title string, text string) {
	notify.Notify(appName, title, text, "")
	return
}
