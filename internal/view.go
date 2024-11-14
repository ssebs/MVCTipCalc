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

// Implement fyne.widget
var _ fyne.Widget = (*TipView)(nil)

type TipView struct {
	widget.BaseWidget
	billAmountEntry  *widget.Entry
	tipPercentSelect *widget.SelectEntry
	finalTipAmount   *widget.Label
	finalTotalAmount *widget.Label
	errorMsg         *widget.Label
	submitBtn        fyne.Widget
	onSubmit         func()
}

func NewTipView() *TipView {
	boldStyle := fyne.TextStyle{Bold: true}

	tv := &TipView{
		billAmountEntry:  widget.NewEntry(),
		tipPercentSelect: widget.NewSelectEntry(DEFAULT_TIP_PERCENTAGES),
		errorMsg:         widget.NewLabel(""),
		finalTipAmount:   widget.NewLabelWithStyle("", fyne.TextAlignLeading, boldStyle),
		finalTotalAmount: widget.NewLabelWithStyle("", fyne.TextAlignLeading, boldStyle),
	}
	tv.billAmountEntry.PlaceHolder = "$"
	tv.tipPercentSelect.SetText(DEFAULT_TIP_PERCENTAGES[1])
	tv.errorMsg.Importance = widget.DangerImportance

	// On btn press, check if onSubmit is set to anything and do that. (Controller should use SetOnSubmit)
	tv.submitBtn = widget.NewButton("Submit", func() {
		if tv.onSubmit != nil {
			tv.onSubmit()
		} else {
			fmt.Println("Submitted, use SetOnTapped to add a callback.")
			fmt.Println(tv)
		}
	})

	tv.ExtendBaseWidget(tv)
	return tv
}

// Run f(s) when Select changes values
func (tv *TipView) SetOnSelectTip(f func(s string)) {
	tv.tipPercentSelect.OnChanged = f
}

// Run f() when Submit button is Tapped
func (tv *TipView) SetOnSubmit(f func()) {
	tv.onSubmit = f
}

func (tv *TipView) SetErrorMsg(err string) {
	tv.errorMsg.SetText(err)
	tv.errorMsg.Refresh()
}

func (tv *TipView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(
		widget.NewLabelWithStyle("Tip Calculator", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel("Bill Amount"),
		tv.billAmountEntry,
		widget.NewLabel("Tip %"),
		tv.tipPercentSelect,
		widget.NewSeparator(),
		tv.submitBtn,
		widget.NewSeparator(),
		container.NewHBox(
			widget.NewLabel(TIP_AMOUNT_LABEL_BASE_TXT),
			tv.finalTipAmount,
		),
		container.NewHBox(
			widget.NewLabel(TOTAL_AMOUNT_LABEL_BASE_TXT),
			tv.finalTotalAmount,
		),
		tv.errorMsg,
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
	tv.finalTipAmount.SetText(fmt.Sprintf("$%.2f", amount))
	tv.finalTipAmount.Refresh()
}
func (tv *TipView) SetFinalTotalAmount(amount float32) {
	tv.finalTotalAmount.SetText(fmt.Sprintf("$%.2f", amount))
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
