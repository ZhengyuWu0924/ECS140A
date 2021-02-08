package term

import (
	"errors"
	// "fmt"
)

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

type P struct{
	term *Term
	tokenSlice []*Token
	idx int
	isValid bool
	m map[string]*Term
}

type Parser interface {
	Parse(string) (*Term, error)
}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	p := &P{term: &Term{}, idx: 0, tokenSlice: make([]*Token, 0), isValid: true, m: make(map[string]*Term)}
	return p
}

func (p *P) isTermArgs() bool {
	leftMinusRight := 0
	for i := p.idx + 1; i < len(p.tokenSlice); i++ {
		if p.tokenSlice[i].typ == tokenComma && leftMinusRight == 0 {
			return true
		}
		if p.tokenSlice[i].typ == tokenLpar {
			leftMinusRight++
		} else if p.tokenSlice[i].typ == tokenRpar {
			leftMinusRight--
		}
	}
	return false
}

func (p *P) toDAG(term *Term) *Term {
	name := term.String()
	if val, ok := p.m[name]; ok{
		return val
	}
	p.m[name] = term
	if term.Typ == TermCompound{
		term.Functor = p.toDAG(term.Functor)
		for i := 0; i < len(term.Args); i++{
			term.Args[i] = p.toDAG(term.Args[i])
		}
	}
	return term
}

func (p *P) Parse(input string) (*Term, error) {
	if input == "" {
		return nil, nil
	}
	// p.tokenSlice = make([]*Token, 0)
	lex := newLexer(input)
	for i := 0; ; i++ {
		token, _ := lex.next()
		p.tokenSlice = append(p.tokenSlice, token)
		// fmt.Println("token type", token.typ)
		if token.typ == tokenEOF {
			break
		}
	}
	// p.term = &Term{}
	p.termF(p.term)
	// fmt.Println(p.idx)
	if p.tokenSlice[p.idx].typ == tokenEOF && p.isValid == true{
		p.toDAG(p.term)
		return p.term, nil
	}
	return nil, errors.New("Invalid Term")
}

func (p *P) termF(term *Term) {
	if p.tokenSlice[p.idx].typ == tokenAtom {
		if p.tokenSlice[p.idx + 1].typ == tokenLpar{
			term.Typ = TermCompound
			p.compoundF(term)
		} else {
			term.Typ = TermAtom
			// term.Literal = p.tokenSlice[p.idx].literal
			// p.m[term.Literal] = term
			p.match(tokenAtom, term)
		}
	} else if p.tokenSlice[p.idx].typ == tokenVariable {
		term.Typ = TermVariable
		// term.Literal = p.tokenSlice[p.idx].literal
		// p.m[term.Literal] = term
		// fmt.Println("line 89 matching")
		p.match(tokenVariable, term)
	} else if p.tokenSlice[p.idx].typ == tokenNumber {
		term.Typ = TermNumber
		// term.Literal = p.tokenSlice[p.idx].literal
		// p.m[term.Literal] = term
		// fmt.Println("line 94 matching")
		p.match(tokenNumber, term)
		// fmt.Println("line 96",p.idx)
	} else {
		p.isValid = false
	}
}

func (p *P) compoundF(term *Term) {
	term.Functor = &Term{}
	// newTerm := &Term{}
	// term.Args = append(term.Args, newTerm)
	if p.tokenSlice[p.idx].typ == tokenAtom {
		// term.Functor.Typ = TermAtom
		// term.Functor.Literal = p.tokenSlice[p.idx].literal
		// if val, ok := p.m[p.tokenSlice[p.idx].literal]; ok{
		// 	term.Functor = val
		// } else {
		// 	term.Functor = &Term{}
		// }
		p.functorF(term.Functor)
		// fmt.Println("line 108 matching")
		p.match(tokenLpar, term)
		p.argsF(term)
		// fmt.Println("line 111 matching")
		p.match(tokenRpar, term)
		// p.m[term.String()] = term
	}
}

func (p *P) functorF(term *Term) {
	term.Typ = TermAtom
	// fmt.Println("line 118 matching")
	p.match(tokenAtom, term)
}

func (p *P) argsF(term *Term) {	
	newTerm := &Term{}
	term.Args = append(term.Args, newTerm)
	if p.isTermArgs() {
		p.termF(newTerm)
		p.match(tokenComma, term)
		p.argsF(term)
	} else {
		p.termF(newTerm)
	}
}

func (p *P) match(typ tokenType, term *Term) {
	// fmt.Println(typ, p.tokenSlice[p.idx].typ)
	if p.tokenSlice[p.idx].typ == typ {
		switch typ {
		case tokenAtom, tokenNumber, tokenVariable:
			term.Literal = p.tokenSlice[p.idx].literal
		}
		p.idx++
		// fmt.Println(p.idx)
	} else {
		p.isValid = false		
	}
}

