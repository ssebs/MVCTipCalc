package internal

type TipModel struct {
	billAmount float32
	tipPercent float32
}

func NewTipModel() *TipModel {
	return &TipModel{
		billAmount: 0,
		tipPercent: 15,
	}
}

func (tm *TipModel) GetBillAmount() float32 {
	return tm.billAmount
}

func (tm *TipModel) GetTipPercent() float32 {
	return tm.tipPercent
}

func (tm *TipModel) SetBillAmount(amount float32) {
	tm.billAmount = amount
}

func (tm *TipModel) SetTipPercent(amount float32) {
	tm.tipPercent = amount
}
