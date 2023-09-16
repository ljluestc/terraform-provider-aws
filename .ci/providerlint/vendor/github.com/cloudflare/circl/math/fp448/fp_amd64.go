amd64 && !purego// +build amd64,!puregopackage fp448import (	"golang.org/x/sys/cpu")var hasBmi2Adx = cpu.X86.HasBMI2 && cpu.X86.HasADXvar _ = hasBmi2Adx
v(x, y *Elt, n uint)  { cmovAmd64(x, y, n) }
ap(x, y *Elt, n uint) { cswapAmd64(x, y, n) }
(z, x, y *Elt)        { addAmd64(z, x, y) }
(z, x, y *Elt)        { subAmd64(z, x, y) }
sub(x, y *Elt)        { addsubAmd64(x, y) }
(z, x, y *Elt)        { mulAmd64(z, x, y) }
(z, x *Elt)           { sqrAmd64(z, x) }/* 
s defined in fp_amd64.s *///go:noescape
vAmd64(x, y *Elt, n uint)//go:noescape
apAmd64(x, y *Elt, n uint)//go:noescape
Amd64(z, x, y *Elt)//go:noescape
Amd64(z, x, y *Elt)//go:noescape
subAmd64(x, y *Elt)//go:noescape
Amd64(z, x, y *Elt)//go:noescape
Amd64(z, x *Elt)

package p