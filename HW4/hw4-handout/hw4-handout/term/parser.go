package term

import (
	"errors"
	"strconv"
)

// use the released solution from instructor

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <term>     ::= ATOM | NUM | VAR | <compound>
// <compound> ::= <functor> LPAR <args> RPAR
// <functor>  ::= ATOM
// <args>     ::= <term> | <term> COMMA <args>
//

// Parser is the interface for the term parser.
// Do not change the definition of this interface.
type Parser interface {
	Parse(string) (*Term, error)
}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	return &ParserImpl{
		lex:         nil,
		peekTok:     nil,
		terms:       make(map[string]*Term),
		termID:      make(map[*Term]int),
		termCounter: 0,
	}
}

type ParserImpl struct {
	lex         *lexer
	peekTok     *Token
	terms       map[string]*Term
	termID      map[*Term]int
	termCounter int
}

func (p *ParserImpl) nextToken() (*Token, error) {
	if tok := p.peekTok; tok != nil {
		p.peekTok = nil
		return tok, nil
	}
	return p.lex.next()
}

func (p *ParserImpl) backToken(tok *Token) {
	p.peekTok = tok
}

func (p *ParserImpl) Parse(input string) (*Term, error) {
	p.lex = newLexer(input)
	p.peekTok = nil

	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}
	if tok.typ == tokenEOF {
		return nil, nil
	}
	p.backToken(tok)

	term, err := p.parseNextTerm()
	if err != nil {
		return nil, ErrParser
	}
	if tok, err := p.nextToken(); err != nil || tok.typ != tokenEOF {
		return nil, ErrParser
	}
	return term, nil
}

func (p *ParserImpl) parseNextTerm() (*Term, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, err
	}

	switch tok.typ {
	case tokenEOF:
		return nil, nil
	case tokenNumber:
		return p.mkSimpleTerm(TermNumber, tok.literal), nil
	case tokenVariable:
		return p.mkSimpleTerm(TermVariable, tok.literal), nil
	case tokenAtom:
		a := p.mkSimpleTerm(TermAtom, tok.literal)
		nxt, err := p.nextToken()
		if err != nil {
			return nil, err
		}
		if nxt.typ != tokenLpar {
			p.backToken(nxt)
			return a, nil
		}
		arg, err := p.parseNextTerm()
		if err != nil {
			return nil, err
		}
		args := []*Term{arg}
		nxt, err = p.nextToken()
		if err != nil {
			return nil, err
		}
		for ; nxt.typ == tokenComma; nxt, err = p.nextToken() {
			arg, err = p.parseNextTerm()
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		}
		if nxt.typ != tokenRpar {
			return nil, ErrParser
		}
		return p.mkCompoundTerm(a, args), nil
	default:
		return nil, ErrParser
	}
}

func (p *ParserImpl) mkSimpleTerm(typ TermType, lit string) *Term {
	key := lit
	term, ok := p.terms[key]
	if !ok {
		term = &Term{Typ: typ, Literal: lit}
		p.insertTerm(term, key)
	}
	return term
}

func (p *ParserImpl) mkCompoundTerm(functor *Term, args []*Term) *Term {
	key := strconv.Itoa(p.termID[functor])
	for _, arg := range args {
		key += "," + strconv.Itoa(p.termID[arg])
	}
	term, ok := p.terms[key]
	if !ok {
		term = &Term{
			Typ:     TermCompound,
			Functor: functor,
			Args:    args,
		}
		p.insertTerm(term, key)
	}
	return term
}

func (p *ParserImpl) insertTerm(term *Term, key string) {
	p.terms[key] = term
	p.termID[term] = p.termCounter
	p.termCounter++
}
