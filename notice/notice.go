package notice

import (
	"fmt"
	"github.com/kaseiaoki/meow/config"
	"github.com/martinlindhe/notify"
	"path/filepath"
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
	t, err := time.ParseDuration(interval)
	if err != nil {
		fmt.Println(err)
	}
	td := time.Duration(t)
	ticker := time.NewTicker(td)
	for range ticker.C {
		Pop(text)
	}

	return
}
