package anki

type Operator interface {
	Do() error
}

type NormalOperator struct {
	Note Note
}

func (n *NormalOperator) Do() error {
	//TODO implement me
	panic("implement me")
}

type VerbOperator struct {
	Note Note
}

func (v *VerbOperator) Do() error {
	//TODO implement me
	panic("implement me")
}

type AdjOperator struct {
	Note Note
}

func (a *AdjOperator) Do() error {
	//TODO implement me
	panic("implement me")
}
