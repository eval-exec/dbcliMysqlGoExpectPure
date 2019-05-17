package main

import (
	"fmt"
	"os/exec"
	"strings"

	"database/sql"

	"github.com/gdamore/tcell"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rivo/tview"
)

func main() {

	var out string

	var getcommands string

	var datatext string
	datatext = "ok I am data text"

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	main := tview.NewTextView().SetText(datatext).SetTextAlign(tview.AlignLeft)

	botTextInput := func(text string) (trxPri tview.Primitive) {

		tinput := tview.NewInputField().SetLabel("mysql commands:").SetFieldWidth(0).SetText(text)

		tinput.SetDoneFunc(func(key tcell.Key) {
			main.SetText("")

			getcommands = tinput.GetText()
			commandsSlice := strings.Fields(getcommands)
			execCommand := exec.Command(commandsSlice[0], commandsSlice[1:]...)
			output, err := execCommand.CombinedOutput()
			if err != nil {
				datatext = fmt.Sprintf("%s", err)
			}

			out = fmt.Sprint(string(output))
			datatext = fmt.Sprintf("%s", out)

			main.SetText(datatext)
			tinput.SetText("")

		})

		trxPri = tinput
		return
	}

	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(50, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("MySQL command client"), 0, 0, 1, 3, 0, 0, false).
		AddItem(botTextInput(""), 2, 0, 1, 3, 0, 0, true)

	grid.AddItem(main, 1, 0, 1, 3, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}

func DBexec(tDB **sql.DB, str string) (returnRows *sql.Rows, err error) {
	returnRows, err = (*tDB).Query(str)
	if err != nil {
		debug(err, "err in dbexec func ")
	}
	return
}

func debug(str interface{}, stxr string) {
	fmt.Println("--------------------------------------------------------")
	fmt.Sprintln(str)
	fmt.Println(stxr)
	fmt.Println("--------------------------------------------------------")
}




