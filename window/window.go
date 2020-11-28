package window

import (
	"context"
	"github.com/elizarpif/goui/internal/binary"
	"github.com/elizarpif/goui/internal/lab3"
	"github.com/elizarpif/goui/ui"
	"github.com/elizarpif/logger"
)

type Window struct {
	uiWindow *ui.UICryptMainWindow
}

func NewWindow(ui *ui.UICryptMainWindow) *Window {
	return &Window{
		uiWindow: ui,
	}
}

func (window *Window) Connect(ctx context.Context) {
	window.uiWindow.Btn.ConnectClicked(func(bool) {
		// widgets.QMessageBox_Information(nil, "OK", uiWindow.LineEdit.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		window.ButtonReact(ctx)
	})
}

func (window *Window) ButtonReact(ctx context.Context) {
	if window.uiWindow.Lab31.IsChecked() {
		window.lab31(ctx)
	}
}

func (window *Window) lab31(ctx context.Context) {
	w := window.uiWindow
	log := logger.GetLogger(ctx)

	text := w.LineEdit.Text()

	if text == "" {
		log.Warning("no input text")
		return
	}

	isBin := binary.IsBinary(text)
	if !isBin {
		var err error
		text, err = binary.ToBinary(text)

		if err != nil {
			log.WithError(err).WithField("text", text).Error("cannot convert to binary")
			return
		}
	}

	res := lab3.PolynomForm(text)
	w.AnswerLine.SetText(res)
}