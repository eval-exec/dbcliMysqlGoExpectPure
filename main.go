package main

import (
	// "bytes"
	"fmt"
	"log"
	"os"
	"regexp"

	// "os/exec"
	// "strings"

	"github.com/gdamore/tcell"
	_ "github.com/go-sql-driver/mysql"
	expect "github.com/google/goexpect"
	"github.com/rivo/tview"
)

func main() {

	var out string
	// var outErr string
	var getcommands string
	// var stdin io.Reader
	// var stdout []byte
	// var stderr []byte

	var datatext string
	// var execCommand *exec.Cmd
	datatext = "ok I am data text"
	var e *expect.GExpect
	var err error

	var chanErr <-chan error
	var start = 0
	var regOut = regexp.MustCompile(" ")
	var regIn = regexp.MustCompile(" ")
	logfile, err := os.Create("mysql.log")
	if err != nil {
		log.Fatal("logfile error ")
	}
	logger := log.New(logfile, "", log.Llongfile)
	logger.SetFlags(log.LstdFlags)

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

			logger.Println(getcommands, "is getcommands")

			if start == 0 {
				logger.Println(start, "is start")

				e, chanErr, err = expect.Spawn(getcommands, -1)
				logger.Println(getcommands, "spawn")
				if err != nil {
					logger.Println(err)
				}
				if getcommands == "mysql" {
					logger.Println(getcommands, " is mysql commands")
					start = 1
				}
			} else if start != 0 {
				logger.Println(start, "is start")
				e.Expect(regIn,-1)
				err = e.Send(getcommands)
				logger.Println(getcommands, "send")
				if err != nil {
					logger.Println(err)
				}
			}
			result, stringx, err := e.Expect(regOut, -1)
			logger.Println(result, stringx, err, "is e.rxprct output")

			if err != nil {
				logger.Println(err)
			}

			if err != nil {
				logger.Println(err)
			}

			out = result
			datatext = fmt.Sprintf("%s\n", out)

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
