package internal

import "fmt"

type TipController struct {
	*TipModel
	*TipView
}

func NewTipController(tm *TipModel, tv *TipView) *TipController {
	tc := &TipController{
		TipModel: tm,
		TipView:  tv,
	}

	tc.TipView.SetOnSubmitTapped(func(tv *TipView) {
		fmt.Println("In SetOnSubmitTapped")

		// update model to current values
		billAmount, err := tc.TipView.GetBillAmount()
		if err != nil {
			// TODO: use view.showerror
			fmt.Println(billAmount, "should be a number!")
		}
		tipPercent, err := tc.TipView.GetTipPercent()
		if err != nil {
			// TODO: use view.showerror
			fmt.Println(tipPercent, "should be a number!")
		}
		tc.TipModel.SetBillAmount(billAmount)
		tc.TipModel.SetTipPercent(tipPercent)

		// calculate val using model
		finalTip := billAmount * (tipPercent / 100)
		finalTotal := billAmount + finalTip

		// update view
		tc.TipView.SetFinalTipAmount(finalTip)
		tc.TipView.SetFinalTotalAmount(finalTotal)
	})

	return tc
}
