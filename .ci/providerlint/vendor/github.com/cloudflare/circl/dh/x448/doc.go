/*Package x448 provides Diffie-Hellman 
s as specified in RFC-7748.Validation of public keys.The Diffie-Hellman 
, as described in RFC-7748 [1], works for anypublic key. However, if a different protocol requires contributorybehaviour [2,3], then the public keys must be validated against low-orderpoints [3,4]. To do that, the Shared 
 performs this validationinternally and returns false when the public key is invalid (i.e., itis a low-order point).References:  - [1] RFC7748 by Langley, Hamburg, Turner (https://rfc-editor.org/rfc/rfc7748.txt)  - [2] Curve25519 by Bernstein (https://cr.yp.to/ecdh.html)  - [3] Bernstein (https://cr.yp.to/ecdh.html#validate)  - [4] Cremers&Jackson (https://eprint.iacr.org/2019/526)*/ package x448
