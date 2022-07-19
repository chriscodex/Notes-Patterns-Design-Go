package main

import (
	emd5 "crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type passwordProtector struct {
	user          string
	passwordName  string
	password      string
	hashAlgorithm hashAlgorithm
}

type hashAlgorithm interface {
	hash(p *passwordProtector)
}

func newPasswordProtector(user string, passwordName string, password string, hash hashAlgorithm) *passwordProtector {
	return &passwordProtector{
		user:          user,
		passwordName:  passwordName,
		password:      password,
		hashAlgorithm: hash,
	}
}

func (p *passwordProtector) setHashAlgorith(hash hashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *passwordProtector) protect() {
	p.hashAlgorithm.hash(p)
}

type sha struct{}

func (sha) hash(p *passwordProtector) {
	h := sha1.New()
	h.Write([]byte(p.password))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using SHA for %s: %s\n", p.passwordName, sha1Hash)
}

type md5 struct{}

func (md5) hash(p *passwordProtector) {
	h := emd5.New()
	h.Write([]byte(p.password))
	md5Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using MD5 for %s: %s\n", p.passwordName, md5Hash)
}

func main() {
	sha := &sha{}
	md5 := &md5{}

	passwordProtector := newPasswordProtector("Christian", "Steam password", "1324", sha)
	passwordProtector.protect()
	passwordProtector.setHashAlgorith(md5)
	passwordProtector.protect()
}
