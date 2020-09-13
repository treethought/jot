package ui

import (
	"github.com/rivo/tview"
	"github.com/treethought/jot/pkg/app"
)

type UI struct {
	app     *tview.Application
	Widgets map[string]WidgetRenderer
	jot     *app.App
	state   *State
}

func NewUI() (ui *UI, err error) {
	ui = &UI{}

	ui.state = NewState()
	ui.app = tview.NewApplication()
	ui.jot = app.NewApp()

	ui.Widgets = make(map[string]WidgetRenderer)
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

func (ui *UI) initWidgets() error {
	notes, err := ui.jot.ListNotes()
	if err != nil {
		return err
	}

	listing := NewNoteList(ui, notes)
	ui.Widgets["note_list"] = listing

	view := NewNoteView(ui)
	ui.Widgets["view"] = view
	return nil

}

func (ui *UI) initGrid() error {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	// menu := newPrimitive("Menu")
	// main := newPrimitive("Main content")

	// sideBar := newPrimitive("Side Bar")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	ui.app.SetRoot(grid, true)
	for _, w := range ui.Widgets {
		w.Render(grid)
	}
	// // Layout for screens narrower than 100 cells (menu and side bar are hidden).
	// grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
	// 	AddItem(main, 1, 0, 1, 3, 0, 0, false).
	// 	AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// // Layout for screens wider than 100 cells.
	// grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
	// 	AddItem(main, 1, 1, 1, 1, 0, 100, false).
	// 	AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	return nil

}
