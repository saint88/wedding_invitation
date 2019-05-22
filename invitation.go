package main

import (
	"github.com/atotto/clipboard"
	"github.com/jroimartin/gocui"
	"github.com/skratchdot/open-golang/open"
	"log"
	"net/http"
	"stash.mail.ru/qafeta/feta-media-tools/wedding_invitation/gui"
	"stash.mail.ru/qafeta/feta-media-tools/wedding_invitation/server"
)

var clear map[string]func()
var s *http.Server

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	server.Start()
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManager(gocui.ManagerFunc(gui.Description), gocui.ManagerFunc(gui.WeddingTime), gocui.ManagerFunc(gui.EventPlace), gocui.ManagerFunc(gui.Controls))

	err = g.SetKeybinding("", 'm', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		clipboard.WriteAll(gui.PLACE_URL)
		open.Run(gui.PLACE_URL)

		return nil
	})
	check(err)

	err = g.SetKeybinding("", 'ь', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		clipboard.WriteAll(gui.PLACE_URL)
		open.Run(gui.PLACE_URL)

		return nil
	})
	check(err)

	err = g.SetKeybinding("", 'i', gocui.ModNone, func(g *gocui.Gui, view *gocui.View) error {
		open.Run("http://127.0.0.1:5432/")

		return nil
	})
	check(err)

	err = g.SetKeybinding("", 'ш', gocui.ModNone, func(g *gocui.Gui, view *gocui.View) error {
		open.Run("http://127.0.0.1:5432/")

		return nil
	})
	check(err)

	err = g.SetKeybinding("", 'q', gocui.ModNone, gui.Quit)
	check(err)

	err = g.SetKeybinding("", 'й', gocui.ModNone, gui.Quit)
	check(err)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
