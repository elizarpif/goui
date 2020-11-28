// WARNING! All changes made in this file will be lost!
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UICryptMainWindow struct {
	Centralwidget     *widgets.QWidget
	TabWidget         *widgets.QTabWidget
	Lab3              *widgets.QWidget
	PlainTextEdit3    *widgets.QPlainTextEdit
	PlainTextEdit     *widgets.QPlainTextEdit
	PlainTextEdit2    *widgets.QPlainTextEdit
	Group3            *widgets.QGroupBox
	Line3             *widgets.QLineEdit
	Answer3           *widgets.QLineEdit
	Label3            *widgets.QLabel
	BaseBox           *widgets.QGroupBox
	BinaryRadio       *widgets.QRadioButton
	DecRadio          *widgets.QRadioButton
	Group1            *widgets.QGroupBox
	Line31            *widgets.QLineEdit
	Answer            *widgets.QLineEdit
	Label31           *widgets.QLabel
	Group2            *widgets.QGroupBox
	LinePoly2         *widgets.QLineEdit
	LinePoly1         *widgets.QLineEdit
	Answer2           *widgets.QLineEdit
	Label32           *widgets.QLabel
	RadioPolynom      *widgets.QRadioButton
	RadioElem         *widgets.QRadioButton
	Lab3LabelMultiply *widgets.QLabel
	Lab2              *widgets.QWidget
	LabelGCD          *widgets.QLabel
	GcdA              *widgets.QLineEdit
	GcdB              *widgets.QLineEdit
	GcdY              *widgets.QLineEdit
	Gcd               *widgets.QLineEdit
	GcdX              *widgets.QLineEdit
	LabelGcdPlus      *widgets.QLabel
	LabelGcdEq        *widgets.QLabel
	BinPowLabel       *widgets.QLabel
	BinPowA           *widgets.QLineEdit
	BinPowB           *widgets.QLineEdit
	LabelBinPowEq     *widgets.QLabel
	BinPow            *widgets.QLineEdit
	LabelLab2Modulo   *widgets.QLabel
	Lab2Modulo        *widgets.QLineEdit
	Menubar           *widgets.QMenuBar
	Statusbar         *widgets.QStatusBar
}

