package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)

	app.Load(core.QUrl_FromLocalFile("./main/qml/application.qml"))

	gui.QGuiApplication_Exec()
}