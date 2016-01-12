package lang
import (
	"bufio"
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

type StringReader struct {
	AFn
}

func (s *StringReader) Invoke(reader interface{}, doublequote interface{}, opts interface{}, pendingForms interface{}) interface{} {
	// sb := StringBuilder{} // TODO: Need my own stringbuilder type thing
	r := bufio.NewScanner(reader.(io.Reader)) // TODO: I'm cheating with this right now.

	for ch := LispReader.Read1(r); ch != "\""; ch = LispReader.Read1(r) {
		if ch == "-1" {
			panic("EOF while reading string")
		}
		if ch == "\\" {
			ch = LispReader.Read1(r)
			if ch == "-1" {
				panic("EOF while reading string")

			}
			switch ch {
			case "t":
				ch = "\t"
			case "r":
				ch = "\r"
			case "n":
				ch = "\n"
			case "\\":
				break
			case "\"": // TODO: verify that this actually is the correct way of handling things
				break
			case "b":
				ch = "\b"
			case "f":
				ch = "\f"
			case "u":
				ch = LispReader.Read1(r)
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
	}
	// TODO: the rest of this. I got stressed.
	// return sb.String()
	return nil

}

/*
	Static methods
 */

type lr struct {}
var LispReader = lr{}

func (lr *lr) Read1(s *bufio.Scanner) string {
	if s.Scan() {
		if err := s.Err(); err != nil {
			Util.SneakyThrow(s.Err())
		}
		return s.Text()
	}
	return "-1" // We've reached the end of the scanner.
}

// TODO
func (lr *lr) ReadUnicodeChar(s *bufio.Scanner, initch int, base int, length int, exact bool) int {
	return 0
}