package languagecode

// LanguageAlpha3 is a Language in ISO 639-2 format.
type LanguageAlpha3 struct {
	Language
}

// String implements fmt.Stringer.
func (l LanguageAlpha3) String() string {
	return FormatAlpha3.Serialize(l.Language)
}

// GoString implements fmt.GoStringer.
func (l LanguageAlpha3) GoString() string {
	return l.String()
}

// MarshalText implements encoding.TextMarshaler.
func (l LanguageAlpha3) MarshalText() ([]byte, error) {
	return l.marshalTextWithFormat(FormatAlpha3)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (l *LanguageAlpha3) UnmarshalText(text []byte) error {
	*l = LanguageAlpha3{}
	return l.unmarshalTextWithFormat(FormatAlpha3, text)
}
