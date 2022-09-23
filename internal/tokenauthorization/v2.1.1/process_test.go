package tokenauthorization_test

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	secp "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)


func TestCreateTokenAuthorization(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		privateKey, err := secp.GeneratePrivateKey()

		if err != nil {
			t.Fatal("GeneratePrivateKey")
		}

		privateKeyBytes := privateKey.Serialize()
		_ = hex.EncodeToString(privateKeyBytes)

		publicKey := privateKey.PubKey()
		publicKeyBytes := publicKey.SerializeCompressed()
		publicKeyHex := hex.EncodeToString(publicKeyBytes)
		t.Logf("publicKeyBytes: %v", publicKeyBytes)

		msg := "Hello World"
		h := sha256.New()
		h.Write([]byte(msg))
		
		hash := h.Sum(nil)
		hashHex := hex.EncodeToString(hash)
		t.Logf("hash: %v", hash)

		signature := ecdsa.Sign(privateKey, hash)
		signatureBytes := signature.Serialize()
		signatureHex := hex.EncodeToString(signatureBytes)

		t.Logf("signatureBytes: %v", signatureBytes)

		ok := signature.Verify(hash, publicKey)

		if !ok {
			t.Fatal("ok")
		}

		// Verify
		receivedPublicKeyBytes, err := hex.DecodeString(publicKeyHex)

		if err != nil {
			t.Fatal("DecodeString publicKeyHex")
		}

		t.Logf("receivedPublicKeyBytes: %v", receivedPublicKeyBytes)

		receivedPublicKey, err := secp.ParsePubKey(receivedPublicKeyBytes)

		if err != nil {
			t.Fatal("ParsePubKey")
		}

		receivedHash, err := hex.DecodeString(hashHex)

		if err != nil {
			t.Fatal("DecodeString receivedHash")
		}

		t.Logf("receivedHash: %v", receivedHash)
		receivedSignatureBytes, err := hex.DecodeString(signatureHex)

		if err != nil {
			t.Fatal("DecodeString receivedSignatureBytes")
		}

		t.Logf("receivedSignatureBytes: %v", receivedSignatureBytes)
		receivedSignature, err := ecdsa.ParseDERSignature(receivedSignatureBytes)
		
		if err != nil {
			t.Fatal("ParseDERSignature")
		}

		if !receivedSignature.Verify(receivedHash, receivedPublicKey) {
			t.Fatal("receivedSignature ok")
		}
	})
}