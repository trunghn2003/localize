package localize

type TranslatableField map[string]string

func (tf TranslatableField) Get(locale string, fallback ...string) string {
	if val, ok := tf[locale]; ok {
		return val
	}
	for _, f := range fallback {
		if val, ok := tf[f]; ok {
			return val
		}
	}
	if val, ok := tf["en"]; ok {
		return val
	}
	return ""
}

func (tf *TranslatableField) Set(locale string, value string) {
	if *tf == nil {
		*tf = make(map[string]string)
	}
	(*tf)[locale] = value
}
