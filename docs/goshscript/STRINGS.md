# GOSHScript - Strings

## Escape Codes
Escape codes are more or less the same as in C:

| Escape sequence                                                                      | Meaning                                                                                                                           |
|--------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------|
| `\b`                                                                                 | [Backspace](https://en.wikipedia.org/wiki/Backspace)                                                                              |
| `\f`                                                                                 | [Formfeed](https://en.wikipedia.org/wiki/Formfeed); [Page Break](https://en.wikipedia.org/wiki/Page_Break)                        |
| `\n`                                                                                 | [Newline](https://en.wikipedia.org/wiki/Newline) (Line Feed)                                                                      |
| `\r`                                                                                 | [Carriage Return](https://en.wikipedia.org/wiki/Carriage_Return)                                                                  |
| `\t`                                                                                 | [Horizontal Tab](https://en.wikipedia.org/wiki/Horizontal_Tab)                                                                    |
| `\v`                                                                                 | [Vertical Tab](https://en.wikipedia.org/wiki/Vertical_Tab)                                                                        |
| `\\`                                                                                 | [Backslash](https://en.wikipedia.org/wiki/Backslash)                                                                              |
| `\'`                                                                                 | [Apostrophe](https://en.wikipedia.org/wiki/Apostrophe) or single [quotation mark](https://en.wikipedia.org/wiki/Quotation_mark)   |
| `\ `                                                                                 | Space character. Needs escaping when used in [identifiers (or unquoted strings)](goshscript/IDENTIFIERS.md)
| `\"`                                                                                 | double [quotation mark](https://en.wikipedia.org/wiki/Quotation_mark)                                                             |
| `\xhh...` <br />`\uhh...` <sup>[[Note 1]](#escape-codes-notes)</sup> <br />`\Uhh...` | [character](https://en.wikipedia.org/wiki/Character_(computing)) in [hexadecimal](https://en.wikipedia.org/wiki/Hexadecimal), each *h* represents a digit, can be an ASCII value or [Unicode](https://en.wikipedia.org/wiki/Unicode) [code point](https://en.wikipedia.org/wiki/Code_point) value <br /> Max Value is [Go's unicode.MaxRune](https://golang.org/pkg/unicode/#pkg-constants), currently 0x10ffff |

<div id="escape-codes-notes" style="padding-left: 30px;">
[Note 1] \x, \u, \U are all synonyms in GoshScript. All three are available for easier switch from other languages
</div>