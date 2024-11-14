package internal

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var DEFAULT_TIP_PERCENTAGES = []string{"12", "15", "18", "20", "22"}

const TIP_AMOUNT_LABEL_BASE_TXT = "Tip amount:"
const TOTAL_AMOUNT_LABEL_BASE_TXT = "Total amount:"

// Should be a fyne.widget
var _ fyne.Widget = (*TipView)(nil)

type TipView struct {
	widget.BaseWidget
	billAmountEntry  *widget.Entry
	tipPercentSelect *widget.SelectEntry
	finalTipAmount   *widget.Label
	finalTotalAmount *widget.Label
	submitBtn        fyne.Widget
}

func NewTipView() *TipView {
	tv := &TipView{
		billAmountEntry:  widget.NewEntry(),
		tipPercentSelect: widget.NewSelectEntry(DEFAULT_TIP_PERCENTAGES),
		finalTipAmount:   widget.NewLabel(""),
		finalTotalAmount: widget.NewLabel(""),
	}
	tv.tipPercentSelect.SetText("15")
	tv.submitBtn = widget.NewButton("Submit", tv.OnSubmit)

	tv.ExtendBaseWidget(tv)
	return tv
}

func (tv *TipView) OnSubmit() {
	fmt.Println("Submitted")
	fmt.Println(tv)
}

func (tv *TipView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(
		widget.NewLabelWithStyle("Tip Calc", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("Bill Amount"),
		tv.billAmountEntry,
		widget.NewLabel("Tip %"),
		tv.tipPercentSelect,
		widget.NewSeparator(),
		tv.submitBtn,
		widget.NewSeparator(),
		widget.NewLabel(TIP_AMOUNT_LABEL_BASE_TXT),
		widget.NewLabel(TOTAL_AMOUNT_LABEL_BASE_TXT),
	)
	return widget.NewSimpleRenderer(c)
}

func (tv *TipView) GetBillAmount() (float32, error) {
	value, err := strconv.ParseFloat(tv.billAmountEntry.Text, 32)
	return float32(value), err
}
func (tv *TipView) GetTipPercent() (float32, error) {
	value, err := strconv.ParseFloat(tv.tipPercentSelect.Text, 32)
	return float32(value), err
}
func (tv *TipView) GetFinalTipAmount() (float32, error) {
	value, err := strconv.ParseFloat(tv.finalTipAmount.Text, 32)
	return float32(value), err
}
func (tv *TipView) GetFinalTotalAmount() (float32, error) {
	value, err := strconv.ParseFloat(tv.finalTotalAmount.Text, 32)
	return float32(value), err
}

func (tv *TipView) SetBillAmount(amount string) {
	tv.billAmountEntry.SetText(amount)
	tv.billAmountEntry.Refresh()
}
func (tv *TipView) SetTipPercent(percent string) {
	tv.tipPercentSelect.SetText(percent)
	tv.tipPercentSelect.Refresh()
}
func (tv *TipView) SetFinalTipAmount(amount float32) {
	tv.finalTipAmount.SetText(fmt.Sprintf("%.2f", amount))
	tv.finalTipAmount.Refresh()
}
func (tv *TipView) SetFinalTotalAmount(amount float32) {
	tv.finalTotalAmount.SetText(fmt.Sprintf("%.2f", amount))
	tv.finalTotalAmount.Refresh()
}
