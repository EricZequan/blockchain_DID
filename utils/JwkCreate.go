package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"encoding/json"
	"math/big"
)

// ECDSAPublicKeyJWK 是用于表示ECDSA公钥的JWK格式的结构
type ECDSAPublicKeyJWK struct {
	Kty string `json:"kty"`
	Crv string `json:"crv"`
	X   string `json:"x"`
	Y   string `json:"y"`
}

func CreateJWK(public *ecdsa.PublicKey) []byte {
	jwk := ECDSAPublicKeyJWK{
		Kty: "EC",
		Crv: "P-256",
		X:   encodeBigInt(public.X),
		Y:   encodeBigInt(public.Y),
	}
	jwkJSON, err := json.Marshal(jwk)
	if err != nil {
		panic(err)
	}
	return jwkJSON
}

func ParseJwkToPublicKey(jwk []byte) *ecdsa.PublicKey {
	var jwk_pub ECDSAPublicKeyJWK
	err := json.Unmarshal([]byte(jwk), &jwk_pub)
	if err != nil {
		panic(err)
	}
	// 将base64URL编码的字符串转换为*big.Int
	x, err := base64ToBigInt(jwk_pub.X)
	if err != nil {
		panic(err)
	}

	y, err := base64ToBigInt(jwk_pub.Y)
	if err != nil {
		panic(err)
	}

	// 使用X和Y坐标构造ecdsa.PublicKey
	pub := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
	return pub
}

// encodeBigInt 将big.Int转换为其base64URL编码的字符串表示形式
func encodeBigInt(number *big.Int) string {
	// 请注意，这里使用了URL安全的base64编码，但没有填充，因为JWK规范要求如此
	return base64.RawURLEncoding.EncodeToString(number.Bytes())
}

// base64ToBigInt 将base64URL编码的字符串转换为*big.Int
func base64ToBigInt(encoded string) (*big.Int, error) {
	decoded, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	number := new(big.Int).SetBytes(decoded)
	return number, nil
}
