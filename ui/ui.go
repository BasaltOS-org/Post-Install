package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ShowUI initializes and shows the Fyne UI
func ShowUI() {
	// Create a new Fyne application
	app := app.New()
	window := app.NewWindow("Basalt Post Install")
	window.Resize(fyne.NewSize(800, 600))

	// Create UI components
	title := widget.NewLabel("Welcome to Basalt Post Install")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	content := container.NewPadded(
		widget.NewLabel("What do you want to do?"),
	)

	// Create buttons
	clickMeBtn := widget.NewButton("Install stuff", func() {
		log.Println("Button clicked!")
	})

	settingsBtn := widget.NewButton("Settings", func() {
		log.Println("Settings clicked!")
	})

	// Create button container with right alignment
	buttons := container.NewHBox(
		layout.NewSpacer(),
		clickMeBtn,
		settingsBtn,
	)

	// Create main layout
	mainContent := container.NewVBox(
		title,
		container.NewPadded(content),
		layout.NewSpacer(),
		container.NewPadded(buttons),
	)

	// Set the main content
	window.SetContent(mainContent)

	// Show and run the application
	window.ShowAndRun()
}