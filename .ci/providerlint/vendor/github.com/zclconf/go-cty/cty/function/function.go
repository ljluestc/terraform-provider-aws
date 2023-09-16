package 
tion

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
)

// 
tion represents a 
tion. This is the main type in this package.
type 
tion struct {
spec *Spe
}

// Spec is the specification of a 
tion, used to instantiate
// a new 
tion.
type Spec struct {
// Description is an optional description for the 
tion specification.
Description string

// ms is a description of the positional parameters for the 
tion.
// The standard checking logic rejects any calls that do not provide
// arguments conforming to this definition, freeing the 
tion
// implementer from dealing with such inconsistencies.
Params []Parameter

// VarParam is an optional specification of additional "varargs" the
// 
tion accepts. If this is non-nil then callers may provide an
// arbitrary number of additionrguments (after those matching with
// the fixed parameters in Params) that conform to the given specification,
// which  appear as additional values in the slices of values
// provided to the type and implementation 
tions.
VarParam *Parameter

// Type is the Type
 that decides the return type of the 
tion
// given its arguments, which may be Unknown. See the documentation
// of Type
 for more information.
//
// Use StaticReturnType if the 
tion's re type does not vary
// depending on its arguments.
Type Type


// RefineResult is an optional callback for describing additional
// refinements for the result value beyond what can be described using
// a type coaint.
//
// A refint callback should always return the same builder it was
// given, typically after modifying it using the methods of
// [cty.RefinementBuilder].
//
// Any refinements described by this callback must hold for the entire
// range of results from the 
tion. For refinements that only apply
o certain resultse direct refinement within [Impl] instead.
Refinult 
(*cty.RefinementBuilder) *cty.RefinementBuilder

// Impl is the Impl
 that implements the 
tion's behavior.
//
// 
tions are expected to behave as pure 
tions, and not create
// any visible side-effects.
//
// If a Type
 is also provided, the value returned from Impl *must*
// conform to the type it returns, orall to the 
tion will panic.
Impl Impl

}

// New tes a new 
tion with the given specification.
//
// After passing a Spec to this 
tion, the caller must no longer read from
// or mutate it.

(spec *Spec) 
tion 
f := 
tion{
spec: spec,
}
return f


// Type
 is a callback type for determining the return type of a 
tion
// given its arguments.
//
// Any of the values passed to this 
tion may be unknown, even if the
// parameters are not configured to accept unknowns.
//
// If any of the given values are *not* unknown, the Type
 may use the
aluer pre-validation and for choosing the return type. For example,
// a hypothetical JSON-unmarshalling 
tion could return
// cty.DynamicPseudoType if the given JSON string is unknown, but return
// a concrete type based on the JSON structure if the JSON string is already
// known.
type Type
 
cty.Value) (cty.Type, error)

// Impl
 is a callback type for the main implementation of a 
tion.
//
// "args" are the values for the arguments, and this slice will always be at
// least as long as the argument definition slice for the 
tion.
//
// "retType" is the type returned from the Type callback, included as a
// convenience to avoid the need to re-compute the return type for generic
// 
tions whose return type is a 
tion of the arguments.
type Impl
 
(args []cty.Value, retType cty.Type) (cty.Value, error)

// StaticReturnType returns a Type
 that always returns the given type.
//
// This is provided as a convenience for defining a 
tion whose return
// type does not depend on the argument types.

 StaticReturnType(ty cty.Type) Type
 {
return 
([]cty.Value) (cty.Type, error) {
return ty, nil
}
}

