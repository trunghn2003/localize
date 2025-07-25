package localize

var DefaultLocale = "en"

func (tf TranslatableField) Auto() string {
	return tf.Get(DefaultLocale)
}
