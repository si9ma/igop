package constant

import (
	"fmt"
	"go/constant"
	"go/token"
	"math/big"
	"strings"
)

func ExactConstant(c constant.Value) (s string) {
	s, _ = ExactConstantEx(c, false)
	return
}

func ExactConstantEx(c constant.Value, toFloat bool) (s string, exact bool) {
	switch c.Kind() {
	case constant.Bool:
		if constant.BoolVal(c) {
			return "true", true
		} else {
			return "false", true
		}
	case constant.String:
		return constant.StringVal(c), true
	case constant.Int:
		return c.ExactString(), true
	case constant.Float:
		s := c.ExactString()
		if pos := strings.IndexByte(s, '/'); pos >= 0 {
			sx := s[:pos]
			sy := s[pos+1:]
			// simplify 314/100 => 3.14
			// 80901699437494742410229341718281905886015458990288143106772431
			// 50000000000000000000000000000000000000000000000000000000000000
			if strings.HasPrefix(sy, "1") && strings.Count(sy, "0") == len(sy)-1 {
				n := len(sy) - len(sx)
				if n == 0 {
					return fmt.Sprintf("%v.%v", sx[:1], sx[1:]), true
				} else if n > 0 {
					if n < 5 {
						return fmt.Sprintf("0.%v%v", strings.Repeat("0", n-1), sx), true
					}
					if len(sx[1:]) == 0 {
						return fmt.Sprintf("%ve-%v", sx[:1], len(sy)-len(sx)), true
					}
					return fmt.Sprintf("%v.%ve-%v", sx[:1], sx[1:], len(sy)-len(sx)), true
				}
			} else if strings.HasPrefix(sy, "5") && strings.Count(sy, "0") == len(sy)-1 {
				if len(sx) == len(sy) {
					c := constant.BinaryOp(constant.MakeFromLiteral(sx, token.INT, 0), token.MUL, constant.MakeInt64(2))
					sx = c.ExactString()
					return fmt.Sprintf("%v.%v", sx[:1], sx[1:]), true
				}
			}
			if toFloat {
				r, ok := new(big.Rat).SetString(s)
				if !ok {
					panic(fmt.Errorf("parser rat %q error", s))
				}
				if v, ok := r.Float64(); ok {
					return fmt.Sprintf("%v", v), false
				}
				return r.FloatString(10), false
			}
			return s, true
		}
		if pos := strings.LastIndexAny(s, "123456789"); pos != -1 {
			sx := s[:pos+1]
			if len(s)-1 > 5 {
				if len(sx[1:]) == 0 {
					return fmt.Sprintf("%ve+%v", sx[:1], len(s)-1), true
				}
				return fmt.Sprintf("%v.%ve+%v", sx[:1], sx[1:], len(s)-1), true
			}
		}
		return s, true
	case constant.Complex:
		re, e1 := ExactConstantEx(constant.Real(c), toFloat)
		im, e2 := ExactConstantEx(constant.Imag(c), toFloat)
		return fmt.Sprintf("%v+%vi", re, im), e1 && e2
	default:
		panic("unreachable")
	}
}
