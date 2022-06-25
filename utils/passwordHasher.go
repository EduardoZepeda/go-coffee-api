package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type DjangoPasswordHash struct {
	Algorithm  string
	Iterations int
	Size       int
	Salt       string
	Hash       string
}

type DjangoPasswordComparer struct {
	Hash DjangoPasswordHash
}

func GenerateDjangoPasswordHash(hash string) *DjangoPasswordHash {
	// Django stores the passwords using the following format <algorithm>$<iterations>$<salt>$<hash>
	// https://docs.djangoproject.com/en/4.0/topics/auth/passwords/
	arr := strings.Split(hash, "$")
	iterations, _ := strconv.Atoi(arr[1])
	return &DjangoPasswordHash{
		Algorithm:  arr[0],
		Iterations: iterations,
		Salt:       arr[2],
		Hash:       arr[3],
	}
}

func NewDjangoPasswordHash(hash string, size int, iterations int, algorithm string) *DjangoPasswordHash {
	return &DjangoPasswordHash{
		Hash:       hash,
		Size:       sha256.Size,
		Iterations: iterations,
		Algorithm:  algorithm,
	}
}

func NewDjangoPassword(hash string) *DjangoPasswordComparer {
	return &DjangoPasswordComparer{
		Hash: *GenerateDjangoPasswordHash(hash),
	}
}

func (PasswordHasher *DjangoPasswordComparer) VerifyPassword(password string) bool {
	hashedPassword := base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(password), []byte(PasswordHasher.Hash.Salt), PasswordHasher.Hash.Iterations, sha256.Size, sha256.New))
	return bytes.Equal([]byte(PasswordHasher.Hash.Hash), []byte(hashedPassword))
}

func (PasswordHasher *DjangoPasswordComparer) GetDjangoDefaultEncoding() string {
	return "pbkdf2_sha256"
}
