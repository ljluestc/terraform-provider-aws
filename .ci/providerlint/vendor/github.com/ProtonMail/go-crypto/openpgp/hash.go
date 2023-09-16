package openpgp

import (
	"crypto"

	"github.com/ProtonMail/go-crypto/openpgp/internal/algorithm"
)

// HashIdToHash returns a crypto.Hash which corresponds to the given OpenPGP
// hash id.

hIdToHash(id byte) (h crypto.Hash, ok bool) {
	return algorithm.HashIdToHash(id)
}

// HashIdToString returns the name of the hash 
 corresponding to the
// given OpenPGP hash id.

hIdToString(id byte) (name string, ok bool) {
	return algorithm.HashIdToString(id)
}

// HashToHashId returns an OpenPGP hash id which corresponds the given Hash.

hToHashId(h crypto.Hash) (id byte, ok bool) {
	return algorithm.HashToHashId(h)
}
