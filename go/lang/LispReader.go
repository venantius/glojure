package lang
import (
	"bufio"
	"bytes"
	"io"
)

var QUOTE *Symbol = InternSymbol("quote")
var THE_VAR *Symbol = InternSymbol("var")
var UNQUOTE *Symbol = InternSymbol("clojure.core", "unquote")
var UNQUOTE_SPLICING *Symbol = InternSymbol("clojure.core", "unqoute-splicing")
var CONCAT *Symbol = InternSymbol("clojure.core", "concat")
var SEQ *Symbol = InternSymbol("clojure.core", "seq")
var LIST *Symbol = InternSymbol("clojure.core", "list")
var APPLY *Symbol = InternSymbol("clojure.core", "apply")
var HASHMAP *Symbol = InternSymbol("clojure.core", "hash-map")
var HASHSET *Symbol = InternSymbol("clojure.core", "hash-set")
var VECTOR *Symbol = InternSymbol("clojure.core", "vector")
var WITH_META *Symbol = InternSymbol("clojure.core", "with-meta")
var META *Symbol = InternSymbol("clojure.core", "meta")
var DEREF *Symbol = InternSymbol("clojure.core", "deref")
var READ_COND *Symbol = InternSymbol("clojure.core", "read-cond")
var READ_COND_SPLICING *Symbol = InternSymbol("clojure.core", "read-cond-splicing")

var unknown = "unknown"
var UNKNOWN *Keyword = InternKeywordNsAndName(nil, &unknown) // TODO: Might want to change this to not take pointers later.

var macros []IFn = make([]IFn, 256)
var dispatchMacros []IFn = make([]IFn, 256)

// TODO: declare symbolPat
// TODO: declare intPat
// TODO: declare ratioPat
// TODO: declare floatPat

// TODO: A large block of code here

type LispReader struct{
	r *bufio.Reader
}

func (lr *LispReader) Read() rune {
	ch, _, err := lr.r.ReadRune()
	if err != nil {
		Util.SneakyThrow(err)
	}
	return ch
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// TODO: ReaderException, RegexReader

type StringReader struct {
	AFn
}

func (s *StringReader) Invoke(reader interface{}, doublequote interface{}, opts interface{}, pendingForms interface{}) interface{} {
	var sb bytes.Buffer
	r := &LispReader{r: bufio.NewReader(reader.(io.Reader))} // TODO: is casting reader to io.Reader legit?

	for ch := r.Read(); ch != '\\'; ch = r.Read() {
		if ch == "-1" {
			panic("EOF while reading string")
		}
		if ch == '\\' {
			ch = r.Read()
			if ch == "-1" {
				panic("EOF while reading string")

			}
			switch ch {
			case 't':
				ch = '\t'
			case 'r':
				ch = '\r'
			case 'n':
				ch = '\n'
			case '\\':
				break
			case '"':
				break
			case 'b':
				ch = '\b'
			case 'f':
				ch = '\f'
			case 'u':
				ch = r.Read()
				// TODO
				/*
				if Character.digit(ch, 16) == -1 {
					panic("Invalid unicode escape") // TODO: flesh this out more
				}
				ch = LispReader.ReadUnicodeChar(r, ch, 16, 4, true)
				*/
			default:
				// TODO
				/*
				if(Character.isDigit(ch)) {
					ch = LispReader.ReadUnicodeChar(r, ch, 8, 3, false);
					if(ch > 0377) {
						panic("Octal escape sequence must be in range [0, 377].");
					} else {}
					panic("Unsupported escape character") // TODO: Flesh this out more
				}
				*/
			}
		}
		sb.WriteRune(ch)
	}
	return sb.String()

}

type CommentReader struct {
	AFn
}

func (cr *CommentReader) Invoke(reader interface{}, semicolon interface{}, opts interface{}, pendingForms interface{}) interface{} {
	r := &LispReader{r: bufio.NewReader(reader.(io.Reader))}
	var ch int
	for ch := r.Read(); ch != '\n' && ch != '\r' && ch != "-1"; ch = r.Read() {
		// Advance the reader through comments
	}
	return r
}

// TODO: Many more readers.


/*
	Static methods
 */