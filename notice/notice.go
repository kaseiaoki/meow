package notice

import (
	"fmt"
	"github.com/kaseiaoki/meow/config"
	"github.com/martinlindhe/notify"
	"path/filepath"
	"strconv"
	"time"
)

func Pop(text string) {
	path, _ := filepath.Abs(config.ToastConf.Icon)
	appName := config.ToastConf.AppName
	title := config.ToastConf.Title
	if appName == "" {
		appName = "meow"
	}
	if title == "" {
		title = "meow!"
	}
	notify.Notify(appName, title, text, path)
	return
}

func Snooze(text string, interval string) {
	timeArgS, es := strconv.Atoi(interval)
	if es != nil {
		fmt.Println(es)
	}
	tds := time.Duration(timeArgS)
	tickers := time.NewTicker(tds * time.Second)
	for range tickers.C {
		Pop(text)
	}

	return
}
