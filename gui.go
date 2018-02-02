package main

import (
	"github.com/andlabs/ui"

	"strconv"
)

type areaHandler struct {

}

func (areaHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {

}

func (areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {

}

func (areaHandler) MouseCrossed(a *ui.Area, left bool) {

}

func (areaHandler) DragBroken(a *ui.Area) {

}

func (areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	return true;
}

var window *ui.Window
var tab *ui.Tab

func HorizontalBoxWrap(child ui.Control) *ui.Box{
	box := ui.NewHorizontalBox()
	box.Append(child, false)
	return box
}
func VerticalBoxWrap(child ui.Control) *ui.Box{
	box := ui.NewVerticalBox()
	box.Append(child, false)
	return box
}

func main() {
	err := ui.Main(func() {
		tab = ui.NewTab()
		tab1()
		tab2()
		tab3()

		window = ui.NewWindow("Hello", 512, 256, true)
		window.SetMargined(true)
		window.SetChild(tab)


		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func tab1() {
	var h areaHandler = areaHandler{}

	//ui.DrawContext{}
	input := ui.NewArea(h)
	//input := ui.NewScrollingArea(h, 500, 200)

	button := ui.NewButton("Greet")
	button2 := ui.NewButton("boy")
	boxButton := ui.NewHorizontalBox()
	boxButton.Append(button, false)
	boxButton.Append(button2, false)

	boxLabel := ui.NewVerticalBox()
	greeting := ui.NewLabel("")
	greeting2 := ui.NewLabel("")
	boxLabel.Append(greeting, false)
	boxLabel.Append(greeting2, false)

	button.OnClicked(func(*ui.Button) {
		greeting.SetText("Hello, " + "!") //input.Text()
	})
	button2.OnClicked(func(*ui.Button) {
		greeting2.SetText("Hello, " + "!") //input.Text()
	})

	check := ui.NewCheckbox("check")
	check.OnToggled(func (*ui.Checkbox) {
		if check.Checked() {
			ui.MsgBox(window, "Check On", "check on")
			greeting.Show()
		} else {
			ui.MsgBox(window, "Check Off", "check off")
			greeting.Hide()
		}
	})
	boxCheck := ui.NewHorizontalBox()
	boxCheck.Append(check, false)


	boxCombo := ui.NewVerticalBox()
	comboBox := ui.NewCombobox()
	//c1 := ui.NewLabel("c1")
	//c2 := ui.NewLabel("c2")
	comboBox.Append("c1")
	comboBox.Append("c2")
	boxCombo.Append(comboBox, false)



	box := ui.NewVerticalBox()
	box.Append(ui.NewLabel("Enter your name:"), false)
	box.Append(input, false)
	box.Append(boxCheck, false)
	box.Append(boxButton, false)
	box.Append(boxLabel, false)
	box.Append(boxCombo, false)

	tab.Append("tab1", box)
}

func tab2() {
	picker := ui.NewDateTimePicker()
	boxTab2 := ui.NewVerticalBox()
	boxTab2.Append(picker, false)

	spinBox := ui.NewSpinbox(1, 10)
	spinLabel := ui.NewLabel("1")
	bar := ui.NewProgressBar()
	spinBox.OnChanged(func (x *ui.Spinbox) {
		spinLabel.SetText(strconv.Itoa(x.Value()))
		bar.SetValue(x.Value() * 10)
	})
	boxTab2.Append(HorizontalBoxWrap(spinBox), false)
	boxTab2.Append(HorizontalBoxWrap(spinLabel), false)
	boxTab2.Append(HorizontalBoxWrap(bar), false)
	boxTab2.Append(ui.NewHorizontalSeparator(), false)

	slider := ui.NewSlider(1, 10)
	sliderLabel := ui.NewLabel("1")
	slider.OnChanged(func (x *ui.Slider) {
		sliderLabel.SetText(strconv.Itoa(x.Value()))
	})
	boxTab2.Append(VerticalBoxWrap(slider), false)
	boxTab2.Append(VerticalBoxWrap(sliderLabel), false)


	tab.Append("tab2", boxTab2)
}

func tab3() {
	box := ui.NewVerticalBox()

	radioButton := ui.NewRadioButtons()//bug here
	entry := ui.NewEntry()
	label := ui.NewLabel("1")
	entry.OnChanged(func(x *ui.Entry) {
		label.SetText(x.Text())
	})

	box.Append(VerticalBoxWrap(radioButton), false)
	box.Append(VerticalBoxWrap(label), false)
	tab.Append("tab3", box)


}