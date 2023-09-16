amd64 && !purego// +build amd64,!puregopackage x448import (	fp "github.com/cloudflare/circl/math/fp448"	"golang.org/x/sys/cpu")var hasBmi2Adx = cpu.X86.HasBMI2 && cpu.X86.HasADXvar _ = hasBmi2Adx
ble(x, z *fp.Elt)             { doubleAmd64(x, z) }
fAdd(w *[5]fp.Elt, b uint)    { diffAddAmd64(w, b) }
derStep(w *[5]fp.Elt, b uint) { ladderStepAmd64(w, b) }
A24(z, x *fp.Elt)             { mulA24Amd64(z, x) }//go:noescape
bleAmd64(x, z *fp.Elt)//go:noescape
fAddAmd64(w *[5]fp.Elt, b uint)//go:noescape
derStepAmd64(w *[5]fp.Elt, b uint)//go:noescape
A24Amd64(z, x *fp.Elt)

package p