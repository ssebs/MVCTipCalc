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
	onTapped         func(*TipView)
}

func NewTipView() *TipView {
	tv := &TipView{
		billAmountEntry:  widget.NewEntry(),
		tipPercentSelect: widget.NewSelectEntry(DEFAULT_TIP_PERCENTAGES),
		finalTipAmount:   widget.NewLabel(TIP_AMOUNT_LABEL_BASE_TXT),
		finalTotalAmount: widget.NewLabel(TOTAL_AMOUNT_LABEL_BASE_TXT),
	}
	tv.tipPercentSelect.SetText("15")

	tv.submitBtn = widget.NewButton("Submit", func() {
		if tv.onTapped != nil {
			tv.onTapped(tv)
		} else {
			fmt.Println("Submitted, use SetOnTapped to add a callback.")
			fmt.Println(tv)
		}
	})

	tv.ExtendBaseWidget(tv)
	return tv
}

func (tv *TipView) SetOnSubmitTapped(f func(*TipView)) {
	tv.onTapped = f
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
		tv.finalTipAmount,
		tv.finalTotalAmount,
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
	tv.finalTipAmount.SetText(fmt.Sprintf("%s: %.2f", TIP_AMOUNT_LABEL_BASE_TXT, amount))
	tv.finalTipAmount.Refresh()
}
func (tv *TipView) SetFinalTotalAmount(amount float32) {
	tv.finalTotalAmount.SetText(fmt.Sprintf("%s: %.2f", TOTAL_AMOUNT_LABEL_BASE_TXT, amount))
	tv.finalTotalAmount.Refresh()
}

func (tv *TipView) String() string {
	billAmount, _ := tv.GetBillAmount()
	tipPercent, _ := tv.GetTipPercent()
	tipAmount, _ := tv.GetFinalTipAmount()
	totalAmount, _ := tv.GetFinalTotalAmount()

	return fmt.Sprintf(
		"bill: %.2f, tip: %.2f, tipAmount: %.2f, totalAmount: %.2f",
		billAmount,
		tipPercent,
		tipAmount,
		totalAmount,
	)
}
