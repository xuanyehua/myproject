package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

type Password struct {

}

const letterBytes = "abcdefghijklmnopqrstuvwxyz123456789"

func (p*Password)RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (p*Password)Encode_Password(psw string)  (md5_psw string ,salt string){

	salt = p.RandStringBytes(5)
	//fmt.Println(salt)
	md5_psw_str := p.Encode_Md5(psw)
	salt_psw := md5_psw_str + salt
	md5_psw = p.Encode_Md5(salt_psw)
	return md5_psw ,salt
}
func (p*Password)Encode_Md5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func (p*Password)Valid_Passwor(md5_psw string,psw string,salt string) bool{
	md5_psw_str := p.Encode_Md5(psw)
	salt_psw := md5_psw_str + salt
	md5_psw_1 := p.Encode_Md5(salt_psw)
	if md5_psw == md5_psw_1 {
		return true
	}
	return false
}

