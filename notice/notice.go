package notice

import (
	"path/filepath"

	"fmt"

	"github.com/martinlindhe/notify"
)

var imagePath, _ = filepath.Abs("neko.png")

func Pop(appName string, title string, text string) {
	fmt.Println(title, text)
	notify.Notify(appName, title, text, imagePath)
}
