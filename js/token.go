package js // import "github.com/tdewolff/parse/js"

import (
	"io"
	"strconv"
	"unicode"

	"github.com/tdewolff/parse"
)

var identifierStart = []*unicode.RangeTable{unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl}
var identifierPart = []*unicode.RangeTable{unicode.Lu, unicode.Ll, unicode.Lt, unicode.Lm, unicode.Lo, unicode.Nl, unicode.Mn, unicode.Mc, unicode.Nd, unicode.Pc}

////////////////////////////////////////////////////////////////

// TokenType determines the type of token, eg. a number or a semicolon.
type TokenType uint32

// TokenType values.
const (
	ErrorToken          TokenType = iota // extra token when errors occur
	UnknownToken                         // extra token when no token can be matched
	WhitespaceToken                      // space \t \v \f
	LineTerminatorToken                  // \r \n \r\n
	CommentToken
	IdentifierToken
	PunctuatorToken /* { } ( ) [ ] . ; , < > <= >= == != === !==  + - * % ++ -- << >>
	   >>> & | ^ ! ~ && || ? : = += -= *= %= <<= >>= >>>= &= |= ^= / /= */
	NumericToken
	StringToken
	RegexpToken
)

// String returns the string representation of a TokenType.
func (tt TokenType) String() string {
	switch tt {
	case ErrorToken:
		return "Error"
	case UnknownToken:
		return "Unknown"
	case WhitespaceToken:
		return "Whitespace"
	case LineTerminatorToken:
		return "LineTerminator"
	case CommentToken:
		return "Comment"
	case IdentifierToken:
		return "Identifier"
	case PunctuatorToken:
		return "Punctuator"
	case NumericToken:
		return "Numeric"
	case StringToken:
		return "String"
	case RegexpToken:
		return "Regexp"
	}
	return "Invalid(" + strconv.Itoa(int(tt)) + ")"
}

////////////////////////////////////////////////////////////////

// Tokenizer is the state for the tokenizer.
type Tokenizer struct {
	r    *parse.ShiftBuffer
	line int

	regexpState bool
}

// NewTokenizer returns a new Tokenizer for a given io.Reader.
func NewTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{
		parse.NewShiftBuffer(r),
		1,
		false,
	}
}

// Line returns the current line that is being tokenized (1 + number of \n, \r or \r\n encountered).
func (z Tokenizer) Line() int {
	return z.line
}

// Err returns the error encountered during tokenization, this is often io.EOF but also other errors can be returned.
func (z Tokenizer) Err() error {
	return z.r.Err()
}

// IsEOF returns true when it has encountered EOF and thus loaded the last buffer in memory.
func (z Tokenizer) IsEOF() bool {
	return z.r.IsEOF()
}

