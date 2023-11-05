package test

import (
	"jwk/utils"
	"log"
	"testing"
)

/*
func TestJwk(t *testing.T) {
	pub, err := utils.GetPublicKeyByPemFile("../Keyset/publicKey.pem")
	if err != nil {
		log.Fatal(err)
	}
	jwkPub := utils.CreateJWK(pub)
	fmt.Printf("%s\n", jwkPub)
	res := utils.ParseJwkToPublicKey(jwkPub)
	fmt.Println(res)
}

func TestJwkSigAndVeri(t *testing.T) {
	randSign := utils.GetRandomString(40)
	//ECC签名加密
	signByEcc, err := utils.CryptSignByEcc("Offer Come On please", "../Keyset/privateKey.pem", randSign)
	if err != nil {
		log.Fatal(err)
	}
	pub, err := utils.GetPublicKeyByPemFile("../Keyset/publicKey.pem")
	if err != nil {
		log.Fatal(err)
	}
	jwkPub := utils.CreateJWK(pub)
	res_pub := utils.ParseJwkToPublicKey(jwkPub)
	fmt.Println(utils.VerifyCryptEcc("Offer Come On please", signByEcc, res_pub))
}
*/
func BenchmarkSigAndVer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 调用 Sum 函数，但为了防止编译器优化，
		// 我们通常使用其结果，例如将其赋给一个变量
		randSign := utils.GetRandomString(40)
		//ECC签名加密
		signByEcc, err := utils.CryptSignByEcc("Offer Come On please", "../Keyset/privateKey.pem", randSign)
		if err != nil {
			log.Fatal(err)
		}
		pub, err := utils.GetPublicKeyByPemFile("../Keyset/publicKey.pem")
		if err != nil {
			log.Fatal(err)
		}
		jwkPub := utils.CreateJWK(pub)
		res_pub := utils.ParseJwkToPublicKey(jwkPub)
		utils.VerifyCryptEcc("Offer Come On please", signByEcc, res_pub)
		//fmt.Println(utils.VerifyCryptEcc("Offer Come On please", signByEcc, res_pub))
	}
}
