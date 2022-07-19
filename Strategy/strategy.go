package main

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
