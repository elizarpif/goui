package window

import (
	"context"

	"github.com/elizarpif/goui/ui"
)

type Window struct {
	uiWindow *ui.UICryptMainWindow
}

func NewWindow(ui *ui.UICryptMainWindow) *Window {
	return &Window{
		uiWindow: ui,
	}
}

func (w *Window) Connect(ctx context.Context) {
	if w.uiWindow.TabWidget.Count() > 0 {
		w.uiWindow.TabWidget.SetTabText(1, "lab2")
		w.uiWindow.TabWidget.SetTabText(0, "lab3")
	}

	w.ConnectLab2(ctx)
	w.ConnectLab3(ctx)
}
