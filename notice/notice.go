package notice

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/kyokomi/emoji"
	"github.com/martinlindhe/notify"
	"gopkg.in/toast.v1"
)

var imagePath, _ = filepath.Abs("./img/neko.png")

func Pop(appName string, title string, text string, endless bool) {
	fmt.Println(title+emoji.Sprint(":cat2:"), text)
	if runtime.GOOS == "windows" {
		notification := toast.Notification{
			AppID:   appName,
			Title:   title,
			Message: text,
			Icon:    imagePath, // This file must exist (remove this line if it doesn't)
			Loop:    endless,
		}
		err := notification.Push()
		if err != nil {
			log.Fatalln(err)
		}
	}
	notify.Notify(appName, title, text, imagePath)
}
