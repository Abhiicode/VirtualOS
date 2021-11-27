package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)
var count int = 1
func showTextEditor(){
	w := myApp.NewWindow("Abhi NotePad")
	w.Resize(fyne.NewSize(500, 500))
	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Text Editor "),
		),
	)
	content.Add(widget.NewButton("Add New File", func(){
		content.Add(widget.NewLabel("New File"+strconv.Itoa(count)))
		count++
	}))
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter Text....")
	input.Resize(fyne.NewSize(400,400))


	saveBtn := widget.NewButton("Save Text File", func(){
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData:= []byte(input.Text)
				uc.Write(textData)
			},w)
		if count==1{
			saveFileDialog.SetFileName("New File"+strconv.Itoa(count)+".txt")
		}else{
			saveFileDialog.SetFileName("New File"+strconv.Itoa(count-1)+".txt")
		}
		
		saveFileDialog.Show()
	})

	openBtn := widget.NewButton("Open Text File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w:= fyne.CurrentApp().NewWindow(
					string(output.StaticName)) 
				w.SetContent(
					container.NewScroll(viewData),
					//need to fix this
					// container.NewGridWrap(saveBtn)
				)
				w.Resize(fyne.NewSize(400,400))
				
				w.Show()
			},w)

			openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
			openFileDialog.Show()
	})
	textEditorContainer := container.NewVBox(
		content,
		input,
		container.NewHBox(
			saveBtn,
			openBtn,
		),
	)
	r, _ := fyne.LoadResourceFromPath("images\\texteditor.png")
	w.SetIcon(r)
	w.CenterOnScreen()
	w.SetContent(container.NewBorder(nil,nil,nil,nil,textEditorContainer),)
	w.Show()
}