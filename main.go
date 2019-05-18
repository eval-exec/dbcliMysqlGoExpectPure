package main

import (
	"fmt"
	"io"
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

			outputBtte, err := execCommand.CombinedOutput()
			if err != nil {
				outErr = fmt.Sprint(err)
			} else {
				outErr = ""
			}
			out = string(outputBtte)


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

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func debug(str interface{}, stxr string) {
	fmt.Println("--------------------------------------------------------")
	fmt.Sprintln(str)
	fmt.Println(stxr)
	fmt.Println("--------------------------------------------------------")
}
