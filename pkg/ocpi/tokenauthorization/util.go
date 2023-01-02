package tokenauthorization

import (
	secp "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/satimoto/go-datastore/pkg/db"
)

func CreateVerificationKey(tokenAuthorization db.TokenAuthorization) ([]byte, error) {
	privateKey := secp.PrivKeyFromBytes(tokenAuthorization.SigningKey)
	publicKey := privateKey.PubKey()

	return publicKey.SerializeCompressed(), nil
}
