package window

import (
	"context"
	"fmt"
	"strconv"

	"github.com/elizarpif/goui/internal/binary"
	"github.com/elizarpif/goui/internal/lab2"
	"github.com/elizarpif/goui/internal/lab3"

	"github.com/elizarpif/logger"
)

func (w *Window) systemBases() {
	w.uiWindow.Answer3.SetText("")
	w.uiWindow.Answer.SetText("")
	w.uiWindow.Answer2.SetText("")

	if w.uiWindow.BinaryRadio.IsChecked() {
		w.uiWindow.Line31.SetInputMask("BBBBBBBB")
		w.uiWindow.LinePoly2.SetInputMask("BBBBBBBB")
		w.uiWindow.LinePoly1.SetInputMask("BBBBBBBB")
		w.uiWindow.Line3.SetInputMask("BBBBBBBB")

		return
	}

	w.uiWindow.Line31.SetInputMask("")
	w.uiWindow.LinePoly2.SetInputMask("")
	w.uiWindow.LinePoly1.SetInputMask("")
	w.uiWindow.Line3.SetInputMask("")
}

func (window *Window) ConnectLab3(ctx context.Context) {
	window.uiWindow.BinaryRadio.SetChecked(true)

	window.uiWindow.Line31.ConnectTextChanged(func(text string) {
		window.lab31(ctx)
	})

	window.uiWindow.LinePoly1.ConnectTextChanged(func(text string) {
		if window.uiWindow.LinePoly1.Text() != "" && window.uiWindow.LinePoly2.Text() != "" {
			window.lab32(ctx)
		}
	})

	window.uiWindow.LinePoly2.ConnectTextChanged(func(text string) {
		if window.uiWindow.LinePoly1.Text() != "" && window.uiWindow.LinePoly2.Text() != "" {
			window.lab32(ctx)
		}
	})

	window.uiWindow.RadioElem.ConnectClicked(func(checked bool) {
		window.lab32(ctx)
	})

	window.uiWindow.RadioPolynom.ConnectClicked(func(checked bool) {
		window.lab32(ctx)
	})

	window.uiWindow.Line3.ConnectTextChanged(func(text string) {
		window.lab33(ctx)
	})
}

func (w *Window) getNumInBase(str string) (int, error) {
	if w.uiWindow.BinaryRadio.IsChecked() {
		return binary.ToBinary(str)
	}
	return strconv.Atoi(str)
}

func (w *Window) setNumInBase(num int) string {
	if w.uiWindow.BinaryRadio.IsChecked() {
		return binary.ToBinaryStr(num)
	}

	return strconv.Itoa(num)
}

func (w *Window) lab33(ctx context.Context) {
	uiw := w.uiWindow
	log := logger.GetLogger(ctx)

	text := uiw.Line3.Text()

	if text == "" {
		uiw.Answer3.SetText("")
		log.Warning("no input text")
		return
	}

	num, err := w.getNumInBase(text)
	if err != nil {
		log.WithError(err).WithField("text", text).Error("cannot convert to number")
		uiw.Answer3.SetText(fmt.Sprintf("cannot convert to number"))
		return
	}

	modulo := 256

	inv, err := lab2.ModIn(num, modulo)
	if err != nil {
		uiw.Answer3.SetText(err.Error())
	} else {
		uiw.Answer3.SetText(w.setNumInBase(inv))
	}
}

func (w *Window) lab32(ctx context.Context) {
	uiw := w.uiWindow
	log := logger.GetLogger(ctx)

	text1 := uiw.LinePoly1.Text()

	if text1 == "" {
		uiw.Answer2.SetText("")
		log.Warning("no input text")
		return
	}

	num1, err := w.getNumInBase(text1)
	if err != nil {
		log.WithError(err).WithField("text", text1).Error("cannot convert to binary")
		return
	}

	text2 := uiw.LinePoly2.Text()

	if text2 == "" {
		uiw.Answer2.SetText("")
		log.Warning("no input text")
		return
	}

	num2, err := w.getNumInBase(text2)
	if err != nil {
		log.WithError(err).WithField("text", text2).Error("cannot convert to binary")
		return
	}

	p1 := lab3.NewBinaryPolynom(num1)
	p2 := lab3.NewBinaryPolynom(num2)

	p := p1.Multiply(p2)

	if uiw.RadioPolynom.IsChecked() {
		uiw.Answer2.SetText(p.String())
	}

	if uiw.RadioElem.IsChecked() {
		uiw.Answer2.SetText(w.setNumInBase(p.GetNum()))
	}
}

func (window *Window) lab31(ctx context.Context) {
	w := window.uiWindow
	log := logger.GetLogger(ctx)

	text := w.Line31.Text()

	if text == "" {
		w.Answer.SetText("")
		log.Warning("no input text")
		return
	}

	num, err := window.getNumInBase(text)
	if err != nil {
		log.WithError(err).WithField("text", text).Error("cannot convert to binary")
		w.Answer.SetText(fmt.Sprintf("cannot convert to binary"))
		return
	}

	polynom := lab3.NewBinaryPolynom(num)
	w.Answer.SetText(polynom.String())
}
