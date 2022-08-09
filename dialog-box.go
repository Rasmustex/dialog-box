package main

import (
	"fmt"
	"flag"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	// let's parse command line flags
	promptPtr := flag.String("p", "Pick an option", "The prompt for the dialog box")
	wnamePtr := flag.String("w", "Dialog Box", "Window title for the dialog box")
	flag.Parse()


	button_names := flag.Args()
	n_buttons := len(button_names)

	a := app.New()
	w := a.NewWindow(*wnamePtr)

	content := container.New(layout.NewVBoxLayout())
	prompt := container.New(layout.NewCenterLayout(), widget.NewLabel(*promptPtr))
	content.Add(prompt)

	buttons := container.New(layout.NewGridLayout(n_buttons))

	// use idx for button offset
	for _, element := range button_names {
		btn := element
		buttons.Add(widget.NewButton(btn, func() {
			fmt.Println(btn)
			w.Close()
		}))
	}

	content.Add(buttons)
	w.SetContent(content)
	w.ShowAndRun()
}
