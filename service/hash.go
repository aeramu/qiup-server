package service

import(
	"crypto/sha1"
	"fmt"
)

func Hash(s string)(string){
	hash := sha1.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x",hash.Sum(nil))
}