// WARNING! All changes made in this file will be lost!
package uigen

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
	Menubar *widgets.QMenuBar
	Statusbar *widgets.QStatusBar
}

func (this *UICryptMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 651, 505))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.Label = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.Label.SetGeometry(core.NewQRect4(20, 70, 221, 51))
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
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Menubar = widgets.NewQMenuBar(MainWindow)
	this.Menubar.SetObjectName("Menubar")
	this.Menubar.SetGeometry(core.NewQRect4(0, 0, 651, 28))
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
}
