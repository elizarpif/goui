// WARNING! All changes made in this file will be lost!
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UICryptMainWindow struct {
	Centralwidget *widgets.QWidget
	Group2 *widgets.QGroupBox
	LinePoly2 *widgets.QLineEdit
	Label2 *widgets.QLabel
	LinePoly1 *widgets.QLineEdit
	Answer2 *widgets.QLineEdit
	Label32 *widgets.QLabel
	RadioPolynom *widgets.QRadioButton
	RadioElem *widgets.QRadioButton
	Multiply *widgets.QLabel
	Group1 *widgets.QGroupBox
	Line1 *widgets.QLineEdit
	Answer *widgets.QLineEdit
	Label31 *widgets.QLabel
	Group3 *widgets.QGroupBox
	Line3 *widgets.QLineEdit
	Answer3 *widgets.QLineEdit
	Label33 *widgets.QLabel
	PlainTextEdit *widgets.QPlainTextEdit
	PlainTextEdit2 *widgets.QPlainTextEdit
	PlainTextEdit3 *widgets.QPlainTextEdit
	LabelLab *widgets.QLabel
	Menubar *widgets.QMenuBar
	Statusbar *widgets.QStatusBar
}

func (this *UICryptMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 951, 958))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.Centralwidget.SetStyleSheet("QWidget{\nbackground-color: white;\n}\n\nQFrame {\n\tpadding: 10px;\n    border-radius: 10px;        \n    border-width: 5px;\n    border-color: #ae32a0;\n    border-style: double;\n\tbackground-color: white;\n    font: medium Ubuntu;\n    font-size: 20px;\n} \n\n\nQLabel {\n    font: medium Ubuntu;\n    font-size: 20px;\n    color: #ae32a0;\n    border-width: 0px;\n}        \n\nQPushButton {\n    background-color: rgb(47, 155, 255);\n    color: white;\n      font: medium Ubuntu;\n    font-size: 20px;\n    border-radius: 20px;        \n    border-width: 1px;\n    border-color: rgb(172, 194, 255);\n    border-style: solid;\n}\nQPushButton:hover {\n    background-color: rgb(255, 170, 0);\n}\nQPushButton:pressed {\n    background-color: rgb(255, 153, 29);\n}    \n\nQLineEdit {\n\tborder-color: #d998ce;  \n\tfont: medium Ubuntu;\n    font-size: 20px;\n\tpadding: 1px;\n\t border-style: double double double double ; \n\tborder-width: 2px ;\n\tborder-radius: 0px;\n}\n\nQLineEdit:hover {\n    background-color: #d998ce;\n}\n\nQSpinBox{\n\tpadding: 1px;\n\tborder-style: none none double none ;\n\tborder-width: 2px ;\n\tborder-radius: 0px;\n    border-color: #ae32a0;\n}\n\nQSpinBox:hover {\n    background-color: #d998ce;\n}\n\nQGroupBox {\n  font: medium Ubuntu;\n  font-size: 20px;\n  color: #d998ce;  \n  font-weight: bold;\n  \t border-style: double double double double ; \n\tborder-width: 2px ;\n\tborder-radius: 0px;\n\tborder-color: #ae32a0;\n  padding: 4px;\n  margin-top: 16px;\n  padding: 1px;\n}\n\nQRadioButton {\n    font: medium Ubuntu;\n    font-size: 20px;\n    color:  #ae32a0; \n}")
	this.Group2 = widgets.NewQGroupBox(this.Centralwidget)
	this.Group2.SetObjectName("Group2")
	this.Group2.SetGeometry(core.NewQRect4(30, 360, 651, 251))
	this.LinePoly2 = widgets.NewQLineEdit(this.Group2)
	this.LinePoly2.SetObjectName("LinePoly2")
	this.LinePoly2.SetGeometry(core.NewQRect4(340, 70, 271, 51))
	this.LinePoly2.SetReadOnly(false)
	this.Label2 = widgets.NewQLabel(this.Group2, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.Label2.SetGeometry(core.NewQRect4(310, 70, 16, 51))
	this.Label2.SetStyleSheet("")
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
	this.RadioPolynom.SetChecked(true)
	this.RadioElem = widgets.NewQRadioButton(this.Group2)
	this.RadioElem.SetObjectName("RadioElem")
	this.RadioElem.SetGeometry(core.NewQRect4(340, 190, 161, 27))
	this.Multiply = widgets.NewQLabel(this.Group2, core.Qt__Widget)
	this.Multiply.SetObjectName("Multiply")
	this.Multiply.SetGeometry(core.NewQRect4(290, 80, 41, 41))
	this.Group1 = widgets.NewQGroupBox(this.Centralwidget)
	this.Group1.SetObjectName("Group1")
	this.Group1.SetGeometry(core.NewQRect4(30, 100, 651, 251))
	this.Line1 = widgets.NewQLineEdit(this.Group1)
	this.Line1.SetObjectName("Line1")
	this.Line1.SetGeometry(core.NewQRect4(170, 60, 261, 51))
	this.Answer = widgets.NewQLineEdit(this.Group1)
	this.Answer.SetObjectName("Answer")
	this.Answer.SetGeometry(core.NewQRect4(170, 140, 261, 51))
	this.Answer.SetReadOnly(true)
	this.Label31 = widgets.NewQLabel(this.Group1, core.Qt__Widget)
	this.Label31.SetObjectName("Label31")
	this.Label31.SetGeometry(core.NewQRect4(10, 30, 101, 41))
	this.Group3 = widgets.NewQGroupBox(this.Centralwidget)
	this.Group3.SetObjectName("Group3")
	this.Group3.SetGeometry(core.NewQRect4(30, 620, 651, 251))
	this.Line3 = widgets.NewQLineEdit(this.Group3)
	this.Line3.SetObjectName("Line3")
	this.Line3.SetGeometry(core.NewQRect4(170, 60, 261, 51))
	this.Answer3 = widgets.NewQLineEdit(this.Group3)
	this.Answer3.SetObjectName("Answer3")
	this.Answer3.SetGeometry(core.NewQRect4(170, 140, 261, 51))
	this.Answer3.SetReadOnly(true)
	this.Label33 = widgets.NewQLabel(this.Group3, core.Qt__Widget)
	this.Label33.SetObjectName("Label33")
	this.Label33.SetGeometry(core.NewQRect4(10, 20, 101, 61))
	this.PlainTextEdit = widgets.NewQPlainTextEdit(this.Centralwidget)
	this.PlainTextEdit.SetObjectName("PlainTextEdit")
	this.PlainTextEdit.SetGeometry(core.NewQRect4(700, 110, 221, 241))
	this.PlainTextEdit.SetReadOnly(true)
	this.PlainTextEdit2 = widgets.NewQPlainTextEdit(this.Centralwidget)
	this.PlainTextEdit2.SetObjectName("PlainTextEdit2")
	this.PlainTextEdit2.SetGeometry(core.NewQRect4(700, 370, 221, 241))
	this.PlainTextEdit2.SetReadOnly(true)
	this.PlainTextEdit3 = widgets.NewQPlainTextEdit(this.Centralwidget)
	this.PlainTextEdit3.SetObjectName("PlainTextEdit3")
	this.PlainTextEdit3.SetGeometry(core.NewQRect4(700, 630, 221, 241))
	this.PlainTextEdit3.SetReadOnly(true)
	this.LabelLab = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.LabelLab.SetObjectName("LabelLab")
	this.LabelLab.SetGeometry(core.NewQRect4(430, 40, 251, 61))
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Menubar = widgets.NewQMenuBar(MainWindow)
	this.Menubar.SetObjectName("Menubar")
	this.Menubar.SetGeometry(core.NewQRect4(0, 0, 951, 28))
	MainWindow.SetMenuBar(this.Menubar)
	this.Statusbar = widgets.NewQStatusBar(MainWindow)
	this.Statusbar.SetObjectName("Statusbar")
	MainWindow.SetStatusBar(this.Statusbar)


    this.RetranslateUi(MainWindow)

}

