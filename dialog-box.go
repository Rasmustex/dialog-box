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
	// custom usage prompt
	flag.CommandLine.Usage = func() {
		fmt.Println("Usage: dialog-box [OPTION]... BUTTONS")
		fmt.Println()
		fmt.Println("Create graphical dialog box with BUTTONS that have the specified names.")
		fmt.Println("Options:")
		fmt.Println("	-p        specify text in prompt (default \"Pick an option\")")
		fmt.Println("	-w        specify window title (default \"Dialog Box\")")
		fmt.Println("	-h        show usage")
	}

	// Parse command line flags
	promptPtr := flag.String("p", "Pick an option", "The prompt for the dialog box")
	wnamePtr := flag.String("w", "Dialog Box", "Window title for the dialog box")
	flag.Parse()

	// declare names of buttons from command line args
	button_names := flag.Args()
	// and get amount of buttons
	n_buttons := len(button_names)

	a := app.New()
	w := a.NewWindow(*wnamePtr)

	// create parent container
	content := container.New(layout.NewVBoxLayout())
	// create centered label with prompt and put in in the parent container
	prompt := container.New(layout.NewCenterLayout(), widget.NewLabel(*promptPtr))
	content.Add(prompt)

	// create grid child container to hold buttons
	buttons := container.New(layout.NewGridLayout(n_buttons))

	// create all the buttons with their respective names
	for _, element := range button_names {
		btn := element
		buttons.Add(widget.NewButton(btn, func() {
			fmt.Println(btn)
			w.Close()
		}))
	}
	// add button container to parent container
	content.Add(buttons)

	w.SetContent(content)
	w.ShowAndRun()
}
