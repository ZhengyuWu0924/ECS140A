package sexpr

import (
	"errors"
	"math/big" // You will need to use this package in your implementation.
)

// ErrEval is the error value returned by the Evaluator if the contains
// an invalid token.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrEval = errors.New("eval error")

func (expr *SExpr) Eval() (*SExpr, error) {
	if expr == nil {
		return nil, ErrEval
	}
	if expr.atom.typ == tokenNumber {
		if expr.cdr == nil {
			return expr, nil
		} else {
			tok := &token{typ: tokenNumber, num: expr.atom.num}
			return &SExpr{atom: tok, car: nil, cdr: nil}, nil
		}
	}
	if expr.atom.typ == tokenSymbol {
		if expr.atom.literal == "+" {
			//(+)
			if expr.cdr != nil {
				total := big.NewInt(0)
				if expr.car.isAtom() == false {
					res,err := expr.car.Eval() 
					if err != nil || res.atom.typ != tokenNumber {
						return nil, err
					}
					total.Add(total, res.atom.num)
				} else {					
					newExpr := expr.cdr
					for newExpr != nil {
						if newExpr.atom == nil {
							break;
						}
						res,err := newExpr.Eval() 
						if err != nil || res.atom.typ != tokenNumber {
							return nil, err
						}
						total.Add(total, res.atom.num)
						newExpr = newExpr.cdr
					}
				}
				tok := &token{typ: tokenNumber, num: total}
				return &SExpr{atom: tok, car: nil, cdr: nil}, nil
			} else {
				//+
				return nil, ErrEval
			}
		}
		if expr.atom.literal == "*" {
			if expr.cdr != nil {
				total := big.NewInt(1)
				if expr.car.isAtom() == false {
					res,err := expr.car.Eval() 
					if err != nil || res.atom.typ != tokenNumber {
						return nil, err
					}
					total.Mul(total, res.atom.num)
				} else {					
					newExpr := expr.cdr
					for newExpr != nil {
						if newExpr.atom == nil {
							break;
						}
						res,err := newExpr.Eval() 
						if err != nil || res.atom.typ != tokenNumber {
							return nil, err
						}
						total.Mul(total, res.atom.num)
						newExpr = newExpr.cdr
					}
				}
				tok := &token{typ: tokenNumber, num: total}
				return &SExpr{atom: tok, car: nil, cdr: nil}, nil
			} else {
				return nil, ErrEval
			}
		}
		if expr.atom.literal == "CAR" {
			expr = expr.cdr
			if expr == nil {
				return nil, ErrEval
			}
			expr = expr.car
			if expr == nil {
				return nil, ErrEval
			}
			if expr.atom.typ == tokenQuote {
				expr = expr.cdr
				res := expr
				for expr != nil {
					res = expr
					expr = expr.car
				}
				return res, nil
			}
			return expr, nil
		}
		if expr.atom.literal == "CDR" {
			expr = expr.cdr
			if expr == nil {
				return nil, ErrEval
			}
			expr = expr.car
			if expr == nil {
				return nil, ErrEval
			}
			if expr.atom.typ == tokenQuote {
				expr = expr.cdr
				res := expr
				for expr != nil {
					res = expr
					if expr.cdr != nil && expr.cdr.atom == nil {
						expr = expr.car
					} else {
						expr = expr.cdr
						if expr == nil {
							return nil, ErrEval
						} else {
							res = expr
							break
						}
					}
				}
				return res, nil
			}
			return expr, nil
		}
		if expr.atom.literal == "ATOM" {

		}
		if expr.atom.literal == "LISTP" {

		}
		if expr.atom.literal == "LENGTH" {
		}
		if expr.atom.literal == "CONS" {
			
		}
		if expr.atom.literal == "QUOTE" {
			
		}
		if expr.atom.literal == "ZEROP" {
		}
	} else {

	}
	return expr, nil
}

