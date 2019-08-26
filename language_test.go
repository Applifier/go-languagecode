package languagecode_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/Applifier/go-languagecode"
)

func TestFormat_Deserialize(t *testing.T) {
	tests := []struct {
		test         string
		languageCode string
		format       languagecode.Format
		language     languagecode.Language
	}{
		{
			"FormatAlpha3B",
			"cze",
			languagecode.FormatAlpha3B,
			languagecode.CES,
		},
		{
			"FormatAlpha2",
			"el",
			languagecode.FormatAlpha2,
			languagecode.ELL,
		},
		{
			"FormatAlpha3",
			"ell",
			languagecode.FormatAlpha3,
			languagecode.ELL,
		},
		{
			"invalid format",
			"ell",
			languagecode.FormatAlpha2,
			languagecode.LanguageUndefined,
		},
		{
			"invalid language",
			"aaa",
			languagecode.FormatAlpha3,
			languagecode.LanguageUndefined,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			requireEqual(t, tt.language, tt.format.Deserialize(tt.languageCode))
		})
	}
}

func TestFormat_Serialize(t *testing.T) {
	tests := []struct {
		test         string
		language     languagecode.Language
		format       languagecode.Format
		languageCode string
	}{
		{
			"FormatAlpha3B",
			languagecode.CES,
			languagecode.FormatAlpha3B,
			"cze",
		},
		{
			"FormatAlpha2",
			languagecode.ELL,
			languagecode.FormatAlpha2,
			"el",
		},
		{
			"FormatAlpha3",
			languagecode.ELL,
			languagecode.FormatAlpha3,
			"ell",
		},
		{
			"invalid language",
			languagecode.LanguageUndefined,
			languagecode.FormatAlpha3,
			"---",
		},
		{
			"missing code",
			languagecode.FIN,
			languagecode.FormatAlpha3B,
			"---",
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			requireEqual(t, tt.languageCode, tt.format.Serialize(tt.language))
		})
	}
}

func TestConversions(t *testing.T) {
	requireTrue(t, languagecode.ENG == languagecode.ENG) // nolint: staticcheck
	requireFalse(t, languagecode.ENG == languagecode.FIN)
	requireTrue(t, languagecode.ENG.Alpha3() == languagecode.ENG.Alpha3()) // nolint: staticcheck
	requireFalse(t, languagecode.ENG.Alpha3() == languagecode.FIN.Alpha3())
	requireTrue(t, languagecode.ENG.Alpha2() == languagecode.ENG.Alpha2()) // nolint: staticcheck
	requireFalse(t, languagecode.ENG.Alpha2() == languagecode.FIN.Alpha2())
	requireTrue(t, languagecode.ENG.Alpha3B() == languagecode.ENG.Alpha3B()) // nolint: staticcheck
	requireFalse(t, languagecode.ENG.Alpha3B() == languagecode.FIN.Alpha3B())
	requireTrue(t, languagecode.ENG.Alpha3().Language == languagecode.ENG.Alpha3().Language) // nolint: staticcheck
	requireFalse(t, languagecode.ENG.Alpha3().Language == languagecode.FIN.Alpha3().Language)
	requireTrue(t, languagecode.ENG.Alpha2().Language == languagecode.ENG.Alpha2().Language) // nolint: staticcheck
	requireFalse(t, languagecode.ENG.Alpha2().Language == languagecode.FIN.Alpha2().Language)
	requireTrue(t, languagecode.ENG.Alpha3B().Language == languagecode.ENG.Alpha3B().Language) // nolint: staticcheck
	requireFalse(t, languagecode.ENG.Alpha3B().Language == languagecode.FIN.Alpha3B().Language)
	requireTrue(t, languagecode.ENG.Alpha3().Alpha2().Alpha3B().Language == languagecode.ENG)
	requireTrue(t, languagecode.ENG.Alpha2().Alpha3().Alpha3B().Language == languagecode.ENG)
	requireTrue(t, languagecode.ENG.Alpha3B().Alpha3().Language == languagecode.ENG)
	requireTrue(t, languagecode.ENG.Alpha3B().Alpha2().Language == languagecode.ENG)
	requireTrue(t, languagecode.ENG.HasAlpha3())
	requireFalse(t, languagecode.ENG.HasAlpha3B())
	requireTrue(t, languagecode.ENG.HasAlpha2())
	requireTrue(t, languagecode.ELL.HasAlpha3())
	requireTrue(t, languagecode.ELL.HasAlpha3B())
	requireTrue(t, languagecode.ELL.HasAlpha2())
}

func TestDebug(t *testing.T) {
	requireEqual(t, fmt.Sprintf("%#v", languagecode.ELL), "languagecode.Language{ell}")
	requireEqual(t, fmt.Sprintf("%#v", languagecode.ELL.Alpha3()), "ell")
	requireEqual(t, fmt.Sprintf("%#v", languagecode.ELL.Alpha2()), "el")
	requireEqual(t, fmt.Sprintf("%#v", languagecode.ELL.Alpha3B()), "gre")
}

