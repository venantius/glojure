package lang

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
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

var UNKNOWN *Keyword = InternKeywordByNsName("unknown")

var macros map[rune]IFn = map[rune]IFn{
	'"': &StringReader{},
	';': &CommentReader{},
	'\'': &WrappingReader{sym: QUOTE},
	'@': &WrappingReader{sym: DEREF},
	'^': &MetaReader{},
	'`': &SyntaxQuoteReader{},
	'~': &UnquoteReader{},
	'(': &ListReader{},
	')': &UnmatchedDelimiterReader{},
	'[': &VectorReader{},
	']': &UnmatchedDelimiterReader{},
	'{': &MapReader{},
	'}': &UnmatchedDelimiterReader{},
	'\\': &RuneReader{},
	'%': &ArgReader{},
	'#': &DispatchReader{},
}

var dispatchMacros map[rune]IFn = map[rune]IFn{
	'^': &MetaReader{},
	'\'': &VarReader{},
	'"': &RegexReader{},
	'(': &FnReader{},
	'{': &SetReader{},
	'=': &EvalReader{},
	'!': &CommentReader{},
	'<': &UnreadableReader{},
	'_': &DiscardReader{},
	'?': &ConditionalReader{},
}

var symbolPat *regexp.Regexp = regexp.MustCompile("[:]?([\\D&&[^/]].*/)?(/|[\\D&&[^/]][^/]*)")
var intPat *regexp.Regexp = regexp.MustCompile("([-+]?)(?:(0)|([1-9][0-9]*)|0[xX]([0-9A-Fa-f]+)|0([0-7]+)|([1-9][0-9]?)[rR]([0-9A-Za-z]+)|0[0-9]+)(N)?")
var radioPat *regexp.Regexp = regexp.MustCompile("([-+]?[0-9]+)/([0-9]+)")
var floatPat *regexp.Regexp = regexp.MustCompile("([-+]?[0-9]+(\\.[0-9]*)?([eE][-+]?[0-9]+)?)(M)?")

// TODO: Var GENSYM_ENV
// TODO: Var ARG_ENV

// TODO: var ctorReader IFn = CtorReader{}
// TODO: Var READ_COND_ENV

// NOTE: isWhiteSpace => unicode.isSpace(ch)

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

func (lr *LispReader) Unread() {
	err := lr.r.UnreadRune()
	if err != nil {
		Util.SneakyThrow(err)
	}
}

// Reader opts
var OPT_EOF *Keyword = InternKeywordByNsName("eof")



// TODO: ReaderException

type RegexReader struct {
	AFn
}

// todo
func (r *RegexReader) Invoke(args ...interface{}) interface{} {
	reader, doublequote, opts, pendingForms := unpackReaderArgs(args)
	return nil
}

type StringReader struct {
	AFn
}

func (s *StringReader) Invoke(args ...interface{}) interface{} {
	reader, doublequote, opts, pendingForms := unpackReaderArgs(args)
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

func (cr *CommentReader) Invoke(args ...interface{}) interface{} {
	reader, semicolon, opts, pendingForms := unpackReaderArgs(args)
	r := &LispReader{r: bufio.NewReader(reader.(io.Reader))}
	var ch int
	for ch := r.Read(); ch != '\n' && ch != '\r' && ch != "-1"; ch = r.Read() {
		// Advance the reader through comments
	}
	return r
}

type DiscardReader struct {
	AFn
}

// TODO
func (dr *DiscardReader) Invoke(args ...interface{}) interface{} {
	reader, underscore, opts, pendingForms := unpackReaderArgs(args)
	return nil
}

type WrappingReader struct {
	AFn

	sym *Symbol
}

// TODO
func (wr *WrappingReader) Invoke(args ...interface{}) interface{} {
	reader, quote, opts, pendingforms := unpackReaderArgs(args)
	return nil
}

// TODO: Many more readers.

type VarReader struct {
	AFn
}

// TODO
func (vr *VarReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type DispatchReader struct {
	AFn
}

// TODO
func (dr *DispatchReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type FnReader struct {
	AFn
}
// TODO
func (fr *FnReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type ArgReader struct {
	AFn
}
// TODO
func (ar *ArgReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type MetaReader struct {
	AFn
}

// TODO
func (mr *MetaReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type SyntaxQuoteReader struct {
	AFn
}

// TODO
func (sr *SyntaxQuoteReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type UnquoteReader struct {
	AFn
}

// TODO
func (ur *UnquoteReader) Invoke(args ...interface{}) interface{} {
	return nil
}

/*
	RuneReader [CharacterReader]
 */

type RuneReader struct {
	AFn
}

// TODO
func (cr *RuneReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type ListReader struct {
	AFn
}

// TODO
func (lr *ListReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type EvalReader struct {
	AFn
}

// TODO
func (er *EvalReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type VectorReader struct {
	AFn
}

// TODO
func (vr *VectorReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type MapReader struct {
	AFn
}

// TODO
func (mr *MapReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type SetReader struct {
	AFn
}

// TODO
func (sr *SetReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type UnmatchedDelimiterReader struct {
	AFn
}

// TODO
func (udr *UnmatchedDelimiterReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type UnreadableReader struct {
	AFn
}

// TODO
func (ur *UnreadableReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type CtorReader struct {
	AFn
}

// TODO
func (cr *CtorReader) Invoke(args ...interface{}) interface{} {
	return nil
}

type ConditionalReader struct {
	AFn
}

// TODO
func (cr *ConditionalReader) Invoke(args ...interface{}) interface{} {
	return nil
}


/*
	Static methods
 */

func unpackReaderArgs(args []interface{}) (interface{}, interface{}, interface{}, interface{}) {
	a, b, c, d := args[0], args[1], args[2], args[3]
	return a, b, c, d
}