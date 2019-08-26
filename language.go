// Package languagecode provides utilities for representing languages in code,
// and handling their serializations and deserializations in a convenient way.
//
// All serializations will result in `LanguageUndefined` if input data is not a
// recognized language code. Some conversions are lossy as not all languages
// have all codes. To avoid loss, you can use the Has<Format> methods to check
// whether the language has a code of the wanted format before conversion.
//
// ISO-639-2 and ISO-639-2/B codes are treated distinctly and the B codes will
// not default to the formal code. If you need defaulting you can perform a
// FormatAlpha3B lookup first and if it returns `LanguageUndefined`, perform a
// FormatAlpha3 lookup. This design choice is made because the defaulting logic
// requirements vary case by case.
//
// All the exported types are one word in memory and as such provide fast
// equality checks, hashing for usage as keys in maps. Conversions between the
// types are zero overhead (don't even escape the stack), serialization is
// worst case O(1), and deserialization is worst case O(n).
package languagecode

// Language represents the language codes of a language.
type Language struct {
	code code
}

// LanguageUndefined represents an undefined language.
var LanguageUndefined = Language{}

// Alpha3 returns the language with the Alpha3 serialization.
func (l Language) Alpha3() LanguageAlpha3 {
	return LanguageAlpha3{Language: l}
}

// HasAlpha3 returns a boolean indicating whether or not the language has a
// designated Alpha3 code.
func (l Language) HasAlpha3() bool {
	return FormatAlpha3.Serialize(l) != empty3
}

// Alpha3B returns the language with the Alpha3B serialization.
func (l Language) Alpha3B() LanguageAlpha3B {
	return LanguageAlpha3B{Language: l}
}

// HasAlpha3B returns a boolean indicating whether or not the language has a
// designated Alpha3B code.
func (l Language) HasAlpha3B() bool {
	return FormatAlpha3B.Serialize(l) != empty3
}

// Alpha2 returns the language with the Alpha2 serialization.
func (l Language) Alpha2() LanguageAlpha2 {
	return LanguageAlpha2{Language: l}
}

// HasAlpha2 returns a boolean indicating whether or not the language has a
// designated Alpha2 code.
func (l Language) HasAlpha2() bool {
	return FormatAlpha2.Serialize(l) != empty2
}

// GoString implements fmt.GoStringer.
func (l Language) GoString() string {
	return "languagecode.Language{" + FormatAlpha3.Serialize(l) + "}"
}

func (l Language) marshalTextWithFormat(format Format) ([]byte, error) {
	return []byte(format.Serialize(l)), nil
}

func (l *Language) unmarshalTextWithFormat(format Format, text []byte) error {
	*l = format.Deserialize(string(text))
	return nil
}
