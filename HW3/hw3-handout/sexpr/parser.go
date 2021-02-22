package sexpr

import "errors"

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <sexpr>       ::= <atom> | <pars> | QUOTE <sexpr>
// <atom>        ::= NUMBER | SYMBOL
// <pars>        ::= LPAR <dotted_list> RPAR | LPAR <proper_list> RPAR
// <dotted_list> ::= <proper_list> <sexpr> DOT <sexpr>
// <proper_list> ::= <sexpr> <proper_list> | \epsilon
//
type Parser interface {
	Parse(string) (*SExpr, error)
}

func NewParser() Parser {
	return &ParserImpl{
		lex: nil,
		peekTok: nil,

	}
}

type ParserImpl struct {
	lex *lexer
	peekTok *token

}

func (p *ParserImpl) nextToken() (*token, error) {
	if tok := p.peekTok; tok != nil {
		p.peekTok = nil
		return tok, nil
	}
	return p.lex.next()
}

func (p *ParserImpl) backToken(tok *token) {
	p.peekTok = tok
}

func (p *ParserImpl) Parse(input string) (*SExpr, error) {
	p.lex = newLexer(input)
	p.peekTok = nil

	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}
	if tok.typ == tokenEOF {
		return nil, ErrParser
	}
	p.backToken(tok)
	sexpr, err := p.parseNextSexpr()
	if err != nil {
		return nil, ErrParser
	}
	if tok, err := p.nextToken(); err != nil || tok.typ != tokenEOF {
		return nil, ErrParser
	}
	return sexpr, nil
}

func (p *ParserImpl) parseNextSexpr() (*SExpr, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}
	switch tok.typ {
	case tokenEOF:
		return nil, nil
	case tokenNumber:
		return p.mkSimpleSexpr(tok, nil, nil), nil
	case tokenSymbol:
		return p.mkSimpleSexpr(tok, nil, nil), nil
	case tokenQuote:
		a := p.mkSimpleSexpr(tok, nil, nil)
		nxt, err := p.nextToken()
		if err != nil {
			return nil, ErrParser
		}
		if nxt.typ == tokenRpar {
			return nil, ErrParser
		}
		if nxt.typ == tokenEOF {
			return nil, ErrParser
		}
		p.backToken(nxt)
		arg, err := p.parseNextSexpr()
		if err != nil {
			return nil, err
		}
		nilArg := p.mkSimpleSexpr(nil, nil, nil)
		return p.mkSexpr(a.atom, a, p.mkSexpr(arg.atom, arg, nilArg)), nil

	case tokenLpar:
		nxt, err := p.nextToken()
		if err != nil {
			return nil, err
		}
		if nxt.typ == tokenEOF {
			return nil, ErrParser
		}
		if nxt.typ == tokenRpar {
			//epsilon
			return p.mkSimpleSexpr(nil, nil, nil), nil
		}
		p.backToken(nxt)
		arg, err := p.parseNextSexpr()
		if err != nil {
			return nil, err
		}
		args := []*SExpr{arg}
		//true means it's proper
		flags := []bool{true}
		nxt, err = p.nextToken() 
		if err != nil {
			return nil, ErrParser
		}
		for ; nxt.typ != tokenRpar && nxt.typ != tokenEOF; nxt, err = p.nextToken() {
			if nxt.typ != tokenDot {
				p.backToken(nxt)
				arg, err = p.parseNextSexpr()
				if err != nil {
					return nil, err
				}
				args = append(args, arg)
				flags = append(flags, true)
			} else {
				//if it's dotted, remove the last one from the args
				arg, err = p.parseNextSexpr()
				if err != nil {
					return nil, err
				}
				if arg == nil || (arg.atom == nil && arg.cdr == nil) {
					return nil, ErrParser
				}
				temp := args[len(args)-1]
				args = args[:len(args)-1]
				flags = flags[:len(flags)-1]
				args = append(args, p.mkSexpr(temp.atom, temp, arg))
				flags = append(flags, false)
			}
		}
		if err != nil || nxt.typ != tokenRpar {
			return nil, ErrParser
		}
		if len(args) == 1 {
			if flags[len(flags)-1] == true {
				arg = p.mkSimpleSexpr(nil, nil, nil)
				return p.mkSexpr(args[0].atom, args[0], arg), nil
			} else {
				return args[0], nil
			}
		}
		if flags[len(flags)-1] == true {
			arg = p.mkSimpleSexpr(nil, nil, nil)
			args[len(args)-1] = p.mkSexpr(args[len(args)-1].atom, args[len(args)-1], arg)
		} 
		for i := len(args)-2; i >= 1; i-- {
			if flags[i] == true {
				args[i] = p.mkSexpr(args[i].atom, args[i], args[i+1])
			} else {
				if flags[i+1] == false && args[i].atom == nil && args[i+1].atom == nil {
					return nil, ErrParser
				}
			}
		} 
		return p.mkSexpr(args[0].atom, args[0], args[1]), nil
	default:
		return nil, ErrParser

	}
}

func (p *ParserImpl) mkSimpleSexpr(tok *token, car *SExpr, cdr *SExpr) *SExpr {
	return &SExpr{atom: tok, car: car, cdr: cdr}
}

func (p *ParserImpl) mkSexpr(tok *token, car *SExpr, cdr *SExpr) *SExpr {
	return &SExpr{atom: tok, car: car, cdr: cdr}
}
