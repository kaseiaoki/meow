package notice

import (
	"fmt"
	"github.com/martinlindhe/notify"
	"time"
)

func Pop(appName string, title string, text string) {
	notify.Notify(appName, title, text, "")
	return
}

func Snooze(appName string, title string, text string, interval string) {
	t, err := time.ParseDuration(interval)
	if err != nil {
		fmt.Println(err)
	}

	td := time.Duration(t)

	tickers := time.NewTicker(td)
	for range tickers.C {
		Pop("meow", "meow", text)
	}

	return
}
