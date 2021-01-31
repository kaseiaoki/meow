package notice

import (
	"fmt"
	"github.com/martinlindhe/notify"
	"strconv"
	"time"
)

func Pop(appName string, title string, text string) {
	notify.Notify(appName, title, text, "")
	return
}

func Snooze(appName string, title string, text string, interval string) {
	timeArgS, es := strconv.Atoi(interval)
	if es != nil {
		fmt.Println(es)
	}
	tds := time.Duration(timeArgS)
	tickers := time.NewTicker(tds * time.Second)
	for range tickers.C {
		Pop("meow", "meow", text)
	}

	return
}