func (this *UICryptMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 1024, 1030))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.Centralwidget.SetStyleSheet("QWidget{\nbackground-color: white;\n}\n\nQFrame {\n\tpadding: 10px;\n    border-radius: 10px;        \n    border-width: 5px;\n    border-color: #ae32a0;\n    border-style: double;\n\tbackground-color: white;\n    font: medium Ubuntu;\n    font-size: 20px;\n} \n\n\nQLabel {\n    font: medium Ubuntu;\n    font-size: 20px;\n    color: #ae32a0;\n    border-width: 0px;\n}        \n\nQPushButton {\n    background-color: rgb(47, 155, 255);\n    color: white;\n      font: medium Ubuntu;\n    font-size: 20px;\n    border-radius: 20px;        \n    border-width: 1px;\n    border-color: rgb(172, 194, 255);\n    border-style: solid;\n}\nQPushButton:hover {\n    background-color: rgb(255, 170, 0);\n}\nQPushButton:pressed {\n    background-color: rgb(255, 153, 29);\n}    \n\nQLineEdit {\n\tborder-color: #d998ce;  \n\tfont: medium Ubuntu;\n    font-size: 20px;\n\tpadding: 1px;\n\t border-style: double double double double ; \n\tborder-width: 2px ;\n\tborder-radius: 0px;\n}\n\nQLineEdit:hover {\n    background-color: #d998ce;\n}\n\nQSpinBox{\n\tpadding: 1px;\n\tborder-style: none none double none ;\n\tborder-width: 2px ;\n\tborder-radius: 0px;\n    border-color: #ae32a0;\n}\n\nQSpinBox:hover {\n    background-color: #d998ce;\n}\n\nQGroupBox {\n  font: medium Ubuntu;\n  font-size: 20px;\n  color: #d998ce;  \n  font-weight: bold;\n  \t border-style: double double double double ; \n\tborder-width: 2px ;\n\tborder-radius: 0px;\n\tborder-color: #ae32a0;\n  padding: 4px;\n  margin-top: 16px;\n  padding: 1px;\n}\n\nQRadioButton {\n    font: medium Ubuntu;\n    font-size: 20px;\n    color:  #ae32a0; \n}")
	this.TabWidget = widgets.NewQTabWidget(this.Centralwidget)
	this.TabWidget.SetObjectName("TabWidget")
	this.TabWidget.SetGeometry(core.NewQRect4(20, 0, 961, 951))
	this.Lab3 = widgets.NewQWidget(this.TabWidget, core.Qt__Widget)
	this.Lab3.SetObjectName("Lab3")
	this.PlainTextEdit3 = widgets.NewQPlainTextEdit(this.Lab3)
	this.PlainTextEdit3.SetObjectName("PlainTextEdit3")
	this.PlainTextEdit3.SetGeometry(core.NewQRect4(680, 610, 221, 241))
	this.PlainTextEdit3.SetReadOnly(true)
	this.PlainTextEdit = widgets.NewQPlainTextEdit(this.Lab3)
	this.PlainTextEdit.SetObjectName("PlainTextEdit")
	this.PlainTextEdit.SetGeometry(core.NewQRect4(680, 90, 221, 241))
	this.PlainTextEdit.SetReadOnly(true)
	this.PlainTextEdit2 = widgets.NewQPlainTextEdit(this.Lab3)
	this.PlainTextEdit2.SetObjectName("PlainTextEdit2")
	this.PlainTextEdit2.SetGeometry(core.NewQRect4(680, 350, 221, 241))
	this.PlainTextEdit2.SetReadOnly(true)
	this.Group3 = widgets.NewQGroupBox(this.Lab3)
	this.Group3.SetObjectName("Group3")
	this.Group3.SetGeometry(core.NewQRect4(10, 600, 651, 251))
	this.Line3 = widgets.NewQLineEdit(this.Group3)
	this.Line3.SetObjectName("Line3")
	this.Line3.SetGeometry(core.NewQRect4(170, 60, 261, 51))
	this.Answer3 = widgets.NewQLineEdit(this.Group3)
	this.Answer3.SetObjectName("Answer3")
	this.Answer3.SetGeometry(core.NewQRect4(170, 140, 261, 51))
	this.Answer3.SetReadOnly(true)
	this.Label3 = widgets.NewQLabel(this.Group3, core.Qt__Widget)
	this.Label3.SetObjectName("Label3")
	this.Label3.SetGeometry(core.NewQRect4(10, 20, 101, 61))
	this.BaseBox = widgets.NewQGroupBox(this.Lab3)
	this.BaseBox.SetObjectName("BaseBox")
	this.BaseBox.SetGeometry(core.NewQRect4(300, -10, 351, 80))
	this.BaseBox.SetStyleSheet("QGroupBox {\n  font: medium Ubuntu;\n  font-size: 20px;\n  color: #d998ce;  \n  font-weight: bold;\n  \t border-style: none none none none ; \n\tborder-width: 2px ;\n\tborder-radius: 0px;\n\tborder-color: #ae32a0;\n  padding: 4px;\n  margin-top: 16px;\n  padding: 1px;\n}")
	this.BinaryRadio = widgets.NewQRadioButton(this.BaseBox)
	this.BinaryRadio.SetObjectName("BinaryRadio")
	this.BinaryRadio.SetGeometry(core.NewQRect4(10, 40, 131, 27))
	this.BinaryRadio.SetChecked(false)
	this.DecRadio = widgets.NewQRadioButton(this.BaseBox)
	this.DecRadio.SetObjectName("DecRadio")
	this.DecRadio.SetGeometry(core.NewQRect4(180, 40, 161, 27))
	this.DecRadio.SetChecked(true)
	this.Group1 = widgets.NewQGroupBox(this.Lab3)
	this.Group1.SetObjectName("Group1")
	this.Group1.SetGeometry(core.NewQRect4(10, 80, 651, 251))
	this.Line31 = widgets.NewQLineEdit(this.Group1)
	this.Line31.SetObjectName("Line31")
	this.Line31.SetGeometry(core.NewQRect4(170, 60, 261, 51))
	this.Answer = widgets.NewQLineEdit(this.Group1)
	this.Answer.SetObjectName("Answer")
	this.Answer.SetGeometry(core.NewQRect4(170, 140, 261, 51))
	this.Answer.SetReadOnly(true)
	this.Label31 = widgets.NewQLabel(this.Group1, core.Qt__Widget)
	this.Label31.SetObjectName("Label31")
	this.Label31.SetGeometry(core.NewQRect4(10, 30, 101, 41))
	this.Group2 = widgets.NewQGroupBox(this.Lab3)
	this.Group2.SetObjectName("Group2")
	this.Group2.SetGeometry(core.NewQRect4(10, 340, 651, 251))
	this.LinePoly2 = widgets.NewQLineEdit(this.Group2)
	this.LinePoly2.SetObjectName("LinePoly2")
	this.LinePoly2.SetGeometry(core.NewQRect4(340, 70, 271, 51))
	this.LinePoly2.SetReadOnly(false)
	this.LinePoly1 = widgets.NewQLineEdit(this.Group2)
	this.LinePoly1.SetObjectName("LinePoly1")
	this.LinePoly1.SetGeometry(core.NewQRect4(20, 70, 261, 51))
	this.Answer2 = widgets.NewQLineEdit(this.Group2)
	this.Answer2.SetObjectName("Answer2")
	this.Answer2.SetGeometry(core.NewQRect4(20, 160, 261, 51))
	this.Answer2.SetReadOnly(true)
	this.Label32 = widgets.NewQLabel(this.Group2, core.Qt__Widget)
	this.Label32.SetObjectName("Label32")
	this.Label32.SetGeometry(core.NewQRect4(10, 20, 101, 41))
	this.RadioPolynom = widgets.NewQRadioButton(this.Group2)
	this.RadioPolynom.SetObjectName("RadioPolynom")
	this.RadioPolynom.SetGeometry(core.NewQRect4(340, 150, 161, 27))
	this.RadioPolynom.SetChecked(false)
	this.RadioElem = widgets.NewQRadioButton(this.Group2)
	this.RadioElem.SetObjectName("RadioElem")
	this.RadioElem.SetGeometry(core.NewQRect4(340, 190, 161, 27))
	this.Lab3LabelMultiply = widgets.NewQLabel(this.Group2, core.Qt__Widget)
	this.Lab3LabelMultiply.SetObjectName("Lab3LabelMultiply")
	this.Lab3LabelMultiply.SetGeometry(core.NewQRect4(290, 80, 41, 41))
	this.TabWidget.AddTab(this.Lab3, "")
	this.Lab2 = widgets.NewQWidget(this.TabWidget, core.Qt__Widget)
	this.Lab2.SetObjectName("Lab2")
	this.LabelGCD = widgets.NewQLabel(this.Lab2, core.Qt__Widget)
	this.LabelGCD.SetObjectName("LabelGCD")
	this.LabelGCD.SetGeometry(core.NewQRect4(40, 30, 361, 51))
	this.GcdA = widgets.NewQLineEdit(this.Lab2)
	this.GcdA.SetObjectName("GcdA")
	this.GcdA.SetGeometry(core.NewQRect4(60, 110, 81, 41))
	this.GcdB = widgets.NewQLineEdit(this.Lab2)
	this.GcdB.SetObjectName("GcdB")
	this.GcdB.SetGeometry(core.NewQRect4(300, 110, 81, 41))
	this.GcdY = widgets.NewQLineEdit(this.Lab2)
	this.GcdY.SetObjectName("GcdY")
	this.GcdY.SetGeometry(core.NewQRect4(390, 110, 81, 41))
	this.GcdY.SetReadOnly(true)
	this.Gcd = widgets.NewQLineEdit(this.Lab2)
	this.Gcd.SetObjectName("Gcd")
	this.Gcd.SetGeometry(core.NewQRect4(550, 110, 81, 41))
	this.Gcd.SetReadOnly(true)
	this.GcdX = widgets.NewQLineEdit(this.Lab2)
	this.GcdX.SetObjectName("GcdX")
	this.GcdX.SetGeometry(core.NewQRect4(150, 110, 81, 41))
	this.GcdX.SetReadOnly(true)
	this.LabelGcdPlus = widgets.NewQLabel(this.Lab2, core.Qt__Widget)
	this.LabelGcdPlus.SetObjectName("LabelGcdPlus")
	this.LabelGcdPlus.SetGeometry(core.NewQRect4(240, 110, 41, 41))
	this.LabelGcdEq = widgets.NewQLabel(this.Lab2, core.Qt__Widget)
	this.LabelGcdEq.SetObjectName("LabelGcdEq")
	this.LabelGcdEq.SetGeometry(core.NewQRect4(490, 110, 41, 41))
	this.BinPowLabel = widgets.NewQLabel(this.Lab2, core.Qt__Widget)
	this.BinPowLabel.SetObjectName("BinPowLabel")
	this.BinPowLabel.SetGeometry(core.NewQRect4(50, 210, 361, 51))
	this.BinPowA = widgets.NewQLineEdit(this.Lab2)
	this.BinPowA.SetObjectName("BinPowA")
	this.BinPowA.SetGeometry(core.NewQRect4(70, 290, 81, 41))
	this.BinPowB = widgets.NewQLineEdit(this.Lab2)
	this.BinPowB.SetObjectName("BinPowB")
	this.BinPowB.SetGeometry(core.NewQRect4(160, 260, 51, 31))
	this.LabelBinPowEq = widgets.NewQLabel(this.Lab2, core.Qt__Widget)
	this.LabelBinPowEq.SetObjectName("LabelBinPowEq")
	this.LabelBinPowEq.SetGeometry(core.NewQRect4(210, 290, 41, 41))
	this.BinPow = widgets.NewQLineEdit(this.Lab2)
	this.BinPow.SetObjectName("BinPow")
	this.BinPow.SetGeometry(core.NewQRect4(270, 290, 81, 41))
	this.BinPow.SetReadOnly(true)
	this.LabelLab2Modulo = widgets.NewQLabel(this.Lab2, core.Qt__Widget)
	this.LabelLab2Modulo.SetObjectName("LabelLab2Modulo")
	this.LabelLab2Modulo.SetGeometry(core.NewQRect4(760, 210, 101, 51))
	this.Lab2Modulo = widgets.NewQLineEdit(this.Lab2)
	this.Lab2Modulo.SetObjectName("Lab2Modulo")
	this.Lab2Modulo.SetGeometry(core.NewQRect4(770, 290, 81, 41))
	this.Lab2Modulo.SetReadOnly(true)
	this.TabWidget.AddTab(this.Lab2, "")
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Menubar = widgets.NewQMenuBar(MainWindow)
	this.Menubar.SetObjectName("Menubar")
	this.Menubar.SetGeometry(core.NewQRect4(0, 0, 1024, 28))
	MainWindow.SetMenuBar(this.Menubar)
	this.Statusbar = widgets.NewQStatusBar(MainWindow)
	this.Statusbar.SetObjectName("Statusbar")
	MainWindow.SetStatusBar(this.Statusbar)

	this.RetranslateUi(MainWindow)
	this.TabWidget.SetCurrentIndex(1)
}

