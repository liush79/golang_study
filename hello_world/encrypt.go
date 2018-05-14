package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	"io"
)

func testEncrypt() {
	testSha512()
	testAES()
	testRSA()
}

func testSha512() {
	ss := "Hello, world!"
	h1 := sha512.Sum512([]byte(ss)) // 문자열의 SHA512 해시 값 추출
	fmt.Printf("%x\n", h1)
}

func testAES() {
	key := "my key 16 bytes."
	s := "abcdefghijklmnopqrstuvwxyz"
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := encrypt(block, []byte(s)) // AES 대칭키 암호화 블록 생성
	fmt.Printf("AES enc: %x\n", ciphertext)

	plaintext := decrypt(block, ciphertext) // AES 알고리즘 암호문을 평문으로 복호화
	fmt.Println("AES dec: " + string(plaintext))
}

func testRSA() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // 개인 키와 공개 키 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey // 개인 키 변수 안에 공개 키가 들어있음

	message := "안녕하세요. Go 언어"
	hash := md5.New()           // 해시 인스턴스 생성
	hash.Write([]byte(message)) // 해시 인스턴스에 문자열 추가
	digest := hash.Sum(nil)     // 문자열의 MD5 해시 값 추출

	var h1 crypto.Hash
	signature, err := rsa.SignPKCS1v15( // 개인 키로 서명
		rand.Reader,
		privateKey, // 개인 키
		h1,
		digest, // MD5 해시 값
	)

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15( // 공개 키로 서명 검증
		publicKey, // 공개 키
		h2,
		digest,    // MD5 해시 값
		signature, // 서명 값
	)

	if err != nil {
		fmt.Println("검증 실패")
	} else {
		fmt.Println("검증 성공")
	}
}

func encrypt(b cipher.Block, plaintext []byte) []byte {
	if mod := len(plaintext) % aes.BlockSize; mod != 0 { // 블록 크기의 배수가 되어야함
		padding := make([]byte, aes.BlockSize-mod) // 블록 크기에서 모자라는 부분을
		plaintext = append(plaintext, padding...)  // 채워줌
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext)) // 초기화 벡터 공간(aes.BlockSize)만큼 더 생성
	iv := ciphertext[:aes.BlockSize]                         // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {  // 랜덤 값을 초기화 벡터에 넣어줌
		fmt.Println(err)
		return nil
	}

	mode := cipher.NewCBCEncrypter(b, iv)                   // 암호화 블록과 초기화 벡터를 넣어서 암호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext) // 암호화 블록 모드 인스턴스로
	// 암호화

	return ciphertext
}

func decrypt(b cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 { // 블록 크기의 배수가 아니면 리턴
		fmt.Println("암호화된 데이터의 길이는 블록 크기의 배수가 되어야합니다.")
		return nil
	}

	iv := ciphertext[:aes.BlockSize]        // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	ciphertext = ciphertext[aes.BlockSize:] // 부분 슬라이스로 암호화된 데이터를 가져옴

	plaintext := make([]byte, len(ciphertext)) // 평문 데이터를 저장할 공간 생성
	mode := cipher.NewCBCDecrypter(b, iv)      // 암호화 블록과 초기화 벡터를 넣어서
	// 복호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(plaintext, ciphertext) // 복호화 블록 모드 인스턴스로 복호화

	return plaintext
}