func TestUnmarshaling(t *testing.T) {
	t.Run("Alpha3 map", func(t *testing.T) {
		// Arrange
		type x struct {
			Y map[languagecode.LanguageAlpha3]int `json:"y"`
		}
		expected := x{
			Y: map[languagecode.LanguageAlpha3]int{
				languagecode.ELL.Alpha3(): 123,
				languagecode.CES.Alpha3(): 234,
			},
		}
		input := `{"y": {"ell": 123, "ces": 234}}`
		var received x

		// Act
		requireNoError(t, json.Unmarshal([]byte(input), &received))

		// Assert
		requireEqual(t, expected, received)
	})
	t.Run("Alpha3B map", func(t *testing.T) {
		// Arrange
		type x struct {
			Y map[languagecode.LanguageAlpha3B]int `json:"y"`
		}
		expected := x{
			Y: map[languagecode.LanguageAlpha3B]int{
				languagecode.ELL.Alpha3B(): 123,
				languagecode.CES.Alpha3B(): 234,
			},
		}
		input := `{"y": {"gre": 123, "cze": 234}}`
		var received x

		// Act
		requireNoError(t, json.Unmarshal([]byte(input), &received))

		// Assert
		requireEqual(t, expected, received)
	})
	t.Run("Alpha2 map", func(t *testing.T) {
		// Arrange
		type x struct {
			Y map[languagecode.LanguageAlpha2]int `json:"y"`
		}
		expected := x{
			Y: map[languagecode.LanguageAlpha2]int{
				languagecode.ELL.Alpha2(): 123,
				languagecode.CES.Alpha2(): 234,
			},
		}
		input := `{"y": {"el": 123, "cs": 234}}`
		var received x

		// Act
		requireNoError(t, json.Unmarshal([]byte(input), &received))

		// Assert
		requireEqual(t, expected, received)
	})
}

func TestMarshaling(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "Alpha3 list",
			input:    []languagecode.LanguageAlpha3{languagecode.ELL.Alpha3(), languagecode.CES.Alpha3()},
			expected: `["ell","ces"]`,
		},
		{
			name:     "Alpha3B list",
			input:    []languagecode.LanguageAlpha3B{languagecode.ELL.Alpha3B(), languagecode.CES.Alpha3B()},
			expected: `["gre","cze"]`,
		},
		{
			name:     "Alpha2 list",
			input:    []languagecode.LanguageAlpha2{languagecode.ELL.Alpha2(), languagecode.CES.Alpha2()},
			expected: `["el","cs"]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.input)
			requireNoError(t, err)
			requireEqual(t, tc.expected, string(data))
		})
	}
}

func ExampleFormat_twoWayConversion() {
	const countryCodeFormat = languagecode.FormatAlpha3

	type externalData struct {
		Language string
	}

	type internalData struct {
		Language languagecode.Language
	}

	toInternalData := func(d externalData) internalData {
		return internalData{Language: countryCodeFormat.Deserialize(d.Language)}
	}

	toExternalData := func(d internalData) externalData {
		return externalData{Language: countryCodeFormat.Serialize(d.Language)}
	}

	edata := externalData{Language: "ell"}
	idata := toInternalData(edata)
	edata = toExternalData(idata)
	// Output: {ell}
	fmt.Println(edata)
}

func ExampleFormat_Deserialize() {
	country := languagecode.FormatAlpha2.Deserialize("en")
	// Output: true
	fmt.Println(country == languagecode.ENG)
}

func ExampleFormat_Serialize() {
	countryCodeAlpha2 := languagecode.FormatAlpha2.Serialize(languagecode.ENG)
	// Output: en
	fmt.Println(countryCodeAlpha2)
}

func Example_unmarshalJSON() {
	data := []byte(`
		{
			"ell": 1234,
			"eng": 2345
		}
	`)
	var mapping map[languagecode.LanguageAlpha3]int
	_ = json.Unmarshal(data, &mapping)
	// Output: map[languagecode.LanguageAlpha3]int{ell:1234, eng:2345}
	fmt.Printf("%#v\n", mapping)
}

func Example_marshalJSON() {
	mapping := map[languagecode.LanguageAlpha3]int{
		languagecode.ELL.Alpha3(): 1234,
		languagecode.ENG.Alpha3(): 2345,
	}
	data, _ := json.Marshal(mapping)
	// Output: {"ell":1234,"eng":2345}
	fmt.Println(string(data))
}

func requireEqual(tb testing.TB, expected, received interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(expected, received) {
		tb.Fatalf("expected %#v, received %#v", expected, received)
	}
}

func requireNoError(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("expected no error, got %#v", err)
	}
}

func requireTrue(tb testing.TB, v bool) {
	tb.Helper()
	if !v {
		tb.Fatal("expected the value to be true")
	}
}

func requireFalse(tb testing.TB, v bool) {
	tb.Helper()
	if v {
		tb.Fatal("expected the value to be false")
	}
}
