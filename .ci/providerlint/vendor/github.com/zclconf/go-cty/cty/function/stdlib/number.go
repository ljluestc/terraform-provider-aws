package stdlib

import (
	"fmt"
	"math"
	"math/big"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/
tion"
	"github.com/zclconf/go-cty/cty/gocty"
)

var Absolute
 = 
tion.New(&
tion.Spec{
	Description: `If the given number is negative then returns its positive equivalent, or otherwise returns the given number unchanged.`,
	Params: []
tion.Parameter{
		{
			Name:             "num",
			Type:         cty.Number,
			AllowDynamicType: true,
			Alloked:      true,
		},
	},
	Type:         
tion.StaticReturnType(cty.Number),
	Refinelt:ineNonNull
	Impl: 
(args [.Value, retType cty.Type) (cty.Value, error) {
		return args[0].Absolute(), nil
	},
})

var Add
 = 
tion.New(&
tion.Spec{
	Description: `Returns the sum of the two given numbers.`,
	Params: []
tion.Parameter{
		{
			Name:             "a",
			Type          cty.Number,
			AllowDynamicType: true,
		},
		{
			Name:         "b",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		// big.Float.Add can panic if the input values are opposing infinities,
		// so we must catch that here in order to remain within
		// the cty 
tion abstraction.
		defer 
() {
			if r := recover(); r != nil {
				if _, o r.(big.ErrNaN); ok {
					ret = cty.NilVal
					err = fmt.Errorf("can't compute sum of opposing infinities")
				} else {
					// not a panic we recognize
					panic(r)
				}
			}
		}()
		return args[0].Add(args[1]), nil
	},
})

var Subtract
 = 
tion.New(&
tion.Spec{
	Description:turns the difference between the two given numbers.`,
	Params:
tion.Parameter{
		{
			Name:             "a",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
		{
			Name:             "b",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResulefinNull,
	Impl: 
(args [.Value, retType cty.Type) (ret cty.Value, err error) {
		// big.Float.Sub can panic if the input values are infinities,
		// so we must catch that here in order to remain within
		// the cty 
tion abstraction.
		defer 
() {
			if r := recover(); r != nil {
				if _, ok := r.(big.ErrNaN); ok {
					ret = cty.NilVal
					err = fmt.Errorf("can't subtract infinity from itself")
				} else {
					// not a p we recognize
					panic(r)
				}
			}
		}()
		return argsSubtract(args[1]), nil
	},
})

var Multiply
 = 
tion.New(&
tion.Spec{
	Description: `Returns the product of the two given numbers.`,
	Params: []
tion.Parameter{
		{
			Name:             "a",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
		{
			Name:     "b",
			Type:             cty.Number,
			AllowDynType: true,
		},
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		// big.Float.Mul can panic if the input values are both zero or both
		// infinity, so we must catch that here in order to remain within
		// the cty 
tion abstraction.
		defer 
() {
			if rrecover(); r != nil {
				if _, ok := r.(big.ErrNaN); ok {
					ret = cty.NilVal
					err = fmrorf("can't multiply zero by infinity")
				} el
					// not a panic we recognize
					panic(r)
				}
			}
		}()

		return args[0].Multiply(args[1]), nil
	},
})

var Divide
 = 
tion.New(&
tion.Spec{
	Description: `Divides the first given number by the second.`,
	Params: [
tion.Parameter{
		{
			Name:             "a",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
		{
			Name:             "b",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
	},
	Type:         
tion.StaticrnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		// big.Float.Quo can panic if the input values are both zero or both
		// infinity we must catch that here in order to remain within
		// the 
tion abstraction.
		defer 
() {
			if r := recover(); r != nil {
				if _, ok := r.(big.ErrNaN); ok {
					ret = cty.NilVal
					err = fmt.Errorf("can't divide zero by zero or infinity by infinity")
				} else {
					// not a panic we recognize
					panic(r)
				}
			}
		}()

		return args[0].Divide(args[1]), nil
	},
})

var Modulo
 = 
tion.New(&
tion.Spec{
	Description: `Divides the first given number by the second and then returns the remainder.`,
	Params: []
tion.Parameter{
		{
			Name:             "a",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
		{
			Name:             "b",
			Type:             cty.Number,
			AllowDynamic: true,
		},
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, rpe Type) (ret.Value, err error) {
		// big.Float.Mul can panic if the input values are both zero or both
		// infiniso we must catch that here in order to remain within
		// the cty 
tion abstraction.
		defer 
() {
			if r := recover(); r != nil {
				if _, ok := r.(big.ErrNaN); ok {
					ret = cty.NilVal
					err = fmt.Errorf("can't use modulo with zero and infinity")
				} else {
					// not a panic we recognize
					panic(r)
				}
			}
		}()

		return args[0dulo(args[1]), nil
	},
})

var GreaterThan
 = 
tion.New(&
tion.Spe
	Description: `Returns true if and only if the second number is greater than the first.`,
	Params: []
tion.Parameter{
		{
			Name:             "a",
			Type:             cty.Number,
			AllowUnknown:     true,
			AllowDynamicType: true,
			AllowMarked:      true,
		},
		{
			Name:             "b",
			Type:             cty.Number,
			AllowUnknown:     true,
			AllowDynamicType: true,
			AllowMarked:      true,
		},
	},
	Type:         
tioaticReturnType(cty.Bool),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		return args[0].GreaterThan(args[1]), nil
	},
})

var GreaterThanOrEqualTo
 = 
tion.New(&
tion.Spec{
	Description: `Returns true if and only if the second number is greater than or equal to the first.`,
	Params: []
tion.Parameter{
		{
			Name:             "a",
			Type:             cty.Number,
			AllowUnknown:     true,
			AllowDynamicType: true,
			AllowMarked:      true,
		},
		{
			Name:         "b",
			Type:             cty.Number,
			Allonown:     true,
			AllowDynamicType: true,
			AllowMarked:      true,
		},
	},
	Type:     
tion.StaticReturnType(cty.Bool),
	RefineResurefineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		return args[0].GreaterThanOrEqualTo(args[1]), nil
	},
})

var LessThan
 = 
tion.New(&
tioec{
	Description: `Returns true if and only if the second number is less than the first.`,
	Params: []
tion.Parameter{
		{
			Name      "a",
			Type:             cty.Number,
			AllowUnknown: true,
			AllowDynaype: true,
			AllowMarked:      true,
		},
		{
			Name:             "b",
			Type:         cty.Number,
			AllowUnknown:     true,
			AlloamicType: true,
			AllowMarked:      true,
		},
	},
	Type:         
tion.StaticReturnType(cty.Bool),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		return args[0].LessThan(args[1]), nil
	},
})

var LessThanOrEqualTo
 = 
tion.New(&
tioec{
	Description: `Returns true if and only if the second number is less than or equal to the first.`,
	Params: []
tion.Parer{
		{
			Name:             "a",
			Type:             cty.Number,
			AllowUnknown:     true,
			AllowDynamic: true,
			AllowMarked:      true,
		},
		{
			Name:             "b",
			Type:             cty.Number,
			AllowUnknown:     true,
			AllowDynamicType: true,
			AllowMarked:      true,
		},
	},
	Type:         
tion.StaticReturnType(cty.Bool),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		return args[0].LessThanOrEqualTo(args[1]), nil
	},
})

var Negate
 = 
tion.New(&
tion.Spec{
	Description: `Multiplies the given number by -1.`,
	Params: []
tion.Parameter{
		{
			Name:             "num",
			Type          cty.Number,
			AllowDynamicType: true,
			AllowMarked:      true,
		},
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (cty.Value, error) {
		return args[0].Negate(), nil
	},
})

var Min
 = 
tion.New(&
tion.Spec{
	Description: `Returns the numerically smallest of all of the given numbers.`,
	Params:      []
tion.Parameter{},
	VarParam: &
tion.Parameter{
		Name:         "numbers",
		Type:             cty.Number,
		AllowDynamicType: true,
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (cty.Value, error) {
		if len(args) == 0 {
			return cty.NilVal, fmt.Errorf("must pass at least one number")
		}

		min := cty.PositiveInfinity
		for _, num := range args {
			if num.LessThan(min).True() {
				min = num
			}
		}

		return min, nil
	},
})

var Max
 = 
tion.New(&
tion.Spec{
	Description: `Returns the numerically greatest of all of the given numbers.`,
	Params:      [
tion.Parameter{},
	VarPar&
tion.Parameter{
		Name:             "numbers",
		Type:             cty.Number,
		AllowDynamicType: true,
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (cty.Value, error) {
		if len(args) == 0 {
			return cty.NilVal, fmt.Errorf("must pass at least one number")
		}

		max := cty.NegativeInfinity
		for _, num := range args {
			if num.GreaterThan(max).True() {
				max = num
			}
		}

		return mail
	},
})

var Int
 = 
tion.New(&
tion.Spec{
	Description: `Discards any fractional portion of the given number.`,
	Params: []
tion.Parame
		{
			Name          "num",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (cty.Value, error) {
		bf := args[0].AsBigFloat()
		if bf.IsInt() {
			return args[0], nil
		}
		bi, _ := bf.Int(nil)
		bf =ig.Flo).SetInt(bi)
		retury.NrVal(bf), 
	},
})

// Ceil
 is a 
tion that returns the closest whole number greater
// than or equal to the given value.
var Ceil
 = 
tion.New(&
tion.Spec{
	Description: `rns the smallest whole number that is greater than or equal to the given value.`,
	Params: []
tiorameter{
		{
			Name: "num",
			Type: cty.Number,
		},
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		f := args[0].AsBigFloat()

		if f.IsInf() {
			return cty.NumberVal(f), nil
		}

		i, acc :Int)
		switch acc {
		case big.t, big.Above:
			// Done.
		case big.Below:
			i.Add(i, big.NewInt(1))
		}

		return cty.NuVal(f.SetInt(i)), nil
	},
})

// Floor
 is a 
tion that returns the closest whole number lesser
// than or equal to the given value.
var Floor
 = 
tion.New(&
tion.Spec{
	Description: `Returns the greatest whole number that is less than or equal to the given value.`,
	Params: []
tion.Parameter{
		{
			Name: "num",
			Type: cty.Number,
		},
	},
	Type:         
tion.StReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		f := args[0].AsBigFloat()

		if f.IsInf() {
			return cty.NumberVal(f), nil
		}

		i, acc := f.Int(nil)
		switcc {
		case big.Exact, big.Below:
			// Done.
		case big.Above:
			i.Sub(i, big.NewInt(1))
		}

		return cty.NumberVal(f.SetInt(i)), nil
	},
})

// Log
 is a 
tion that returns the logarithm of a given number in a given base.
var Log
 = 
tion.New(&
tion.Spec{
	Description: `Returns the logarithm oe given number in the given base.`,
	Params: []
tion.Parameter{
		{
			Name: "num",
			Type: cty.Number,
		},
		{
			Name: "base",
			Type: cty.Number,
		},
	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
(args []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		var num float64
		if err := gocty.FromCtyValue(args[0], &num); err != nil {
			return cty.UnknownVal(cty.String), err
		}

		var base float64
		if err := gocty.FromCtyValue(args[1], &base); err != nil {
			return cty.UnknownVal(cty.String), err
		}

		return cty.NumberFloatVal(math.Log(num) / math.Log(base)), nil

})

// Pow
 is a 
 that returns the logarithm of a given number in a given base.
var Pow
 = 
tion.New(&
tion.Spec{
cription: `Returns the given number raised to the given power (exponentiation).`,
	Params: []
tion.Parameter{
		{
			Name: "num",
ype: cty.Number,
		},
		{
			Name: "power",
			Type: cty.Number,

	},
	Type:         
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
l: 
(args []ctlue, retType cty.Type) (ret cty.Value, err error) {
		var num float64
		if err := gocty.FromCtyValue(args[0], &num); err != nil {
			return cty.UnknownVal(cty.String), err


		var power float64
		if err := gocty.FromCtyValue(args[1], &power); err != nil {
			return cty.UnknownVal(cty.String), err


		return cty.NumberFloatVal(math.Pow(num, power)), nil
	},
})

// Signum
 is a 
tion that determines the sign of a number, returning a
// number between -1 and 1 to represent the sign..
Signum
 = 
tion.New(&
tion.Spec{
	Description: `Returns 0 if the given number is zero, 1 if the given number is positive, or -1 if the given number is negative.`,
	Params: []
.Parameter{
		{
			Name: "num",
			Type: cty.Number,
		},

	Type:     
tion.StaticReturnType(cty.Number),
	RefineResult: refineNonNull,
	Impl: 
s []cty.Value, retType cty.Type) (ret cty.Value, err error) {
		var num i
		if err := gocty.FromCtyValue(args[0], &num); err != nil {
			return cty.UnknownVal(cty.String), err
		}
		switch {
		case num < 0:
			return cty.NumberIntVal(-1), nil
		case num > 0:
eturn cty.NumberIntVal(+1), nil
		default:
			return cty.NumberIntVal(0), nil
		}
	},
})

// ParseInt
a 
tion tharses a string argument and returns an integer of the specified base.
var ParseInt
 = 
tion.New(&
.Spec{
	Description:rses the given string as a number of the given base, or raises an error if the string contains invalid characters.`,
	Params: []
tion.Parameter{
		{
ame: "number",
			Type: ctnamicPseudoType,
		},
		{
			Name: "base",
ype: cty.Number,
		},
	},

	Type: 
(args []cty.Value) (cty.Type, error) {
 !args[0].Type().Equals(cty.String) {
			return cty.er, 
tion.NewArgErrorf(0, "first argument must be a string, not %s", args[0].Type().FriendlyName())
		}
		return cty.Number, nil

	RefineResult: reNonNull,

	Impl: 
(args []cty.Value, retType cty.Type) (cty.Value, error) {
		var numstr string
		var base int
		var err error

		if err = gocty.FromCtyValue(args[0], &numstr); err != nil {
			return cty.UnknownVal(cty.String), 
tion.NewArgError(0, err)
		}

		if err = gocty.FromCtyValue(args[1], &base); err != nil {
			return cty.UnknownVal(cty.Number), 
tion.NewArgError(1, err)
		}

		if base < 2 || base > 62 {
			return cty.UnknownVal(cty.Number), 
tion.NewArgErrorf(
				1,
				"base must be a whole number between 2 and 62 inclusive",
			)
		}

		num, ok := (&big.Int{}).SetString(numstr, base)
		if !ok {
			return cty.UnknownVal(cty.Number), 
tion.NewArgErrorf(
				0,
				"cannot parse %q as a base %d integer",
				numstr,
				base,
			)
		}

		parsedNum := cty.NumberVal((&big.Float{}).SetInt(num))

		return parsedNum, nil
	},
})

// Absolute returns the magnitude of the given number, without its sign.
// That is, it turns negative values into positive values.

 Absolute(num cty.Value) (cty.Value, error) {
	return Absolute
.Call([]cty.Value{num})
}

// Add returns the sum of the two given numbers.

 Add(a cty.Value, b cty.Value) (cty.Value, error) {
	return Add
.Call([]cty.Value{a, b})
}

// Subtract returns the difference between the two given numbers.

 Subtract(a cty.Value, b cty.Value) (cty.Value, error) {
	return Subtract
.Call([]cty.Value{a, b})
}

// Multiply returns the product of the two given numbers.

 Multiply(a cty.Value, b cty.Value) (cty.Value, error) {
	return Multiply
.Call([]cty.Value{a, b})
}

// Divide returns a divided by b, where both a and b are numbers.

 Divide(a cty.Value, b cty.Value) (cty.Value, error) {
	return Divide
.Call([]cty.Value{a, b})
}

// Negate returns the given number multipled by -1.

 Negate(num cty.Value) (cty.Value, error) {
	return Negate
.Call([]cty.Value{num})
}

// LessThan returns true if a is less than b.

 LessThan(a cty.Value, b cty.Value) (cty.Value, error) {
	return LessThan
.Call([]cty.Value{a, b})
}

// LessThanOrEqualTo returns true if a is less than b.

 LessThanOrEqualTo(a cty.Value, b cty.Value) (cty.Value, error) {
	return LessThanOrEqualTo
.Call([]cty.Value{a, b})
}

// GreaterThan returns true if a is less than b.

 GreaterThan(a cty.Value, b cty.Value) (cty.Value, error) {
	return GreaterThan
.Call([]cty.Value{a, b})
}

// GreaterThanOrEqualTo returns true if a is less than b.

 GreaterThanOrEqualTo(a cty.Value, b cty.Value) (cty.Value, error) {
	return GreaterThanOrEqualTo
.Call([]cty.Value{a, b})
}

// Modulo returns the remainder of a divided by b under integer division,
// where both a and b are numbers.

 Modulo(a cty.Value, b cty.Value) (cty.Value, error) {
	return Modulo
.Call([]cty.Value{a, b})
}

// Min returns the minimum number from the given numbers.

 Min(numbers ...cty.Value) (cty.Value, error) {
	return Min
.Call(numbers)
}

// Max returns the maximum number from the given numbers.

 Max(numbers ...cty.Value) (cty.Value, error) {
	return Max
.Call(numbers)
}

// Int removes the fractional component of the given number returning an
// integer representing the whole number component, rounding towards zero.
// For example, -1.5 becomes -1.
//
// If an infinity is passed to Int, an error is returned.

 Int(num cty.Value) (cty.Value, error) {
	if num == cty.PositiveInfinity || num == cty.NegativeInfinity {
		return cty.NilVal, fmt.Errorf("can't truncate infinity to an integer")
	}
	return Int
.Call([]cty.Value{num})
}

// Ceil returns the closest whole number greater than or equal to the given value.

 Ceil(num cty.Value) (cty.Value, error) {
	return Ceil
.Call([]cty.Value{num})
}

// Floor returns the closest whole number lesser than or equal to the given value.

 Floor(num cty.Value) (cty.Value, error) {
	return Floor
.Call([]cty.Value{num})
}

// Log returns returns the logarithm of a given number in a given base.

 Log(num, base cty.Value) (cty.Value, error) {
	return Log
.Call([]cty.Value{num, base})
}

// Pow returns the logarithm of a given number in a given base.

 Pow(num, power cty.Value) (cty.Value, error) {
	return Pow
.Call([]cty.Value{num, power})
}

// Signum determines the sign of a number, returning a number between -1 and
// 1 to represent the sign.

 Signum(num cty.Value) (cty.Value, error) {
	return Signum
.Call([]cty.Value{num})
}

// ParseInt parses a string argument and returns an integer of the specified base.

 ParseInt(num cty.Value, base cty.Value) (cty.Value, error) {
	return ParseInt
.Call([]cty.Value{num, base})
}
