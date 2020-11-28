// WARNING! All changes made in this file will be lost!
package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UICryptMainWindow struct {
	Centralwidget *widgets.QWidget
	Label *widgets.QLabel
	LineEdit *widgets.QLineEdit
	Btn *widgets.QPushButton
	AnswerLine *widgets.QLineEdit
	AnswerLb *widgets.QLabel
	GroupBox *widgets.QGroupBox
	Lab31 *widgets.QRadioButton
	Lab32 *widgets.QRadioButton
	Lab33 *widgets.QRadioButton
	Lab44 *widgets.QRadioButton
	Menubar *widgets.QMenuBar
	Statusbar *widgets.QStatusBar
}

func (this *UICryptMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 951, 505))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.Centralwidget.SetStyleSheet("QWidget{\nbackground-color: white;\n}\n\nQFrame {\n  border-radius: 55px;        \n    border-width: 5px;\n    border-color: #ae32a0;\n    border-style: double;\nbackground-color: white;\n} \n\n\nQLabel {\n    font: medium Ubuntu;\n    font-size: 20px;\n    color: #d998ce;  \n    border-width: 0px;\nbackground-image: url(:/bktop_Qt_5_3_MinGW_32bit-Debug/debug/candy.jpg);\n}        \n\nQPushButton {\n    background-color: rgb(47, 155, 255);\n    color: white;\n  \n    border-radius: 20px;        \n    border-width: 1px;\n    border-color: rgb(172, 194, 255);\n    border-style: solid;\n}\nQPushButton:hover {\n    background-color: rgb(255, 170, 0);\n}\nQPushButton:pressed {\n    background-color: rgb(255, 153, 29);\n}    \n\nQLineEdit {\n\tfont: medium Ubuntu;\n    font-size: 20px;\n\tpadding: 1px;\n\tborder-style: none none double none ;\n\tborder-width: 2px ;\n\tborder-radius: 0px;\n\tborder-color: #ae32a0;\n\tbackground-image: url(:/bktop_Qt_5_3_MinGW_32bit-Debug/debug/candy.jpg);\n}\n\nQLineEdit:hover {\n    background-color: #d998ce;\nbackground-image: url(:/bktop_Qt_5_3_MinGW_32bit-Debug/debug/candy.jpg);\n}\n\nQSpinBox{\npadding: 1px;\nborder-style: none none double none ;\nborder-width: 2px ;\nborder-radius: 0px;\n border-color: #ae32a0;\n}\n\nQSpinBox:hover {\n    background-color: #d998ce;\n}\n\nQGroupBox {\n  font: medium Ubuntu;\n  font-size: 20px;\n  color: #d998ce;  \n  font-weight: bold;\n\n  padding: 4px;\n  margin-top: 16px;\n  padding: 1px;\n\n}\n\nQRadioButton {\n    font: medium Ubuntu;\n    font-size: 20px;\n    color:  #ae32a0; \n}")
	this.Label = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.Label.SetGeometry(core.NewQRect4(20, 70, 221, 51))
	this.Label.SetStyleSheet("")
	this.LineEdit = widgets.NewQLineEdit(this.Centralwidget)
	this.LineEdit.SetObjectName("LineEdit")
	this.LineEdit.SetGeometry(core.NewQRect4(20, 140, 431, 51))
	this.Btn = widgets.NewQPushButton(this.Centralwidget)
	this.Btn.SetObjectName("Btn")
	this.Btn.SetGeometry(core.NewQRect4(490, 140, 122, 51))
	this.AnswerLine = widgets.NewQLineEdit(this.Centralwidget)
	this.AnswerLine.SetObjectName("AnswerLine")
	this.AnswerLine.SetGeometry(core.NewQRect4(20, 280, 431, 51))
	this.AnswerLine.SetReadOnly(true)
	this.AnswerLb = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.AnswerLb.SetObjectName("AnswerLb")
	this.AnswerLb.SetGeometry(core.NewQRect4(20, 210, 221, 51))
	this.GroupBox = widgets.NewQGroupBox(this.Centralwidget)
	this.GroupBox.SetObjectName("GroupBox")
	this.GroupBox.SetGeometry(core.NewQRect4(800, 20, 120, 261))
	this.Lab31 = widgets.NewQRadioButton(this.GroupBox)
	this.Lab31.SetObjectName("Lab31")
	this.Lab31.SetGeometry(core.NewQRect4(10, 60, 81, 27))
	this.Lab31.SetChecked(true)
	this.Lab32 = widgets.NewQRadioButton(this.GroupBox)
	this.Lab32.SetObjectName("Lab32")
	this.Lab32.SetGeometry(core.NewQRect4(10, 100, 81, 27))
	this.Lab33 = widgets.NewQRadioButton(this.GroupBox)
	this.Lab33.SetObjectName("Lab33")
	this.Lab33.SetGeometry(core.NewQRect4(10, 140, 81, 27))
	this.Lab44 = widgets.NewQRadioButton(this.GroupBox)
	this.Lab44.SetObjectName("Lab44")
	this.Lab44.SetGeometry(core.NewQRect4(10, 180, 81, 27))
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
	MainWindow.SetWindowTitle(_translate("MainWindow", "MainWindow", "", -1))
	this.Label.SetText(_translate("MainWindow", "Enter", "", -1))
	this.Btn.SetText(_translate("MainWindow", "ok", "", -1))
	this.AnswerLb.SetText(_translate("MainWindow", "Answer", "", -1))
	this.GroupBox.SetTitle(_translate("MainWindow", "Lab3", "", -1))
	this.Lab31.SetText(_translate("MainWindow", "1", "", -1))
	this.Lab32.SetText(_translate("MainWindow", "2", "", -1))
	this.Lab33.SetText(_translate("MainWindow", "3", "", -1))
	this.Lab44.SetText(_translate("MainWindow", "4", "", -1))
}
