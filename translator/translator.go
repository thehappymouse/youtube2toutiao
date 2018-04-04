package translator

type Translator interface {
	Translate(in string) string
}

func Translate(tan Translator, str string) string {
	return tan.Translate(str)
}
