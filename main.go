package main

import (
	"fmt"
	"jwk/utils"
	"log"
)

func TestECDSA() {
	//生成随机钥字符串长度40字节，用于生产公私钥证书

	randKey := utils.GetRandomString(40)
	//生成随机签名字符串40字节，用于加密数据
	randSign := utils.GetRandomString(40)

	//使用随机钥字符串生成公私钥文件
	e := utils.GenerateKey(randKey)
	if e != nil {
		log.Fatal(e)
	}

	//签名附加信息
	srcInfo := "Offer Come On please"
	fmt.Println(srcInfo)

	//ECC签名加密
	signByEcc, e := utils.CryptSignByEcc(srcInfo, utils.PRIVATEFILE, randSign)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("ECDSA私钥加密签名为：", signByEcc)

	//ECC签名算法校验
	verifyCryptEcc, e := utils.VerifyCryptEcc(srcInfo, signByEcc, "../Keyset/publicKey.pem")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("ECDSA公钥解密后验签校验结果：", verifyCryptEcc)

}

func main() {
	TestECDSA()
}