func (this *UICryptMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
    _translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "Lab", "", -1))
	this.Group2.SetTitle(_translate("MainWindow", "", "", -1))
	this.LinePoly2.SetInputMask(_translate("MainWindow", "", "", -1))
	this.Label2.SetText(_translate("MainWindow", "*", "", -1))
	this.LinePoly1.SetInputMask(_translate("MainWindow", "", "", -1))
	this.Label32.SetText(_translate("MainWindow", "2", "", -1))
	this.RadioPolynom.SetText(_translate("MainWindow", "Многочлены", "", -1))
	this.RadioElem.SetText(_translate("MainWindow", "Элементы", "", -1))
	this.Multiply.SetText(_translate("MainWindow", "*", "", -1))
	this.Group1.SetTitle(_translate("MainWindow", "", "", -1))
	this.Line1.SetInputMask(_translate("MainWindow", "", "", -1))
	this.Label31.SetText(_translate("MainWindow", "1", "", -1))
	this.Group3.SetTitle(_translate("MainWindow", "", "", -1))
	this.Line3.SetInputMask(_translate("MainWindow", "", "", -1))
	this.Label33.SetText(_translate("MainWindow", "3", "", -1))
	this.PlainTextEdit.SetPlainText(_translate("MainWindow", "Напишите функцию, представляющую элемент из 𝐺𝐹(256) в полиномиальной форме.", "", -1))
	this.PlainTextEdit2.SetPlainText(_translate("MainWindow", "Напишите функцию\n    - умножения двух двоичных многочленов; \n    - умножения двух элементов из 𝐺𝐹(256).", "", -1))
	this.PlainTextEdit3.SetPlainText(_translate("MainWindow", "Напишите функцию для поиска мультипликативного обратного для элемента из 𝐺𝐹(256).", "", -1))
	this.LabelLab.SetText(_translate("MainWindow", "Lab3", "", -1))
}
