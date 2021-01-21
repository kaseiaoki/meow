package notice

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/kyokomi/emoji"
	"gopkg.in/toast.v1"
)

var imagePath, _ = filepath.Abs("neko.png")

func Pop(appName string, title string, text string, endless bool) {
	fmt.Println(title+emoji.Sprint(":cat2:"), text)

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
