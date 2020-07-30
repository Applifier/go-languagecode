# go-languagecode

[![CI status](https://github.com/Unity-Technologies/go-languagecode/workflows/CI/badge.svg)](https://github.com/Unity-Technologies/go-languagecode/actions)
[![GoDoc](https://godoc.org/github.com/Unity-Technologies/go-languagecode?status.svg)](https://godoc.org/github.com/Unity-Technologies/go-languagecode)

Package `languagecode` provides utilities for representing languages in code, and handling their serializations and deserializations in a convenient way.

All serializations will result in `LanguageUndefined` if input data is not a recognized language code. Some conversions are lossy as not all languages have all codes. To avoid loss, you can use the `Has<Format>` method to check whether the language has a code of the wanted format before conversion.

ISO-639-2 and ISO-639-2/B codes are treated distinctly and the B codes will not default to the formal code. If you need defaulting you can perform a `FormatAlpha3B` lookup first and if it returns `LanguageUndefined`, perform a FormatAlpha3 lookup. This design choice is made because the defaulting logic requirements vary case by case.

All the exported types are one word in memory and as such provide fast equality checks, hashing for usage as keys in maps. Conversions between the types are zero overhead (don't even escape the stack), serialization is worst case `O(1)`, and deserialization is worst case `O(n)`.

## License

MIT License. See [LICENSE](LICENSE) for more details.
