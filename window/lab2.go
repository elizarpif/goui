package window

import (
	"context"
	"strconv"

	"github.com/elizarpif/goui/internal/lab2"
	"github.com/elizarpif/logger"
)

func (w *Window) ConnectLab2(ctx context.Context) {
	w.uiWindow.BinPowA.ConnectTextChanged(func(text string) {
		w.binPow(ctx)
	})

	w.uiWindow.BinPowB.ConnectTextChanged(func(text string) {
		w.binPow(ctx)
	})

	w.uiWindow.GcdA.ConnectTextChanged(func(text string) {
		w.gcd(ctx)
	})

	w.uiWindow.GcdB.ConnectTextChanged(func(text string) {
		w.gcd(ctx)
	})
}

func (w *Window) binPow(ctx context.Context) {
	if w.uiWindow.BinPowA.Text() == "" || w.uiWindow.BinPowB.Text() == "" {
		return
	}

	text1 := w.uiWindow.BinPowA.Text()

	num1, err := strconv.Atoi(text1)
	if err != nil {
		logger.GetLogger(ctx).WithError(err).Error("cannot convert to num")
		return
	}

	text2 := w.uiWindow.BinPowB.Text()

	num2, err := strconv.Atoi(text2)
	if err != nil {
		logger.GetLogger(ctx).WithError(err).Error("cannot convert to num")
		return
	}

	modulo, err := strconv.Atoi(w.uiWindow.Lab2Modulo.Text())
	if err != nil {
		logger.GetLogger(ctx).WithError(err).Error("cannot convert modulo to num")
		return
	}

	res := lab2.BinPow(num1, num2, modulo)
	w.uiWindow.BinPow.SetText(strconv.Itoa(res))
}

func (w *Window) gcd(ctx context.Context) {
	if w.uiWindow.GcdA.Text() == "" || w.uiWindow.GcdB.Text() == "" {
		return
	}

	text1 := w.uiWindow.GcdA.Text()

	num1, err := strconv.Atoi(text1)
	if err != nil {
		logger.GetLogger(ctx).WithError(err).Error("cannot convert to num")
		return
	}

	text2 := w.uiWindow.GcdB.Text()

	num2, err := strconv.Atoi(text2)
	if err != nil {
		logger.GetLogger(ctx).WithError(err).Error("cannot convert to num")
		return
	}

	gcd, x, y := lab2.RecursiveGCD(num1, num2)

	w.uiWindow.Gcd.SetText(strconv.Itoa(gcd))
	w.uiWindow.GcdX.SetText(strconv.Itoa(x))
	w.uiWindow.GcdY.SetText(strconv.Itoa(y))
}
