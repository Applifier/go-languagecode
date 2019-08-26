package languagecode

// Format represents a specific language code format with a specific
// serialization.
type Format int

const (
	// FormatAlpha3 is an ISO-639-2 language code.
	FormatAlpha3 Format = iota
	// FormatAlpha3B is an ISO-639-2/B language code.
	FormatAlpha3B
	// FormatAlpha2 is an ISO-639-1 language code.
	FormatAlpha2
	formatsCount
)

// Serialize the specified Language into a language code string of the Format.
func (f Format) Serialize(language Language) string {
	return codes[language.code][f]
}

// Deserialize the specified language code string of the Format into a
// Language.
func (f Format) Deserialize(languageCode string) Language {
	return languages[f][languageCode]
}

var languages = func() (l [formatsCount]map[string]Language) {
	for f := Format(0); f < formatsCount; f++ {
		l[f] = make(map[string]Language, len(codes))
		for j, languageCodes := range codes {
			l[f][languageCodes[f]] = Language{code: code(j)}
		}
	}
	return
}()
