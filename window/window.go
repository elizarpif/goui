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

func (w *Window) Connect(ctx context.Context){
	w.ConnectLab2(ctx)
	w.ConnectLab3(ctx)
}