// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	log "github.com/sirupsen/logrus"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"taskr/parseTasks"
	"fmt"
)

func displayDescription(description string) {
	ui.Clear()
  p0 := widgets.NewParagraph()
	p0.Text = description
	p0.SetRect(104, 50, 20, 5)
	p0.Border = false
	p0.WrapText = true
	ui.Render(p0)
}

func main() {
  log.SetFormatter(&log.JSONFormatter{})
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	parsed := parseTasks.ParseTasks() 

	l := widgets.NewList()
	l.Title = "List"
	for i := 0; i < len(parsed.Tasks); i++ {
		title := fmt.Sprintf("[%s] %s", parsed.Tasks[i].ID, parsed.Tasks[i].Title)
		l.Rows = append(l.Rows, title)
	}

	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 100, 20) // left margin/padding? top margin/padding, width, height?

	ui.Render(l)

//	log.Info("PARSED: ")
//	log.Info(parsed.Tasks[0])

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
	    ui.Render(l)
		case "k", "<Up>":
			l.ScrollUp()
	    ui.Render(l)
		case "<C-d>":
			l.ScrollHalfPageDown()
	    ui.Render(l)
		case "<C-u>":
			l.ScrollHalfPageUp()
	    ui.Render(l)
		case "<C-f>":
			l.ScrollPageDown()
	    ui.Render(l)
		case "<C-b>":
			l.ScrollPageUp()
	    ui.Render(l)
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
	      ui.Render(l)
			}
		case "<Home>":
			l.ScrollTop()
	    ui.Render(l)
		case "G", "<End>":
			l.ScrollBottom()
	    ui.Render(l)
		case "o":
	  //  log.Info(l.SelectedRow)
		  displayDescription(parsed.Tasks[l.SelectedRow].Description)	
			ui.Clear()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}


	}
}
