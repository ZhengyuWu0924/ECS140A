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
var flag = 0

func (expr *SExpr) Eval() (*SExpr, error) {
	if expr == nil {
		return nil, ErrEval
	}
	if (expr.atom == nil || expr.atom.literal == "NIL") && expr.car == nil && expr.cdr == nil {
		return expr, nil
	}
	if expr.atom.typ == tokenNumber {
		if flag == 0 && expr.cdr != nil {
			return nil, ErrEval
		}
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
						return nil, ErrEval
					}
					total.Add(total, res.atom.num)
				} else {					
					newExpr := expr.cdr
					if newExpr.car != nil && newExpr.car.atom.typ != tokenNumber {
						flag++
						res,err := newExpr.car.Eval()
						flag-- 
						if err != nil || res.atom.typ != tokenNumber {
							return nil, ErrEval
						}
						total.Add(total, res.atom.num)
						newExpr = newExpr.cdr
					}
					for newExpr != nil {
						if newExpr.atom == nil {
							break;
						}
						flag++
						res,err := newExpr.Eval()
						flag-- 
						if err != nil || res.atom.typ != tokenNumber {
							return nil, ErrEval
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
					flag++
					res,err := expr.car.Eval()
					flag-- 
					if err != nil || res.atom.typ != tokenNumber {
						return nil, ErrEval
					}
					total.Mul(total, res.atom.num)
				} else {					
					newExpr := expr.cdr
					if newExpr.car != nil && newExpr.car.atom.typ != tokenNumber {
						flag++
						res,err := newExpr.car.Eval()
						flag-- 
						if err != nil || res.atom.typ != tokenNumber {
							return nil, ErrEval
						}
						total.Mul(total, res.atom.num)
						newExpr = newExpr.cdr
					}
					for newExpr != nil {
						if newExpr.atom == nil {
							break;
						}
						flag++
						res,err := newExpr.Eval()
						flag-- 
						if err != nil || res.atom.typ != tokenNumber {
							return nil, ErrEval
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
			if expr == nil || expr.isNil() == true || expr.cdr.atom != nil {
				return nil, ErrEval
			}
			if expr.atom.literal == "NIL" && expr.cdr != nil && expr.cdr.isNil() == true {
				return expr.cdr, nil
			}
			expr = expr.car
			if expr == nil || expr.cdr == nil || expr.atom.typ != tokenQuote {
				return nil, ErrEval
			}
			expr = expr.cdr
			res := expr
			for expr != nil {
				res = expr
				expr = expr.car
			}
			return res, nil
		}
		if expr.atom.literal == "CDR" {
			expr = expr.cdr
			if expr == nil || expr.isNil() == true {
				return nil, ErrEval
			}
			if expr.atom.literal == "NIL" && expr.cdr != nil && expr.cdr.isNil() == true {
				return expr.cdr, nil
			}
			expr = expr.car
			if expr == nil || expr.atom.typ != tokenQuote {
				return nil, ErrEval
			}
			expr = expr.cdr
			res := expr
			for expr != nil {
				res = expr
				if expr.cdr != nil && expr.cdr.atom == nil {
					expr = expr.car
				} else {
					expr = expr.cdr
					if expr != nil {
						res = expr
						break
					} 
				}
			}
			return res, nil
			
		}
		if expr.atom.literal == "ATOM" {
			expr = expr.cdr
			if expr.cdr == nil ||  expr.isNil() == true{
				return nil, ErrEval
			}
			if expr.cdr.isNil() == false && expr.cdr.atom.literal != "NIL" {
				return nil, ErrEval
			}
			if expr != nil {
				expr = expr.car
				flag++
				res, err := expr.Eval()
				flag-- 
				if err != nil {
					return nil, ErrEval
				}
				if res.isAtom() == true {
					tok := &token{typ: tokenSymbol, literal: "T"}
					return &SExpr{atom: tok, car: nil, cdr: nil}, nil
				} else {
					return &SExpr{atom: nil, car: nil, cdr: nil}, nil
				}
			}
		}
		if expr.atom.literal == "LISTP" {
			expr = expr.cdr 
			if expr == nil || expr.atom == nil && expr.car == nil && expr.cdr == nil {
				return nil, ErrEval
			}
			if expr.cdr.atom != nil && expr.cdr.atom.literal != "NIL" {
				return nil, ErrEval
			}
			expr = expr.car
			flag++
			res, err := expr.Eval()
			flag--
			if err != nil {
				return nil, ErrEval
			}
			//NIL
			if (res.atom == nil || res.atom.literal == "NIL") && res.car == nil && res.cdr == nil  {
				tok := &token{typ: tokenSymbol, literal: "T"}
				return &SExpr{atom: tok, car: nil, cdr: nil}, nil
			}
			if res.cdr == nil {
				return &SExpr{atom: nil, car: nil, cdr: nil}, nil
			} else {
				tok := &token{typ: tokenSymbol, literal: "T"}
				return &SExpr{atom: tok, car: nil, cdr: nil}, nil
			}
		}
		if expr.atom.literal == "LENGTH" {
			expr = expr.cdr
			if expr == nil || expr.atom == nil || expr.atom.typ != tokenQuote {
				return nil, ErrEval
			}
			expr = expr.car
			if expr != nil {expr = expr.cdr
				if expr.car != nil {
					return expr.car.getLength();
				}
			}
		}
		if expr.atom.literal == "CONS" {
			expr = expr.cdr
			if expr == nil || expr.car == nil {
				return nil, ErrEval
			}
			flag++
			res1, err := expr.car.Eval()
			flag--
			if err != nil {
				return nil, ErrEval
			}
			expr = expr.cdr 
			if expr == nil || expr.cdr.atom != nil{
				return nil, ErrEval
			}
			expr = expr.car 
			flag++
			res2, err := expr.Eval()
			flag--
			if err != nil {
				return nil, ErrEval
			}
			return &SExpr{atom: res1.atom, car: res1, cdr: res2}, nil
		}
		if expr.atom.literal == "ZEROP" {
			expr = expr.cdr
			if expr == nil || expr.cdr == nil {
				return nil, ErrEval
			}
			if expr.cdr.isNil() == false && expr.cdr.atom.literal != "NIL" {
				return nil, ErrEval
			}
			if expr.atom == nil {
				return nil, ErrEval
			}
			if expr.car.atom.typ == tokenNumber && expr.car.isAtom() == false {
				return nil, ErrEval
			}
			flag++
			res, err := expr.Eval()
			flag--
			if err != nil {
				return nil, ErrEval
			}
			if res.atom.typ == tokenNumber && res.atom.num.Cmp(big.NewInt(0)) == 0 {
				tok := &token{typ: tokenSymbol, literal: "T"}
				return &SExpr{atom: tok, car: nil, cdr: nil}, nil
			} else {
				return &SExpr{atom: nil, car: nil, cdr: nil}, nil
			} 
		}
		if expr.atom.literal != "QUOTE" {
			return nil, ErrEval
		}
	} 
	if expr.atom.literal == "QUOTE" {
		expr = expr.cdr
		if expr == nil || expr.cdr == nil {
			return nil, ErrEval
		}
		if expr.cdr.atom != nil && expr.cdr.atom.literal != "NIL" {
			return nil, ErrEval
		}
		expr = expr.car
		if expr != nil {
			return expr, nil
		} 
	}
	return nil, ErrEval
}

func (expr *SExpr) getLength() (*SExpr, error) {
	if expr.atom == nil && expr.car == nil && expr.cdr == nil {
		tok := &token{typ: tokenNumber, num: big.NewInt(0)}
		return &SExpr{atom: tok, car: nil, cdr: nil}, nil 
	}
	length := big.NewInt(1)
	expr = expr.cdr
	if expr == nil {
		return nil, ErrEval
	}
	res, err := expr.getLength();
	if err != nil {
		return nil, ErrEval
	}
	length.Add(length, res.atom.num)
	tok := &token{typ: tokenNumber, num: length}
	return &SExpr{atom: tok, car: nil, cdr: nil}, nil 
}