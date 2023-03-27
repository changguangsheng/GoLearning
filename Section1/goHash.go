package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestString := "Hi,pandaman!"

	md5Inst := md5.New()

	md5Inst.Write([]byte(TestString))
	Result := md5Inst.Sum([]byte(""))

	fmt.Printf("%x\n\n", Result)

	Sha1Int := sha1.New()
	Sha1Int.Write([]byte(TestString))
	Result = Sha1Int.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}
