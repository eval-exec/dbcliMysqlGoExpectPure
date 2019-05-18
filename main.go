package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rivo/tview"
)

func main() {

	var out string
	var outErr string
	var getcommands string
	// var stdin io.Reader
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var datatext string
	var execCommand *exec.Cmd
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
			execCommand = exec.Command(commandsSlice[0], commandsSlice[1:]...)

			execCommand.Stdin = strings.NewReader(getcommands)
			execCommand.Stdout = &stdout
			execCommand.Stderr = &stderr

			outErr = stderr.String()
			err := execCommand.Run()
			if err != nil {
				outErr = fmt.Sprint(err)
			} else {
				outErr = fmt.Sprint(err) + outErr
			}
			out = stdout.String()

			datatext = fmt.Sprintf("%s\n%s", out, outErr)

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

func debug(str interface{}, stxr string) string {
	x := "-------------------------------------------------------------------\n" + fmt.Sprint(str) + "\n" + stxr + "\n---------------------------------------------"
	return x
}