// Next returns the next Token. It returns ErrorToken when an error was encountered. Using Err() one can retrieve the error message.
func (z *Tokenizer) Next() (TokenType, []byte) {
	var tt TokenType
	c := z.r.Peek(0)
	switch c {
	case '(', ')', '[', ']', '{', '}', ';', ',', '~', '?', ':':
		z.r.Move(1)
		tt = PunctuatorToken
	case '<', '>', '=', '!', '+', '-', '*', '%', '&', '|', '^':
		if z.consumeLongPunctuatorToken() {
			tt = PunctuatorToken
		}
	case '/':
		if z.consumeCommentToken() {
			tt = CommentToken
		} else if z.regexpState && z.consumeRegexpToken() {
			tt = RegexpToken
		} else if z.consumeLongPunctuatorToken() {
			tt = PunctuatorToken
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
		if z.consumeNumericToken() {
			tt = NumericToken
		} else if c == '.' {
			z.r.Move(1)
			tt = PunctuatorToken
		}
	case '\'', '"':
		if z.consumeStringToken() {
			tt = StringToken
		}
	case ' ', '\t', '\v', '\f':
		z.r.Move(1)
		for z.consumeWhitespace() {
		}
		tt = WhitespaceToken
	case '\n', '\r':
		z.r.Move(1)
		for z.consumeLineTerminator() {
		}
		tt = LineTerminatorToken
	default:
		if c >= 0xC0 && z.consumeWhitespace() {
			for z.consumeWhitespace() {
			}
			tt = WhitespaceToken
		} else if c >= 0xC0 && z.consumeLineTerminator() {
			for z.consumeLineTerminator() {
			}
			tt = LineTerminatorToken
		} else if z.consumeIdentifierToken() {
			tt = IdentifierToken
		} else if c >= 0xC0 {
			if z.consumeWhitespace() {
				for z.consumeWhitespace() {
				}
				tt = WhitespaceToken
			} else if z.consumeLineTerminator() {
				for z.consumeLineTerminator() {
				}
				tt = LineTerminatorToken
			} else {
				z.consumeRune()
				tt = UnknownToken
			}
		} else if c == 0 && z.Err() != nil {
			return ErrorToken, []byte{}
		} else {
			z.r.Move(1)
			tt = UnknownToken
		}
	}

	// differentiate between divisor and regexp state, because the '/' character is ambiguous!
	if tt != WhitespaceToken && tt != CommentToken {
		z.regexpState = false
		if tt == PunctuatorToken && z.r.Pos() == 1 {
			if c := z.r.Buffered()[0]; c == '(' || c == ',' || c == '=' || c == ':' || c == '[' || c == '!' || c == '&' || c == '|' || c == '?' || c == '{' || c == '}' || c == ';' {
				z.regexpState = true
			}
		}
	}
	return tt, z.r.Shift()
}

////////////////////////////////////////////////////////////////

/*
The following functions follow the specifications at http://www.ecma-international.org/ecma-262/5.1/
*/

func (z *Tokenizer) consumeRune() bool {
	c := z.r.Peek(0)
	if c < 0xC0 {
		z.r.Move(1)
	} else if c < 0xE0 {
		z.r.Move(2)
	} else if c < 0xF0 {
		z.r.Move(3)
	} else {
		z.r.Move(4)
	}
	return true
}

func (z *Tokenizer) consumeWhitespace() bool {
	c := z.r.Peek(0)
	if c == ' ' || c == '\t' || c == '\v' || c == '\f' {
		z.r.Move(1)
		return true
	} else if c >= 0xC0 {
		if r := z.r.PeekRune(0); r == '\u00A0' || r == '\uFEFF' || unicode.Is(unicode.Zs, r) {
			return z.consumeRune()
		}
	}
	return false
}

func (z *Tokenizer) consumeLineTerminator() bool {
	c := z.r.Peek(0)
	if c == '\n' {
		z.line++
		z.r.Move(1)
		return true
	} else if c == '\r' {
		z.line++
		if z.r.Peek(1) == '\n' {
			z.r.Move(2)
		} else {
			z.r.Move(1)
		}
		return true
	} else if c >= 0xC0 {
		if r := z.r.PeekRune(0); r == '\u2028' || r == '\u2029' {
			z.line++
			return z.consumeRune()
		}
	}
	return false
}

func (z *Tokenizer) consumeDigit() bool {
	c := z.r.Peek(0)
	if c >= '0' && c <= '9' {
		z.r.Move(1)
		return true
	}
	return false
}

func (z *Tokenizer) consumeHexDigit() bool {
	c := z.r.Peek(0)
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
		z.r.Move(1)
		return true
	}
	return false
}

func (z *Tokenizer) consumeEscape() bool {
	if z.r.Peek(0) != '\\' {
		return false
	} else if z.consumeHexEscape() || z.consumeUnicodeEscape() {
		return true
	}
	z.r.Move(1)
	if z.r.Peek(0) == '0' {
		nOld := z.r.Pos()
		z.r.Move(1)
		if !z.consumeDigit() {
			return true
		}
		z.r.MoveTo(nOld)
		return false
	}
	z.r.Move(1)
	return true
}
func (z *Tokenizer) consumeHexEscape() bool {
	if z.r.Peek(0) != '\\' || z.r.Peek(1) != 'x' {
		return false
	}
	nOld := z.r.Pos()
	z.r.Move(2)
	for k := 0; k < 2; k++ {
		if !z.consumeHexDigit() {
			z.r.MoveTo(nOld)
			return false
		}
	}
	return true
}

func (z *Tokenizer) consumeUnicodeEscape() bool {
	if z.r.Peek(0) != '\\' || z.r.Peek(1) != 'u' {
		return false
	}
	nOld := z.r.Pos()
	z.r.Move(2)
	for k := 0; k < 4; k++ {
		if !z.consumeHexDigit() {
			z.r.MoveTo(nOld)
			return false
		}
	}
	return true
}

////////////////////////////////////////////////////////////////

func (z *Tokenizer) consumeCommentToken() bool {
	if z.r.Peek(0) != '/' || z.r.Peek(1) != '/' && z.r.Peek(1) != '*' {
		return false
	}
	if z.r.Peek(1) == '/' {
		z.r.Move(2)
		// single line
		for {
			c := z.r.Peek(0)
			if c == '\r' || c == '\n' || c == 0 {
				break
			} else if c >= 0xC0 {
				nOld := z.r.Pos()
				if z.consumeLineTerminator() {
					z.r.MoveTo(nOld)
					break
				}
				z.consumeRune()
				continue
			}
			z.r.Move(1)
		}
	} else {
		z.r.Move(2)
		// multi line
		for {
			c := z.r.Peek(0)
			if c == 0 {
				break
			} else if c == '*' && z.r.Peek(1) == '/' {
				z.r.Move(2)
				return true
			} else if c >= 0xC0 {
				z.consumeRune()
				continue
			}
			z.r.Move(1)
		}
	}
	return true
}

