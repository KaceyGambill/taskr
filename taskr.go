// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	log "github.com/sirupsen/logrus"

	"fmt"
	"taskr/parseTasks"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

)

//type Tasks struct {
//	ID string,
//	Title string,
//	Description string,
//	Flags []string,
//}

func displayDescription(description string) {
	ui.Clear()
	p0 := widgets.NewParagraph()
	p0.Text = description
	p0.SetRect(104, 50, 20, 5)
	p0.Border = false
	p0.WrapText = true
	ui.Render(p0)

	//	uiEvents := ui.PollEvents()
	//	for {
	//		e := <-uiEvents
	//		switch e.ID {
	//		case "q":
	//			return
	//		}
	//	}

}

//func displayTaskList(parsedList parseTasks.Tasks) {
////  log.Info(parsedList)
//  ui.Clear()
//	taskList := widgets.NewList()
//	taskList.Title = "Task List"
//	for i := 0; i < len(parsedList.Tasks); i++ {
//		title := fmt.Sprintf("[%s] %s", parsedList.Tasks[i].ID, parsedList.Tasks[i].Title)
//		taskList.Rows = append(taskList.Rows, title)
//	}
//	taskList.TextStyle = ui.NewStyle(ui.ColorYellow)
//	taskList.WrapText = true
//	taskList.SetRect(0, 0, 100, 20)
//	ui.Render(taskList)
//}
func userInput(description string) {
	ui.Clear()
	p0 := widgets.NewParagraph()
	p0.Text = description
	p0.SetRect(104, 50, 20, 5)
	p0.Border = false
	p0.WrapText = true
	ui.Render(p0)
	fmt.Println("Enter something: ")
	var stringo string
	fmt.Scanln(&stringo)
	p0.Text = stringo
	fmt.Println("Entered: ", stringo)
	ui.Clear()
	ui.Render(p0)
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	keyList := widgets.NewList()
	keyList.Rows = []string{
		"q - quit",
		"j - down, k - up",
		"o - open, e - edit",
		"a - add, d - delete",
	}
	keyList.SetRect(0, 0, 30, 7)
	keyList.Border = true
	keyList.BorderStyle.Fg = ui.ColorYellow

	parsed := parseTasks.ParseTasks()
	//	displayTaskList(parsed)
	taskList := widgets.NewList()
	taskList.Title = "List"
	for i := 0; i < len(parsed.Tasks); i++ {
		title := fmt.Sprintf("[%s] %s", parsed.Tasks[i].ID, parsed.Tasks[i].Title)
		taskList.Rows = append(taskList.Rows, title)
	}

	taskList.TextStyle = ui.NewStyle(ui.ColorYellow)
	taskList.WrapText = false
	taskList.SetRect(0, 7, 100, 20) // left margin/padding? top margin/padding, width, height?

	ui.Render(keyList, taskList)

	//	log.Info("PARSED: ")
	//	log.Info(parsed.Tasks[0])

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
			//    case "q":
			//			ui.Clear()
			//			ui.Render(taskList)
		case "j", "<Down>":
			taskList.ScrollDown()
			ui.Render(keyList, taskList)
		case "k", "<Up>":
			taskList.ScrollUp()
			ui.Render(keyList, taskList)
		case "<C-d>":
			taskList.ScrollHalfPageDown()
			ui.Render(keyList, taskList)
		case "<C-u>":
			taskList.ScrollHalfPageUp()
			ui.Render(keyList, taskList)
		case "<C-f>":
			taskList.ScrollPageDown()
			ui.Render(keyList, taskList)
		case "<C-b>":
			taskList.ScrollPageUp()
			ui.Render(keyList, taskList)
		case "g":
			if previousKey == "g" {
				taskList.ScrollTop()
				ui.Render(keyList, taskList)
			}
		case "<Home>":
			taskList.ScrollTop()
			ui.Render(keyList, taskList)
		case "G", "<End>":
			taskList.ScrollBottom()
			ui.Render(keyList, taskList)
		case "o":
			//  log.Info(l.SelectedRow)
			displayDescription(parsed.Tasks[taskList.SelectedRow].Description)
			ui.Clear()
		case "e":
			userInput(parsed.Tasks[taskList.SelectedRow].Description)
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

	}
}
