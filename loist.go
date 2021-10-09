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
