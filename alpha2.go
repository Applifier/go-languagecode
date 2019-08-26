package languagecode

// LanguageAlpha2 is a Language in ISO 639-1 format.
type LanguageAlpha2 struct {
	Language
}

// String implements fmt.Stringer.
func (l LanguageAlpha2) String() string {
	return FormatAlpha2.Serialize(l.Language)
}

// GoString implements fmt.GoStringer.
func (l LanguageAlpha2) GoString() string {
	return l.String()
}

// MarshalText implements encoding.TextMarshaler.
func (l LanguageAlpha2) MarshalText() ([]byte, error) {
	return l.marshalTextWithFormat(FormatAlpha2)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (l *LanguageAlpha2) UnmarshalText(text []byte) error {
	*l = LanguageAlpha2{}
	return l.unmarshalTextWithFormat(FormatAlpha2, text)
}
