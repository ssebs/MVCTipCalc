package internal

import "time"

type TipController struct {
	*TipModel
	*TipView
}

func NewTipController(tm *TipModel, tv *TipView) *TipController {
	tc := &TipController{
		TipModel: tm,
		TipView:  tv,
	}

	tc.TipView.SetOnSubmit(func() {
		tc.UpdateModelFromView()
		tc.CalcTipAndUpdate()

		tc.TipView.SetErrorMsg("This button is redundant")
		// TODO: https://gobyexample.com/timers use TimedFunc
	})

	tc.TipView.SetOnSelectTip(func(s string) {
		tc.UpdateModelFromView()
		tc.CalcTipAndUpdate()
	})
	tc.TipView.SetBillAmountEntryOnChanged(func(s string) {
		tc.UpdateModelFromView()
		tc.CalcTipAndUpdate()
	})

	return tc
}

func (tc *TipController) UpdateModelFromView() {
	// update model to current values
	billAmount, err := tc.TipView.GetBillAmount()
	if err != nil {
		tc.TipView.SetErrorMsg("Bill amount must be a number.")
		return
	}
	tipPercent, err := tc.TipView.GetTipPercent()
	if err != nil {
		tc.TipView.SetErrorMsg("Tip % must be a number.")
		return
	}
	tc.TipModel.SetBillAmount(billAmount)
	tc.TipModel.SetTipPercent(tipPercent)
}

func (tc *TipController) CalcTipAndUpdate() {
	// calculate tip and total
	finalTip := tc.TipModel.GetBillAmount() * (tc.TipModel.GetTipPercent() / 100)
	finalTotal := tc.TipModel.GetBillAmount() + finalTip

	// update view
	tc.TipView.SetFinalTipAmount(finalTip)
	tc.TipView.SetFinalTotalAmount(finalTotal)
}

func (tc *TipController) TimedFunc(d time.Duration, f func(*TipController)) {
	// TODO: implement
}
