package notice

import (
	"fmt"
	"github.com/martinlindhe/notify"
	"path/filepath"
	"strconv"
	"time"
)

func Pop(appName string, title string, text string, icon string) {
	path, _ := filepath.Abs(icon)
	if appName == "" {
		appName = "meow"
	}
	if title == "" {
		appName = "meow!"
	}
	notify.Notify(appName, title, text, path)
	return
}

func Snooze(appName string, title string, text string, interval string, icon string) {
	timeArgS, es := strconv.Atoi(interval)
	if es != nil {
		fmt.Println(es)
	}
	tds := time.Duration(timeArgS)
	tickers := time.NewTicker(tds * time.Second)
	for range tickers.C {
		Pop("meow", "meow", text, icon)
	}

	return
}