func (this *UICryptMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
	_translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "Lab", "", -1))
	this.PlainTextEdit3.SetPlainText(_translate("MainWindow", "Напишите функцию для поиска мультипликативного обратного для элемента из 𝐺𝐹(256).", "", -1))
	this.PlainTextEdit.SetPlainText(_translate("MainWindow", "Напишите функцию, представляющую элемент из 𝐺𝐹(256) в полиномиальной форме.", "", -1))
	this.PlainTextEdit2.SetPlainText(_translate("MainWindow", "Напишите функцию\n    - умножения двух двоичных многочленов; \n    - умножения двух элементов из 𝐺𝐹(256).", "", -1))
	this.Group3.SetTitle(_translate("MainWindow", "", "", -1))
	this.Line3.SetInputMask(_translate("MainWindow", "", "", -1))
	this.Label3.SetText(_translate("MainWindow", "3", "", -1))
	this.BaseBox.SetTitle(_translate("MainWindow", "", "", -1))
	this.BinaryRadio.SetText(_translate("MainWindow", "Двоичная", "", -1))
	this.DecRadio.SetText(_translate("MainWindow", "Десятичная", "", -1))
	this.Group1.SetTitle(_translate("MainWindow", "", "", -1))
	this.Line31.SetInputMask(_translate("MainWindow", "", "", -1))
	this.Label31.SetText(_translate("MainWindow", "1", "", -1))
	this.Group2.SetTitle(_translate("MainWindow", "", "", -1))
	this.LinePoly2.SetInputMask(_translate("MainWindow", "", "", -1))
	this.LinePoly1.SetInputMask(_translate("MainWindow", "", "", -1))
	this.Label32.SetText(_translate("MainWindow", "2", "", -1))
	this.RadioPolynom.SetText(_translate("MainWindow", "Многочлены", "", -1))
	this.RadioElem.SetText(_translate("MainWindow", "Элементы", "", -1))
	this.Lab3LabelMultiply.SetText(_translate("MainWindow", "*", "", -1))
	this.TabWidget.SetTabText(this.TabWidget.IndexOf(this.Lab3), _translate("MainWindow", "Tab 1", "", -1))
	this.LabelGCD.SetText(_translate("MainWindow", "Расширенный алгоритм Евклида", "", -1))
	this.GcdY.SetText(_translate("MainWindow", "", "", -1))
	this.LabelGcdPlus.SetText(_translate("MainWindow", "+", "", -1))
	this.LabelGcdEq.SetText(_translate("MainWindow", "=", "", -1))
	this.BinPowLabel.SetText(_translate("MainWindow", "Бинарное возведение в степень", "", -1))
	this.BinPowA.SetText(_translate("MainWindow", "", "", -1))
	this.BinPowB.SetText(_translate("MainWindow", "", "", -1))
	this.LabelBinPowEq.SetText(_translate("MainWindow", "=", "", -1))
	this.BinPow.SetText(_translate("MainWindow", "", "", -1))
	this.LabelLab2Modulo.SetText(_translate("MainWindow", "Модуль", "", -1))
	this.Lab2Modulo.SetText(_translate("MainWindow", "256", "", -1))
	this.TabWidget.SetTabText(this.TabWidget.IndexOf(this.Lab2), _translate("MainWindow", "Tab 2", "", -1))
}
