package notice

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/kyokomi/emoji"
	"gopkg.in/toast.v1"
)

var imagePath, _ = filepath.Abs("neko.png")

func Pop(appName string, title string, text string) {
	fmt.Println(title+emoji.Sprint(":cat2:"), text)
	// notify.Alert(appName, title, text, imagePath)
	notification := toast.Notification{
		AppID:   appName,
		Title:   title,
		Message: text,
		Icon:    imagePath, // This file must exist (remove this line if it doesn't)
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
