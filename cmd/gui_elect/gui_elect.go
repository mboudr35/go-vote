package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"log"
	"os"
)

func run(win *app.Window) error {
	theme := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		ev := <-win.Events()
		switch ev := ev.(type) {
		case system.DestroyEvent:
			return ev.Err
		case system.FrameEvent:
			ctx := layout.NewContext(&ops, ev)
			title := material.H1(theme, "go-vote")
			title.Alignment = text.Middle
			title.Layout(ctx)
			ev.Frame(ctx.Ops)
		}
	}
}

func main() {
	go func() {
		win := app.NewWindow()
		if err := run(win); err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
