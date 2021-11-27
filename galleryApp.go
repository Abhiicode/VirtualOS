package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"

	// "fyne.io/fyne/v2/app"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/dialog"
)

func showGalleryApp() {
	// a := app.New()
	// dialog.NewFolderOpen()
	w := myApp.NewWindow("ImageViewer")
	w.Resize(fyne.NewSize(800,600))
	var root_src string;
	selectFolderButton := widget.NewButton("Select target folder", func() {
		dialog.NewFolderOpen(func(dir fyne.ListableURI, err error) {
			if err == nil && dir.Name() != "" {
				fmt.Println(dir.Name())
			}
			root_src=dir.Path()
			files, err := ioutil.ReadDir(root_src)
			if err != nil {
				log.Fatal(err)
			}
			tabs:= container.NewAppTabs()
			for _, file := range files {
				fmt.Println(file.Name(), file.IsDir())
				if !file.IsDir(){
					extensions:= strings.Split(file.Name(),".")[1];
					if extensions == "png" || extensions == "jpeg" || extensions == "jpg"{
						image:= canvas.NewImageFromFile(root_src+"\\"+file.Name())
						image.FillMode = canvas.ImageFillContain
						tabs.Append(container.NewTabItem(file.Name(),image))
					}
				}
			}
			tabs.SetTabLocation(container.TabLocationLeading)
			w.SetContent(container.NewBorder(nil,nil,nil,nil,tabs))
		}, w).Show()
		
	})
	

	
	
	r, _ := fyne.LoadResourceFromPath("images\\gallery2.png")
	w.SetIcon(r)
	w.SetContent(container.NewBorder(selectFolderButton,nil,nil,nil))
	w.Show();
}

