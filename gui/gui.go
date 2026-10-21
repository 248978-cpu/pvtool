package gui

import (
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func Toon() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("PV-tool"))
		if err := run(w); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			material.Body1(th, "Hallo, Loser!").Layout(gtx)
			e.Frame(gtx.Ops)
		case app.DestroyEvent:
			return e.Err
		}
	}
}
