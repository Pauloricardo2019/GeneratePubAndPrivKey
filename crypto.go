package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/sha3"
)

func marshalRSAPrivate(priv *rsa.PrivateKey) string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}))
}

func marshalRSAPublic(pub *rsa.PublicKey) string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(pub),
	}))
}

func generateKey() (string, string, error) {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return "", "", err
	}

	pubKeyStr := marshalRSAPublic(key.Public().(*rsa.PublicKey))

	privKeyStr := marshalRSAPrivate(key)

	return pubKeyStr, privKeyStr, nil
}

func publicKeyFromPrivateKey(privKey string) (string, error) {
	block, _ := pem.Decode([]byte(privKey))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	fmt.Println("=========================\n", "block.Headers\n", block.Headers, "\n=========================")
	fmt.Println("=========================\n", "block.Type\n", block.Type, "\n=========================")
	fmt.Println("=========================\n", "block.Bytes\n", block.Bytes, "\n=========================")

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=========================\n", "key.PublicKey(public part)\n", key.PublicKey, "\n=========================")
	fmt.Println("=========================\n", "key.D ( private exponent )\n", key.D, "\n=========================")
	fmt.Println("=========================\n", "key.Primes ( prime factors of N, has >= 2 elements.) \n", key.Primes, "\n=========================")

	publicKeyDer, errpk := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if errpk == nil {
		return string(pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyDer})), nil
	} else {
		return "", errpk
	}
}

func encrypt(msg, publicKey string) (string, error) {
	p, decoded := pem.Decode([]byte(publicKey))
	if decoded == nil {
		log.Fatal("Failed to decode PEM")
	}
	pub, err := x509.ParsePKCS1PublicKey(p.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	encryptedBytes, err := rsa.EncryptOAEP(sha3.New512(), rand.Reader, pub, []byte(msg), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

func decrypt(data, priv string) (string, error) {
	data2, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode([]byte(priv))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	decrypted, err := rsa.DecryptOAEP(sha3.New512(), rand.Reader, key, data2, nil)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

type DateHoliday struct {
	Day   int
	Month int
	Year  int
	Date  time.Time
}

type Holiday struct {
	Date        DateHoliday
	Description string
}

func GetEasterCatholicByYear(year int) (*DateHoliday, error) {
	var a, b, c, d, e, r int

	if year < 0 {
		return nil, errors.New("year have to be greater than 0")
	}

	a = year % 19
	if year >= 1583 {
		var f, g, h, i, k, l, m int
		b = year / 100
		c = year % 100
		d = b / 4
		e = b % 4
		f = (b + 8) / 25
		g = (b - f + 1) / 3
		h = (19*a + b - d - g + 15) % 30
		i = c / 4
		k = c % 4
		l = (32 + 2*e + 2*i - h - k) % 7
		m = (a + 11*h + 22*l) / 451
		r = 22 + h + l - 7*m
	} else {
		b = year % 7
		c = year % 4
		d = (19*a + 15) % 30
		e = (2*c + 4*b - d + 34) % 7
		r = 22 + d + e
	}

	var dh DateHoliday
	dh.Date = time.Date(year, time.March, r, 0, 0, 0, 0, time.Local)
	dh.Day = dh.Date.Day()
	dh.Month = int(dh.Date.Month())
	dh.Year = dh.Date.Year()
	return &dh, nil

}

func GetCarnivalByYear(year int) (*DateHoliday, error) {
	easter, errEaster := GetEasterCatholicByYear(year)
	if errEaster == nil {
		var dh DateHoliday
		dh.Date = easter.Date.Add(-47 * 24 * time.Hour)
		dh.Day = dh.Date.Day()
		dh.Month = int(dh.Date.Month())
		dh.Year = dh.Date.Year()
		return &dh, nil
	} else {
		return nil, errEaster
	}
}

func generateKeys() {
	a, e := GetCarnivalByYear(2050)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	fmt.Println(a)

	pubKey, privKey, _ := generateKey()
	os.WriteFile("public.key", []byte(pubKey), 0644)
	os.WriteFile("private.key", []byte(privKey), 0644)

	bytesPubKey, err := os.ReadFile("public.key")
	if err != nil {
		panic(err)
	}

	sourceText, err := os.ReadFile("source.txt")
	if err != nil {
		panic(err)
	}

	encrypted, err := encrypt(string(sourceText), string(bytesPubKey))
	if err != nil {
		panic(err)
	}
	fmt.Println("encrypted\n", encrypted)
	os.WriteFile("text.enc", []byte(encrypted), 0644)

	bytesPrivKey, err := os.ReadFile("private.key")
	if err != nil {
		panic(err)
	}
	decrypted, err := decrypt(encrypted, string(bytesPrivKey))
	if err != nil {
		panic(err)
	}
	fmt.Println("decrypted\n", decrypted)
	os.WriteFile("text.dec", []byte(decrypted), 0644)

	//Chave Gerada

	pbk, errPbk := publicKeyFromPrivateKey(string(bytesPrivKey))
	if errPbk != nil {
		fmt.Println("não foi possivel gerar a chave publica a partir da privada")
	}
	os.WriteFile("publicKeyFromPrivateKey.key", []byte(pbk), 0644)

	encryptedGerada, errEncryptedGerada := encrypt(string(sourceText), string(bytesPubKey))
	if errEncryptedGerada != nil {
		fmt.Println("não foi possivel encriptar com a chave publica gerada")
	}
	fmt.Println("encriptado com chave publica gerada\n", encryptedGerada)

	decryptedGerada, errDecryptedGerada := decrypt(encryptedGerada, string(bytesPrivKey))
	if errDecryptedGerada != nil {
		fmt.Println("não foi possivel desencriptar com a chave publica gerada")
	}
	fmt.Println("desencriptado com chave privada a partir de uma criptografia com a chave publica gerada gerada\n", decryptedGerada)

}
