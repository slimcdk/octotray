package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/slimcdk/octotray/internal/icon"

	"github.com/getlantern/systray"
	octo "github.com/mcuadros/go-octoprint"
)

var (
	client *octo.Client
)

func main() {

	if client == nil {
		fmt.Println("Please set a host")
		// Open input dialog here
	}

	onExit := func() {
		now := time.Now()
		ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle(" OctoTray")

	systray.AddSeparator()

	// Quit app goroutine
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuitOrig.SetIcon(icon.Data)
}
