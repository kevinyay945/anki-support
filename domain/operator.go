package domain

type Operator interface {
	Do() error
}

type NormalOperator struct {
	Note           Note
	gpter          GPTer
	textToSpeecher TextToSpeecher
	ankier         Ankier
}

func (n *NormalOperator) Do() error {
	return nil
}

type VerbOperator struct {
	Note           Note
	gpter          GPTer
	textToSpeecher TextToSpeecher
	ankier         Ankier
}

func (v *VerbOperator) Do() error {
	//TODO implement me
	panic("implement me")
}

type AdjOperator struct {
	Note           Note
	gpter          GPTer
	textToSpeecher TextToSpeecher
	ankier         Ankier
}

func (a *AdjOperator) Do() error {
	//TODO implement me
	panic("implement me")
}