func (z *Tokenizer) consumeLongPunctuatorToken() bool {
	c := z.r.Peek(0)
	if c == '!' || c == '=' || c == '+' || c == '-' || c == '*' || c == '/' || c == '%' || c == '&' || c == '|' || c == '^' {
		z.r.Move(1)
		if z.r.Peek(0) == '=' {
			z.r.Move(1)
			if (c == '!' || c == '=') && z.r.Peek(0) == '=' {
				z.r.Move(1)
			}
		} else if (c == '+' || c == '-' || c == '&' || c == '|') && z.r.Peek(0) == c {
			z.r.Move(1)
		}
		return true
	}
	if c == '<' || c == '>' {
		z.r.Move(1)
		if z.r.Peek(0) == c {
			z.r.Move(1)
			if c == '>' && z.r.Peek(0) == '>' {
				z.r.Move(1)
			}
		}
		if z.r.Peek(0) == '=' {
			z.r.Move(1)
		}
		return true
	}
	return false
}

func (z *Tokenizer) consumeIdentifierToken() bool {
	c := z.r.Peek(0)
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '$' || c == '_' {
		z.r.Move(1)
	} else if c >= 0xC0 {
		if r := z.r.PeekRune(0); unicode.IsOneOf(identifierStart, r) {
			z.consumeRune()
		} else {
			return false
		}
	} else if !z.consumeUnicodeEscape() {
		return false
	}
	for {
		c := z.r.Peek(0)
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '$' || c == '_' {
			z.r.Move(1)
		} else if c >= 0xC0 {
			if r := z.r.PeekRune(0); r == '\u200C' || r == '\u200D' || unicode.IsOneOf(identifierPart, r) {
				z.consumeRune()
			} else {
				break
			}
		} else {
			break
		}
	}
	return true
}

func (z *Tokenizer) consumeNumericToken() bool {
	nOld := z.r.Pos()
	c := z.r.Peek(0)
	firstDigit := false
	if firstDigit = z.r.Peek(0) == '0'; firstDigit {
		z.r.Move(1)
		if z.r.Peek(0) == 'x' || z.r.Peek(0) == 'X' {
			z.r.Move(1)
			if z.consumeHexDigit() {
				for z.consumeHexDigit() {
				}
			} else {
				z.r.Move(-1) // return just the zero
			}
			return true
		}
		firstDigit = true
	} else if firstDigit = z.consumeDigit(); firstDigit {
		for z.consumeDigit() {
		}
	}
	if z.r.Peek(0) == '.' {
		z.r.Move(1)
		if z.consumeDigit() {
			for z.consumeDigit() {
			}
		} else if firstDigit {
			// . could belong to the next token
			z.r.Move(-1)
			return true
		} else {
			z.r.MoveTo(nOld)
			return false
		}
	} else if !firstDigit {
		z.r.MoveTo(nOld)
		return false
	}
	nOld = z.r.Pos()
	c = z.r.Peek(0)
	if c == 'e' || c == 'E' {
		z.r.Move(1)
		c = z.r.Peek(0)
		if c == '+' || c == '-' {
			z.r.Move(1)
		}
		if !z.consumeDigit() {
			// e could belong to the next token
			z.r.MoveTo(nOld)
			return true
		}
		for z.consumeDigit() {
		}
	}
	return true
}

func (z *Tokenizer) consumeStringToken() bool {
	delim := z.r.Peek(0)
	if delim != '"' && delim != '\'' {
		return false
	}
	nOld := z.r.Pos()
	z.r.Move(1)
	for {
		if z.consumeLineTerminator() {
			z.r.MoveTo(nOld)
			return false
		}
		c := z.r.Peek(0)
		if c == 0 {
			break
		} else if c == delim {
			z.r.Move(1)
			break
		} else if c == '\\' {
			if !z.consumeEscape() {
				break
			}
			continue
		} else if c >= 0xC0 {
			z.consumeRune()
			continue
		}
		z.r.Move(1)
	}
	return true
}

func (z *Tokenizer) consumeRegexpToken() bool {
	if z.r.Peek(0) != '/' || z.r.Peek(1) == '*' {
		return false
	}
	nOld := z.r.Pos()
	z.r.Move(1)
	inClass := false
	for {
		if z.consumeLineTerminator() {
		}
		c := z.r.Peek(0)
		if c == 0 {
			break
		} else if !inClass && c == '/' {
			z.r.Move(1)
			break
		} else if c == '\\' {
			if z.consumeLineTerminator() {
				z.r.MoveTo(nOld)
				return false
			}
			z.r.Move(1)
		} else if c == '[' {
			inClass = true
		} else if c == ']' {
			inClass = false
		} else if c >= 0xC0 {
			z.consumeRune()
			continue
		}
		z.r.Move(1)
	}
	// flags
	for {
		c := z.r.Peek(0)
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '$' || c == '_' {
			z.r.Move(1)
		} else if c >= 0xC0 {
			if r := z.r.PeekRune(0); r == '\u200C' || r == '\u200D' || unicode.IsOneOf(identifierPart, r) {
				z.consumeRune()
			} else {
				break
			}
		} else {
			break
		}
	}
	return true
}