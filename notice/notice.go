package notice

import (
	"fmt"
	"path/filepath"

	"github.com/kyokomi/emoji"
	"github.com/martinlindhe/notify"
)

var imagePath, _ = filepath.Abs("./img/neko.png")

func Pop(appName string, title string, text string) {
	fmt.Println(title+emoji.Sprint(":cat2:"), text)

	notify.Notify(appName, title, text, imagePath)
	return
}
