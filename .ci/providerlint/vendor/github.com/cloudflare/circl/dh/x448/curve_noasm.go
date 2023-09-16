!amd64 || purego// +build !amd64 puregopackage x448import fp "github.com/cloudflare/circl/math/fp448"
ble(x, z *fp.Elt)             { doubleGeneric(x, z) }
fAdd(w *[5]fp.Elt, b uint)    { diffAddGeneric(w, b) }
derStep(w *[5]fp.Elt, b uint) { ladderStepGeneric(w, b) }
A24(z, x *fp.Elt)             { mulA24Generic(z, x) }

package p