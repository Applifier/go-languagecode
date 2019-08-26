package languagecode

// LanguageAlpha3B is a Language in ISO 639-2/B format.
type LanguageAlpha3B struct {
	Language
}

// String implements fmt.Stringer.
func (l LanguageAlpha3B) String() string {
	return FormatAlpha3B.Serialize(l.Language)
}

// GoString implements fmt.GoStringer.
func (l LanguageAlpha3B) GoString() string {
	return l.String()
}

// MarshalText implements encoding.TextMarshaler.
func (l LanguageAlpha3B) MarshalText() ([]byte, error) {
	return l.marshalTextWithFormat(FormatAlpha3B)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (l *LanguageAlpha3B) UnmarshalText(text []byte) error {
	*l = LanguageAlpha3B{}
	return l.unmarshalTextWithFormat(FormatAlpha3B, text)
}
