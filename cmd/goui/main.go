package main

import (
	"github.com/elizarpif/goui/window"
	"github.com/elizarpif/logger"
	"os"

	"github.com/elizarpif/goui/ui"
	"github.com/therecipe/qt/widgets"
)

func main() {
	log := logger.NewLogger()
	ctx := logger.SetContext(log)

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// create a mainWindow
	// with a minimum size of 250*200
	// and sets the title to "Hello Widgets Example"
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetMinimumSize2(250, 200)
	mainWindow.SetWindowTitle("Hello Widgets Example")

	uiWindow := &ui.UICryptMainWindow{}
	uiWindow.SetupUI(mainWindow)

	w := window.NewWindow(uiWindow)
	w.Connect(ctx)

	// make the mainWindow visible
	mainWindow.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the mainWindow is closed by the user
	app.Exec()
}
