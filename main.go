package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
	"github.com/skratchdot/open-golang/open"
)

const defaultParams string = "-noLauncher"

func main() {

	gtk.Init(&os.Args)
	if builder, err := gtk.BuilderNew(); err != nil {
		log.Fatal(err)
	} else if err := builder.AddFromFile("gtk3/launcher.glade"); err != nil {
		log.Fatal(err)
	} else if winObj, err := builder.GetObject("window"); err != nil {
		log.Fatal(err)
	} else if NoSplashChkbttnObj, err := builder.GetObject("chkbutton_noSplash"); err != nil {
		log.Fatal(err)
	} else if NoPauseChkbttnObj, err := builder.GetObject("chkbutton_noPause"); err != nil {
		log.Fatal(err)
	} else if PlayBttnObj, err := builder.GetObject("button_play"); err != nil {
		log.Fatal(err)
	} else if AddParamsTextObj, err := builder.GetObject("text_addParams"); err != nil {
		log.Fatal(err)
	} else {
		window := winObj.(*gtk.Window)
		window.Connect("destroy", func() {
			gtk.MainQuit()
		})
		window.ShowAll()

		chkbtnNoSplash := NoSplashChkbttnObj.(*gtk.CheckButton)
		chkbtnNoPause := NoPauseChkbttnObj.(*gtk.CheckButton)
		textAddParams := AddParamsTextObj.(*gtk.Entry)

		PlayBttn := PlayBttnObj.(*gtk.Button)
		PlayBttn.Connect("clicked", func() {
			var Params string
			if chkbtnNoSplash.GetActive() == true {
				Params += " -noSplash "
			}
			if chkbtnNoPause.GetActive() == true {
				Params += " -noPause "
			}
			AddParams, _ := textAddParams.GetText()

			RunCommand := "steam://run/107410//" + defaultParams + Params + AddParams + "/"
			open.Run(RunCommand)
			fmt.Printf("Starting Arma with: \"%v\"\n", RunCommand)
		})
	}
	gtk.Main()

}
