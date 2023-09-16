//go:build !amd64 || purego
// +build !amd64 purego

package fp448


v(x, y *Elt, n uint)  { cmovGeneric(x, y, n) }

ap(x, y *Elt, n uint) { cswapGeneric(x, y, n) }

(z, x, y *Elt)        { addGeneric(z, x, y) }

(z, x, y *Elt)        { subGeneric(z, x, y) }

sub(x, y *Elt)        { addsubGeneric(x, y) }

(z, x, y *Elt)        { mulGeneric(z, x, y) }

(z, x *Elt)           { sqrGeneric(z, x) }
