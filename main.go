package main

import (
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("My Own system")

var WeatherApp fyne.Widget 
var calculator fyne.Widget 
var galleryApp fyne.Widget 
var notesApp fyne.Widget 

var img fyne.CanvasObject

var DeskBtn fyne.Widget
var NewsBtn fyne.Widget
var ModeBtn fyne.Widget
var songsApp fyne.Widget
var panelContent *fyne.Container
var quitBtn fyne.Widget
var bgChangeBtn fyne.Widget

func main(){ 
	version := false
	bgPic := 1
	myApp.Settings().SetTheme(theme.DarkTheme())
	img = canvas.NewImageFromFile("D:\\Development Pepcoding\\Freelancing-1\\virtual device using go lang & fyne\\images\\wallpaper1.jpeg")


	datetime := canvas.NewText(getTime(), color.White)
	datetime.TextSize = 15
	datetime.Alignment = fyne.TextAlignCenter
	datetime.TextStyle = fyne.TextStyle{Bold: true}


	go func() {
		for {
			datetime.Text = getTime()
			datetime.Refresh()
		}
	}()


	imgk, _ := http.Get("https://w7.pngwing.com/pngs/100/239/png-transparent-arrow-computer-icons-symbol-next-button-angle-text-triangle.png")
	bodyk, _ := ioutil.ReadAll(imgk.Body)
	iconk := canvas.NewImageFromResource(fyne.NewStaticResource("", bodyk))
	iconk.FillMode = canvas.ImageFillContain

	
	

	img5, _ := http.Get("https://user-images.githubusercontent.com/89251393/140250230-cdddb4bd-8403-4b1a-a586-78bc3b9ad24d.png")
	body5, _ := ioutil.ReadAll(img5.Body)
	icon4 := canvas.NewImageFromResource(fyne.NewStaticResource("", body5))
	icon4.FillMode = canvas.ImageFillContain
	WeatherApp = widget.NewButton("", func(){
		showWeatherApp()
	})

	img2, _ := http.Get("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRdYX0ixRZsGv5sC3353tyQ_DV_piqG8zeRRQ&usqp=CAU")
	body1, _ := ioutil.ReadAll(img2.Body)
	icon1 := canvas.NewImageFromResource(fyne.NewStaticResource("", body1))
	icon1.FillMode = canvas.ImageFillContain
	calculator = widget.NewButton("", func(){
		showCalc()
	})




	img99, _ := http.Get("https://lh3.googleusercontent.com/Iip-8Yn3PLAzecCMb4ZaHTvFObl3ETUWZmd5zLflhbB6BXKyNc5aM4hrGAA9NXSs7i0")
	body99, _ := ioutil.ReadAll(img99.Body)
	icon99 := canvas.NewImageFromResource(fyne.NewStaticResource("", body99))
	icon99.FillMode = canvas.ImageFillContain	
	NewsBtn = widget.NewButton("",func() {
		newsApp()
	})



	img88, _ := http.Get("https://developer.apple.com/design/human-interface-guidelines/macos/images/app-icon-consistent-shape_2x.png")
	body88, _ := ioutil.ReadAll(img88.Body)
	icon88 := canvas.NewImageFromResource(fyne.NewStaticResource("", body88))
	icon88.FillMode = canvas.ImageFillContain	
	ModeBtn = widget.NewButton("",func() {
		version = !version
		if !version{
			myApp.Settings().SetTheme(theme.DarkTheme())
		}else{
			myApp.Settings().SetTheme(theme.LightTheme())
		}
	})


	img4, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249556-e51b6917-2805-4086-8d31-e70eeac4c708.png")
	body4, _ := ioutil.ReadAll(img4.Body)
	icon3 := canvas.NewImageFromResource(fyne.NewStaticResource("", body4))
	icon3.FillMode = canvas.ImageFillContain
	galleryApp = widget.NewButton("", func(){
		showGalleryApp()
	})



	img3, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249614-adc807e1-0d93-417f-bb13-b8903c12ee9d.png")
	body3, _ := ioutil.ReadAll(img3.Body)
	icon2 := canvas.NewImageFromResource(fyne.NewStaticResource("", body3))
	icon2.FillMode = canvas.ImageFillContain
	notesApp = widget.NewButton("", func(){
		showTextEditor()
	})


	img7, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249608-9c1f5282-3172-42ca-bb8c-c5f7d7a45730.png")
	body7, _ := ioutil.ReadAll(img7.Body)
	icon7 := canvas.NewImageFromResource(fyne.NewStaticResource("", body7))
	icon7.FillMode = canvas.ImageFillContain
	songsApp := widget.NewButton("", func(){
		musicApp()
	}) 


	img9, _ := http.Get("https://user-images.githubusercontent.com/89251393/140249619-cb66d5e9-0f2e-439c-a222-431cba8e5cac.png")
	body9, _ := ioutil.ReadAll(img9.Body)
	icon6 := canvas.NewImageFromResource(fyne.NewStaticResource("", body9))
	icon6.FillMode = canvas.ImageFillContain
	quitBtn = widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		myWindow.Close()
	})


	
	img8, _ := http.Get("https://play-lh.googleusercontent.com/cFMbr4zwNMhbxXzDcq8wR7k86B2LNYD_lV9nekYDgh74ANdOMQb1b594a1Hhltccfw")
	body8, _ := ioutil.ReadAll(img8.Body)
	icon8 := canvas.NewImageFromResource(fyne.NewStaticResource("", body8))
	icon8.FillMode = canvas.ImageFillContain
	DeskBtn = widget.NewButtonWithIcon("",theme.HomeIcon(), func(){
		temp := container.New(layout.NewHBoxLayout(),
	layout.NewSpacer(),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(DeskBtn,icon8)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(WeatherApp,icon4)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(calculator,icon1)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(songsApp,icon7)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(galleryApp,icon3)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(notesApp,icon2)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(NewsBtn,icon99)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(ModeBtn,icon88)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(quitBtn,icon6)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(bgChangeBtn,iconk)),
	layout.NewSpacer(),
)
		myWindow.SetContent(
			container.New(layout.NewMaxLayout(), img,
				container.New(layout.NewBorderLayout(datetime, temp, nil, nil), datetime, temp),
			),
		)
	})


	bgChangeBtn = widget.NewButton("", func() {
		temp2 := container.New(layout.NewHBoxLayout(),
	layout.NewSpacer(),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(DeskBtn,icon8)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(WeatherApp,icon4)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(calculator,icon1)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(songsApp,icon7)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(galleryApp,icon3)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(notesApp,icon2)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(NewsBtn,icon99)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(ModeBtn,icon88)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(quitBtn,icon6)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(bgChangeBtn,iconk)),
	layout.NewSpacer(),
)
		

		if bgPic==5{
			bgPic=1
		}else{
			bgPic++
		}

		if bgPic==1{
			img = canvas.NewImageFromFile("D:\\Development Pepcoding\\Freelancing-1\\virtual device using go lang & fyne\\images\\wallpaper1.jpeg")
		}else if bgPic==2{
			img = canvas.NewImageFromFile("D:\\Development Pepcoding\\Freelancing-1\\virtual device using go lang & fyne\\image2.jpg")
		}else if bgPic==3{
			img = canvas.NewImageFromFile("D:\\Development Pepcoding\\Freelancing-1\\virtual device using go lang & fyne\\images\\wallpaper2.jpeg")
		}else if bgPic==4{
			img = canvas.NewImageFromFile("D:\\Development Pepcoding\\Freelancing-1\\virtual device using go lang & fyne\\images\\wallpaper3.jpg")
		}else{
			img = canvas.NewImageFromFile("D:\\Development Pepcoding\\Freelancing-1\\virtual device using go lang & fyne\\images\\wallpaper4.jpeg")
		}

		myWindow.SetContent(
			container.New(layout.NewMaxLayout(), img,
				container.New(layout.NewBorderLayout(datetime, temp2, nil, nil), datetime, temp2),
			),
		)
	})
	// WeatherApp.Size(fyne.NewSize(12,11))
	// WeatherApp.Move(fyne.NewPos(780,700))
	// calculator.Move(fyne.NewPos(540,300))
	// DeskBtn.Move(fyne.NewPos(540,300))
	// musicApp.Move(fyne.NewPos(540,320))
	// galleryApp.Move(fyne.NewPos(540,301))
	// btn4.Move(fyne.NewPos(101,301))
	// ModeBtn.Move(fyne.NewPos(240,111))
	// NewsBtn.Move(fyne.NewPos(540,301))
	// panelContent = container.NewVBox((container.NewGridWithColumns(8,DeskBtn,WeatherApp,calculator,galleryApp,NewsBtn,btn4,musicApp,ModeBtn)))
	
	apps := container.New(layout.NewHBoxLayout(),
	layout.NewSpacer(),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(DeskBtn,icon8)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(WeatherApp,icon4)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(calculator,icon1)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(songsApp,icon7)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(galleryApp,icon3)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(notesApp,icon2)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(NewsBtn,icon99)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(ModeBtn,icon88)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(quitBtn,icon6)),
	container.NewGridWrap(fyne.NewSize(50,50), container.NewPadded(bgChangeBtn,iconk)),
	layout.NewSpacer(),
)


	
	myWindow.Resize(fyne.NewSize(1280,720))
	myWindow.CenterOnScreen()
	r, _ := fyne.LoadResourceFromPath("images\\oslogo.png")
	myWindow.SetIcon(r)
	myWindow.SetContent(
		container.New(layout.NewMaxLayout(), img,
			container.New(layout.NewBorderLayout(datetime, apps, nil, nil), datetime, apps),
		),
	)
	k, _:=fyne.LoadResourceFromPath("images\\oslogo.png")
	myWindow.SetIcon(k)
	myWindow.ShowAndRun()
}