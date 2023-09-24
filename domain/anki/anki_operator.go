package anki

import "fmt"

type OperatorGenerate struct {
}

func (g *OperatorGenerate) GetByNote(note Note) (o Operator, err error) {
	switch note.ModelName {
	case "Japanese (recognition&recall) 動詞篇":
		o = &VerbOperator{
			Note: note,
		}
	case "Japanese (recognition&recall) 形容詞":
		o = &AdjOperator{
			Note: note,
		}
	case "Japanese (recognition&recall)":
		o = &NormalOperator{
			Note: note,
		}
	default:
		err = fmt.Errorf("don't support for this modelType: %s", note.ModelName)
	}
	return
}

type OperatorGenerator interface {
	GetByNote(note Note) (o Operator, err error)
}
