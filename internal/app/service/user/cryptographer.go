package service

type Cryptographer interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}
