package main

import "fmt"

type passwordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm hashAlgorithm
}

type hashAlgorithm interface {
	hash(p *passwordProtector)
}

func newPasswordProtector(user string, passwordName string, hash hashAlgorithm) *passwordProtector {
	return &passwordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

func (p *passwordProtector) setHashAlgorith(hash hashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *passwordProtector) hash() {
	p.hashAlgorithm.hash(p)
}

type sha struct{}

func (sha) hash(p *passwordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

type md5 struct{}

func (md5) hash(p *passwordProtector) {
	fmt.Printf("Hashing using MD5 for %s\n", p.passwordName)
}
