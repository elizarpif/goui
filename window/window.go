package window

import (
	"context"
	"fmt"
	"github.com/elizarpif/goui/internal/binary"
	"github.com/elizarpif/logger"

	"github.com/elizarpif/goui/internal/lab3"
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

func (window *Window) Connect(ctx context.Context) {
	window.uiWindow.LineEdit.SetInputMask("BBBBBBBB")
	window.uiWindow.Btn.ConnectClicked(func(bool) {
		// widgets.QMessageBox_Information(nil, "OK", uiWindow.LineEdit.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		window.ButtonReact(ctx)
	})
	window.uiWindow.LineEdit.ConnectTextChanged(func(text string) {
		window.ButtonReact(ctx)
	})
}

func (window *Window) ButtonReact(ctx context.Context) {
	if window.uiWindow.Lab31.IsChecked() {
		window.lab31(ctx)
		return
	}

	window.uiWindow.AnswerLine.SetText("")
}

func (window *Window) lab31(ctx context.Context) {
	w := window.uiWindow
	log := logger.GetLogger(ctx)

	text := w.LineEdit.Text()

	if text == "" {
		w.AnswerLine.SetText("")
		log.Warning("no input text")
		return
	}

	num, err := binary.ToBinary(text)
	if err != nil {
		log.WithError(err).WithField("text", text).Error("cannot convert to binary")
		window.BinaryNumError()
		return
	}

	polynom := lab3.NewBinaryPolynom(num)
	w.AnswerLine.SetText(polynom.String())
}

func (w *Window) BinaryNumError() {
	w.uiWindow.AnswerLine.SetText(fmt.Sprintf("cannot convert to binary"))
}