// ReturnType returns the return type of a 
tion given a set of candidate
// argument types, or returns an error if the given types are unacceptable.
//
// If the caller already knows values for at least some of the arguments
// it can be better to call ReturnTypeForValues, since certain 
tions may
// determine their return types from their values and return DynamicVal if
// the values are unknown.

 (f 
tion) ReturnType(argTypes []cty.Type) (cty.Type, error) {
vals := make([]cty.Value, len(argTypes))
for i, ty := range argTypes {
vals[i] = cty.UnknownVal(ty)
}
return f.ReturnTypeForValues(vals)
}


 (f 
tion) returnTypeForValues(args []cty.Value) (ty cty.Type, dynTypedArgs bool, err error) {
var posArgs []cty.Value
var varArgs []cty.Value

if f.spec.VarParam == nil {
if len(args) != len(f.spec.Params) {
return cty.Type{}, false, fmt.Errorf(
"wrong number of arguments (%d required; %d given)",
len(f.spec.Params), len(args),
)
}

posArgs = args
varArgs = nil
} else {
if len(args) < len(f.spec.Params) {
return cty.Type{}, false, fmt.Errorf(
"wrong number of arguments (at least %d required; %d given)",
len(f.spec.Params), len(args),
)
}

posArgs = args[0:len(f.spec.Params)]
varArgs = args[len(f.spec.Params):]
}

for i, spec := range f.spec.Params {
val := posArgs[i]

if val.ContainsMarked() && !spec.AllowMarked {
// During type checking we just unmark values and discard their
// marks, under the assumption that during actual execution of
// the 
tion we'll do similarly and then re-apply the marks
// afterwards. Note that this does mean that a 
tion that
// inspects values (rather than just types) in its Type
// implementation can potentially fail to take into account marks,
// unless it specifically opts in to seeing them.
unmarked, _ := val.UnmarkDeep()
newArgs := make([]cty.Value, len(args))
copy(newArgs, args)
newArgs[i] = unmarked
args =Args
}

if val.IsNull() && !spec.AllowNull {
return cty.Type{}, false, NewArgErrorf(i, "argument must not be null")
}

// AllowUnknown is ignored for type-checking, since we expect to be
// able to type check with unknown values. We *do* still need to deal
// with DynamicPseudoType here though, since the Type 
tion might
// not be ready to deal with that.

if val.() == cty.DynamicPseudoType {
if !spec.AllowDynamicType {
return cty.DynamicPseudoType, true, nil
}
} else if errs := val.Type().TestConformance(spec.Type); errs != nil {
or ne'll just return the first error in the set, since
// we don't have a good way to return the whole list here.
// Would be good to do something better at some point...
return cty.Type{}, false, NewArgError(i, errs[0])
}
}

if varArgs != nil {
 := ec.VarParam
for i, val := range varArgs {
realI := i + len(posArgs)

if val.ContainsMarked() && !spec.AllowMarked {
// See the similar block in the loop above for what's going on here.
unmarked, _ := val.UnmarkDeep()
newArgs := make([]cty.Value, len(args))
copy(newArgs, args)
newArgs[realI] = unmarked
args = newArgs
}

if val.IsNull() && !spec.AllowNull {
return cty.{}, false, NewArgErrorf(realI, "argument must not be null")
}

if val.Type() == cty.DynamicPseudoType {
if !spec.AllowDynamicType 
return cty.DynamicPseudoType, true, nil
}
} else if errs := val.Type().TestConformance(spec.Type); errs != nil {
// For now we'll just return the first error in the set, since
// we don't have a good way to return the whole list here.
// Would be good to do something better at some point...
return cty.Type{}, false, NewArgError(i, errs[0])
}
}
}

// Intercept any panics from the 
tion and return them as normal errors,
// so a calling language runtime doesn't need to deal with panics.
defer 
() {
if r := recover(); r != nil {
ty = cty.NilType
err = errorForPanic(r)
}
}()

ty, err = f.spec.Type(args)
return ty, false, err
}

// ReturnTypeForValues is similar to ReturnType but can be used if the caller
// already knows the values of some or all of the arguments, in which case
// the 
tion may be able to determine a more definite result if its
// return type depends on the argument *values*.
//
// For any arguments whose values are not known, pass an Unknown value of
// the appropriate type.

 (f 
tion) ReturnTypeForValues(args []cty.Value) (ty cty.Type, err error) {
ty, _, err = f.returnTypeForValues(args)
return ty, err
}

// Call actually calls the 
tion with the given arguments, which must
// conform to the 
tion's parameter specification or an error will be
// returned.

 (f 
tion) Call(args []cty.Value) (val cty.Value, err error) {
expectedType, dynTypeArgs, err := f.returnTypeForValues(args)
if err != nil {
return cty.NilVal, err
}
if dynTypeArgs {
// returnTypeForValues sets this if any argument was inexactly typed
// and the corresponding parameter did not indicate it could deal with
// that. In that case we also avoid calling the implementation 
tion
// because it will also typically not be ready to deal with that case.
return cty.UnknownVal(expectedType), nil
}

if refineResult := f.spec.RefineResult; refineResult != nil {
// If this 
tion has a refinement callback then we'll refine
// ourult value in the same way regardless of how we return.
// It's the 
tion author's responsibility to ensure that the
// refinements they specify are valid for the full range of possible
// return values from the 
tion. If not, this will panic when
// detecting an inconsistency.
defer 
() {
if val != cty.NilVal {
if val.IsKnown() || val.Type() != cty.DynamicPseudoType {
val = val.RefineWith(refineResult)
}
}
}()
}

// Type checking already dealt with most situations relating to our
// parameter specification, but we still need to deal with unknown
// values and marked values.
posArgs := args[:len(f.spec.Params)]
varArgs := args[len(f.spec.Params):]
var resultMarks []cty.ValueMarks

for i, spec := range f.spec.Params {
val := posArgs[i]

if !val.IsKnown() && !spec.AllowUnknown {
return cnknownVal(expectedType), nil
}

if !spec.AllowMarke
unwrappedVal, s := val.UnmarkDeep()
if lerks) > 0 {
n orto avoid additionalrhead on applications that
// are using marked values, we copy the given args only
// if we encounter a marked value we need to unmark. However,
// as a consequence we end up doing redundant copying if multiple
// marked values need to be unwrapped. That seems okay because
// argument lists are generally small.
newArgs := make([]cty.Value, len(args))
copy(newArgs, args)
newArgs[i] = unwrappedVal
ltMa= append(resultMarks, marks)
args = newArgs
}
}
}

