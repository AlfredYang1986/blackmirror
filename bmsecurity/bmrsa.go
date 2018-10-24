package bmsecurity

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 加密
func RsaEncrypt(publicKey, origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(privateKey, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// ph加密
func PhRsaEncrypt(publicKey string, origData []byte) ([]byte, error) {
	pubKeyBytes,_ := base64.StdEncoding.DecodeString(publicKey)
	pubInterface, err := x509.ParsePKIXPublicKey(pubKeyBytes)

	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// ph解密
func PhRsaDecrypt(privateKey string, ciphertext []byte) ([]byte, error) {
	priKeyBytes,_ := base64.StdEncoding.DecodeString(privateKey)
	priv, err := x509.ParsePKCS8PrivateKey(priKeyBytes)

	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), ciphertext)
}