package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)
type RSA struct {
	Credential string `json:"credential"`
}

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 4096)
	return privkey, &privkey.PublicKey
}

func ConvertPrivateKeyToFilePem(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ConvertFilePemToPrivateKey(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ConvertPublicKeyToFilePem(pubkey *rsa.PublicKey) (string) {
	pubkey_bytes, _ := x509.MarshalPKIXPublicKey(pubkey)

	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem)
}

func ConvertFilePemToPublicKey(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}
func Encryption(data string) string{
	stringPublicKey, _ := os.ReadFile("pkg/utils/publicKey.bem")
	publicKey, _ := ConvertFilePemToPublicKey(string(stringPublicKey))
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte(data),
		nil)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes)
}
func Decryption(credential string) string {
	data, err := base64.StdEncoding.DecodeString(credential)
	stringPrivateKey, _ := os.ReadFile("pkg/utils/privateKey.bem")
	privateKey, _ := ConvertFilePemToPrivateKey(string(stringPrivateKey))
	decryptedBytes, err := privateKey.Decrypt(nil, []byte(data), &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	return string(decryptedBytes)
}
func Signing(data string) {
	stringPublicKey, _ := os.ReadFile("pkg/utils/publicKey.bem")
	publicKey, _ := ConvertFilePemToPublicKey(string(stringPublicKey))
	stringPrivateKey, _ := os.ReadFile("pkg/utils/privateKey.bem")
	privateKey, _ := ConvertFilePemToPrivateKey(string(stringPrivateKey))

	hash := md5.New()
	hash.Write([]byte(data))
	msgHashSum := hash.Sum(nil)

	signature, _ := rsa.SignPSS(rand.Reader, privateKey, crypto.MD5, msgHashSum, nil)

	err := rsa.VerifyPSS(publicKey, crypto.MD5, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
	} else {
		fmt.Println("signature verified")

	}

}