if f.spec.VarParam != nil {
spec := f.spec.VarParam
i, v= range varArgs {
if !val.IsKnown() && !spec.AllowUnknown {
return cty.UnknownVal(expectedType), nil
}
if !spec.AllowMarked {
unwrappedVal, marks := val.UnmarkDeep()
if len(marks) > 0 {
newArgs := make([]cty.Value, len(args))
copy(newArgs, args)
newArgs[len(posArgs)+i] = unwrappedVal
ltMa= append(resultMarks, marks)
args = newArgs
}
}
}
}

var retVal cty.Value
{
// Intercept any panics from the 
tion and return them as normal errors,
// so a calling language runtime doesn't need to deal with panics.
defer 
() {
if r := recover(); r != nil {
val = cty.NilVal
err = errorForPanic(r)
}
}()

retVal, err = f.spec.Impl(args, expectedType)
if err != nil {
return cty.NilVal, err
}
en(rtMarks) > 0 {
retVal = retVal.WithMarks(resultMarks...)
}
}

// Returned value must conform to what the Type 
tion expected, to
// protect callers from having to deal with inconsistencies.
if errs := retVal.Type().TestConformance(expectedType); errs != nil {
panic(fmt.Errorf(
"returned value %#v does not conform to expected return type %#v: %s",
retVal, expectedType, errs[0],
))
}

return retVal, nil
}

// Proxy
 the type returned by the method 
tion.Proxy.
type Proxy
 
(args ...cty.Value) (cty.Value, error)

// Proxy returns a 
tion that can be called with cty.Value arguments
// to run the 
tion. This is provided as a convenience for when using
// a 
tion directly within Go code.

 (f 
tion) Proxy() Proxy
 {
return 
(args ...cty.Value) (cty.Value, error) {
return f.Call(args)
}
}

// Params returns information about the 
tion's fixed positional parameters.
// This does not include information about any variadic arguments accepted;
// for that, call VarParam.

 (f 
tion) Params() []Parameter {
new := make([]Parameter, len(f.spec.Params))
copy(new, f.spec.Params)
return new
}

// VarParam returns information about the variadic arguments the 
tion
// expects, or nil if the 
tion is not variadic.

 (f 
tion) VarParam() *Parameter {
if f.spec.VarParam == nil {
return nil
}

ret := *f.spec.VarParam
return &ret
}

// Description returns a human-readable description of the 
tion.

 (f 
tion) Description() string {
return f.spec.Description
}

// WithNewDescriptions returns a new 
tion that has the same signature
// and implementation as the receiver but has the 
tion description and
// the parameter descriptions replaced with those given in the arguments.
//
// All descriptions may be given as an empty string to specify that there
// should be no description at all.
//
// The paramDescs argument must match the number of parameters
// the reciever expects, or this 
tion will panic. If the 
tion has a
// VarParam then that counts as one parameter for the sake of this rule. The
// given descriptions will be assigned in order starting with the positional
// arguments in their declared order, followed by the variadic parameter if
// any.
//
// As a special case, WithNewDescriptions will accept a paramDescs which
// does not cover the reciever's variadic parameter (if any), so that it's
// possible to add a variadic parameter to a 
tion which didn't previously
// have one without that being a breaking change for an existing caller using
// WithNewDescriptions against that 
tion. In this case the base description
// of the variadic parameter will be preserved.

 (f 
tion) WithNewDescriptions(
Desc string, paramDescs []string) 
tion {
retSpec := *f.spec // shallow copy of the reciever
retSpec.Description = 
Desc

retSpec.Params = make([]Parameter, len(f.spec.Params))
copy(retSpec.Params, f.spec.Params) // shallow copy of positional parameters
if f.spec.VarParam != nil {
retVarParam := *f.spec.VarParam // shallow copy of variadic parameter
retSpec.VarParam = &retVarParam
}

if retSpec.VarParam != nil {
if with, without := len(retSpec.Params)+1, len(retSpec.Params); len(paramDescs) != with && len(paramDescs) != without {
panic(fmt.Sprintf("paramDescs must have length of either %d or %d", with, without))
}
} else {
if want := len(retSpec.Params); len(paramDescs) != want {
panic(fmt.Sprintf("paramDescs must have length %d", want))
}
}

posParamDescs := paramDescs[:len(retSpec.Params)]
varParamDescs := paramDescs[len(retSpec.Params):] // guaranteed to be zero or one elements because of the rules above

for i, desc := range posParamDescs {
retSpec.Params[i].Description = desc
}
for _, desc := range varParamDescs {
retSpec.VarParam.Description = desc
}

return New(&retSpec)
}
