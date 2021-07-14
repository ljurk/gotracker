// Demo code for the Table primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)

	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	headers := [...]string{"KK", "SD", "HH", "CY"}
	rows := 16
	for r := 0; r < rows; r++ {
		for c := 0; c < len(headers); c++ {
			color := tcell.ColorWhite
			if r == 0 {
				color = tcell.ColorYellow
				table.SetCell(r, c,
					tview.NewTableCell(headers[c]).
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
			} else {
				table.SetCell(r, c,
					tview.NewTableCell("").
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
			}
		}
	}

	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}

	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		if table.GetCell(row, column).Text == "" {
			table.GetCell(row, column).SetText("X")
		} else {
			table.GetCell(row, column).SetText("")
		}
		table.SetSelectable(false, false)
	})
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyUp {
			table.GetCell(table.GetSelection()).SetText("Y")
		}
		return event
	})

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(table, 1, 0, 1, 3, 0, 0, true).
		AddItem(list, 2, 0, 1, 3, 0, 0, false)
	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
