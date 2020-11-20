package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/treethought/jot/app"
)

type UI struct {
	app     *tview.Application
	Widgets []WidgetRenderer
	jot     *app.App
	state   *State
	grid    *tview.Grid
}

func NewUI() (ui *UI, err error) {
	ui = &UI{}

	ui.state = NewState()
	ui.app = tview.NewApplication()
	ui.jot = app.NewApp()

	ui.Widgets = []WidgetRenderer{}
	return

}

func (ui *UI) Start() {

	err := ui.initWidgets()
	if err != nil {
		return
	}

	err = ui.initGrid()
	if err != nil {
		return
	}

	err = ui.app.Run()
	if err != nil {
		panic(err)
	}
	return

}

func (ui *UI) reload() error {
	for _, w := range ui.Widgets {
		err := w.Render(ui.grid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ui *UI) getWidget(name string) (w WidgetRenderer, ok bool) {
	for _, w := range ui.Widgets {
		if w.Name() == name {
			return w, true
		}
	}
	return nil, false
}

func (ui *UI) initWidgets() error {
	notes, err := ui.jot.ListNotes()
	if err != nil {
		return err
	}

	listing := NewNoteList(ui, notes)
	ui.Widgets = append(ui.Widgets, listing)

	view := NewNoteView(ui)
	ui.Widgets = append(ui.Widgets, view)
	return nil

}

func (ui *UI) initGrid() error {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("jot"), 0, 0, 1, 3, 0, 0, false)

	grid.SetBackgroundColor(tcell.ColorDefault)
	ui.grid = grid
	ui.app.SetRoot(grid, true)
	for _, w := range ui.Widgets {
		w.Render(grid)
	}

	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		switch key {
		case tcell.KeyTab:
			if ui.state.currentIdx == len(ui.Widgets)-1 {
				ui.state.currentIdx = -1
			}
			ui.state.currentIdx += 1
			widget := ui.Widgets[ui.state.currentIdx]

			ui.app.SetFocus(widget.View())
			return nil

		}
		return event

	})

	return nil

}
