package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GetPrivateKeyByPemFile(priKeyFile string) (*ecdsa.PrivateKey, error) {
	//将私钥文件中的私钥读出，得到使用pem编码的字符串
	file, err := os.Open(priKeyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := fileInfo.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	//将得到的字符串解码
	block, _ := pem.Decode(buffer)

	//使用x509将编码之后的私钥解析出来
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
func GetPublicKeyByPemFile(pubKeyFile string) (*ecdsa.PublicKey, error) {
	var err error
	//从公钥文件获取钥匙字符串
	file, err := os.Open(pubKeyFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	//将得到的字符串解码
	block, _ := pem.Decode(buffer)

	//使用x509将编码之后的公钥解析出来
	pubInner, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey := pubInner.(*ecdsa.PublicKey)

	return publicKey, nil
}
