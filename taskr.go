// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"taskr/parseTasks"
	"fmt"
)

func main() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()


	p0 := widgets.NewParagraph()
	p0.Text = "Borderless Text"
	p0.SetRect(104, 50, 20, 5)
	p0.Border = false





	parsed := parseTasks.ParseTasks() 

	l := widgets.NewList()
	l.Title = "List"
	for i := 0; i < len(parsed.Tasks); i++ {
		title := fmt.Sprintf("[%s] %s", parsed.Tasks[i].ID, parsed.Tasks[i].Title)
		l.Rows = append(l.Rows, title)
	}
//	l.Rows = append(l.Rows, "TEST")

//	l.Rows = []string{
//		"[0] github.com/gizak/termui/v3",
//		"[1] [你好，世界](fg:blue)",
//		"[2] [こんにちは世界](fg:red)",
//		"[3] [color](fg:white,bg:green) output",
//		"[4] output.go",
//		"[5] random_out.go",
//		"[6] dashboard.go",
//		"[7] foo",
//		"[8] bar",
//		"[9] baz",
//	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 100, 20) // left margin/padding? top margin/padding, width, height?

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		case "o":
			ui.Close()
			ui.Render(p0)
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}
}
