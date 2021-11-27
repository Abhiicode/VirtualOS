package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {
	// a := app.New()
	// w := myapp.NewWindow("Calculator")
	// w.Resize(fyne.NewSize(450,280))
	output:=""
	input := widget.NewLabel(output)
	isHistory:=false;
	historyStr := ""
	history:=widget.NewLabel(historyStr);
	var historyArr []string;
	historyBtn:=widget.NewButton("History",func(){
		if isHistory{
			historyStr="";
		}else{
			for i := len(historyArr)-1; i >= 0; i-- {
				historyStr = historyStr+historyArr[i];
				historyStr+="\n";
			}
		}
		isHistory=!isHistory;
		history.SetText((historyStr));
	})
	backBtn:=widget.NewButton("Back",func(){
		if len(output)>0{
			output=output[:len(output)-1];
			input.SetText(output);
		}
	})
	clearBtn:=widget.NewButton("Clear",func(){
		output="";
		input.SetText(output);
	})
	openBtn:=widget.NewButton("(",func(){
		output+="(";
		input.SetText(output);
	})
	closeBtn:=widget.NewButton(")",func(){
		output+=")";
		input.SetText(output);
	})
	divideBtn:=widget.NewButton("/",func(){
		output+="/";
		input.SetText(output);
	})
	sevenBtn:=widget.NewButton("7",func(){
		output+="7";
		input.SetText(output);
	})
	eightBtn:=widget.NewButton("8",func(){
		output+="8";
		input.SetText(output);
	})
	nineBtn:=widget.NewButton("9",func(){
		output+="9";
		input.SetText(output);
	})
	multiplyBtn:=widget.NewButton("*",func(){
		output+="*";
		input.SetText(output);
	})
	fourBtn:=widget.NewButton("4",func(){
		output+="4";
		input.SetText(output);
	})
	fiveBtn:=widget.NewButton("5",func(){
		output+="5";
		input.SetText(output);
	})
	sixBtn:=widget.NewButton("6",func(){
		output+="6";
		input.SetText(output);
	})
	minusBtn:=widget.NewButton("-",func(){
		output+="-";
		input.SetText(output);
	})
	oneBtn:=widget.NewButton("1",func(){
		output+="1";
		input.SetText(output);
	})
	twoBtn:=widget.NewButton("2",func(){
		output+="2";
		input.SetText(output);
	})
	threeBtn:=widget.NewButton("3",func(){
		output+="3";
		input.SetText(output);
	})
	plusBtn:=widget.NewButton("+",func(){
		output+="+";
		input.SetText(output);
	})
	zeroBtn:=widget.NewButton("0",func(){
		output+="0";
		input.SetText(output);
	})
	dotBtn:=widget.NewButton(".",func(){
		output+=".";
		input.SetText(output);
	})
	equaltoBtn:=widget.NewButton("=",func(){
		// output+="=";
		// input.SetText(output);
		expression, err:= govaluate.NewEvaluableExpression(output);
		if err == nil{
			result, err := expression.Evaluate(nil);
			if err == nil{
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64);
				strToAppend:= output+" = "+ans;
				historyArr = append(historyArr, strToAppend);
				output = ans;
			}else{
				output = "error";
			}
		}else{
			output= "error"
		}
		input.SetText((output))
	})
	equaltoBtn.Importance = widget.HighImportance
	calcContainer:=container.NewVBox(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
			container.NewGridWithColumns(4,
				clearBtn,
				openBtn,
				closeBtn,
				divideBtn,
			),
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				multiplyBtn,
			),container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				minusBtn,
			),container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
			),container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn,
				),
				equaltoBtn,
			),
		),
	))
	w := myApp.NewWindow("Calculator");
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(450,280));
	r, _ := fyne.LoadResourceFromPath("images\\kaint.png")
	w.SetIcon(r)
	w.SetContent(container.NewBorder(nil,nil,nil,nil,calcContainer))
	w.Show()
}