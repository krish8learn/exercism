package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	// Generate a random number between 2 and p-1 (inclusive)
	// rand.Int returns a value in [0, max), so we need max = p - 2
	max := new(big.Int).Sub(p, big.NewInt(2))
	random, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	// Add 2 to shift range to [2, p-1]
	return new(big.Int).Add(random, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	// Calculate g^private mod p
	gBig := big.NewInt(g)
	return new(big.Int).Exp(gBig, private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	// Generate a private key and compute its corresponding public key
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	// Calculate public2^private1 mod p
	return new(big.Int).Exp(public2, private1, p)
}
