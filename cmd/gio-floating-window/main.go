package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Hello World"),
			app.TopMost(true),
			app.Decorated(false),
		)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

func run(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			e.Frame(gtx.Ops)
		}
	}
}
