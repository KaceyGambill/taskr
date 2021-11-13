package main

import (
	"github.com/rivo/tview"
	log "github.com/sirupsen/logrus"
)
func showBox() {
	box := tview.NewBox()
	box.SetTitle("BIGBOX")
	box.SetBorder(true)
	app := tview.NewApplication()
	err :=	app.SetRoot(box, true).Run()
	if err != nil {
		log.Fatal(err)
	}
}
func textEnter() {
	app := tview.NewApplication()
	inputField := tview.NewInputField()
	inputField.SetTitle("AlligatorBox")
	inputField.SetText("Toasty")
	
	err := app.SetRoot(inputField, true).SetFocus(inputField).Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(inputField.GetText())
}
func main() {
	log.SetFormatter(&log.JSONFormatter{})

	box := tview.NewBox()
	box.SetTitle("BIGBOX")
	box.SetBorder(true)
  app := tview.NewApplication()
	list := tview.NewList()
	list.AddItem("List item 1", "Some explanatory text", 'a', nil)
  list.AddItem("List item 2", "Some explanatory text", 'b', nil)
  list.AddItem("List item 3", "Some explanatory text", 'c', func() {
		textEnter()
	})
	list.AddItem("List item 4", "Some explanatory text", 'd', func() {
		showBox()
	})
	list.AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
		log.Info(list)
	err := app.SetRoot(list, true).SetFocus(list).Run()
//	if err != nil {
//		log.Fatal(err)
//	}
	inputField := tview.NewInputField()
	inputField.SetTitle("AlligatorBox")
	inputField.SetText("Toasty")
	
//	err := app.SetRoot(inputField, false).SetFocus(inputField).Run()
	log.Info(inputField.GetText())

 if err != nil {
	 log.Fatal(err)
 }


}
