package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const DJANGO_DEFAULT_SALT_SIZE int = 22
const DJANGO_DEFAULT_ITERATIONS int = 320000
const DJANGO_DEFAULT_ALGORITHM string = "pbkdf2_sha256"
const DJANGO_DEFAULT_ALLOWED_CHAR_SET string = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func (PasswordHasher *DjangoPasswordComparer) VerifyPassword(password string) bool {
	hashedPassword := base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(password), []byte(PasswordHasher.Hash.Salt), PasswordHasher.Hash.Iterations, sha256.Size, sha256.New))
	// Prevent timing attacks
	if subtle.ConstantTimeCompare([]byte(PasswordHasher.Hash.Hash), []byte(hashedPassword)) == 0 {
		return false
	}
	return true
}

func (PasswordHasher *DjangoPasswordComparer) GetDjangoDefaultEncoding() string {
	return "pbkdf2_sha256"
}

func NewDjangoPassword(hash string) (*DjangoPasswordComparer, error) {
	djangoHash, err := GenerateDjangoPasswordHash(hash)
	if err != nil {
		return nil, err
	}
	return &DjangoPasswordComparer{
		Hash: *djangoHash,
	}, nil
}

func GenerateDjangoPasswordHash(hash string) (*DjangoPasswordHash, error) {
	// Django stores the passwords using the following format <algorithm>$<iterations>$<salt>$<hash>
	// https://docs.djangoproject.com/en/4.0/topics/auth/passwords/
	arr := strings.Split(hash, "$")
	if len(arr) != 4 {
		return nil, errors.New("The hash received is not in the format: <algorithm>$<iterations>$<salt>$<hash>")
	}
	iterations, _ := strconv.Atoi(arr[1])
	return &DjangoPasswordHash{
		Algorithm:  arr[0],
		Iterations: iterations,
		Salt:       arr[2],
		Hash:       arr[3],
	}, nil
}

func NewDjangoPasswordHash(password string) (*DjangoPasswordHash, error) {
	asciiSalt, err := GenerateDjangoSalt()
	if err != nil {
		return nil, err
	}
	passwordHash := base64.StdEncoding.EncodeToString(pbkdf2.Key([]byte(password), []byte(asciiSalt), DJANGO_DEFAULT_ITERATIONS, sha256.Size, sha256.New))
	return &DjangoPasswordHash{
		Hash:       passwordHash,
		Size:       sha256.Size,
		Salt:       asciiSalt,
		Iterations: DJANGO_DEFAULT_ITERATIONS,
		Algorithm:  DJANGO_DEFAULT_ALGORITHM,
	}, nil
}

func GenerateDjangoHashedPassword(password string) (string, error) {
	HashedPassword, err := NewDjangoPasswordHash(password)
	if err != nil {
		return "", err
	}
	// Django stores the passwords using the following format <algorithm>$<iterations>$<salt>$<hash>
	return fmt.Sprintf("%s$%d$%s$%s", HashedPassword.Algorithm, HashedPassword.Iterations, HashedPassword.Salt, HashedPassword.Hash), nil
}

func GenerateRandomString(length int, chars string) (string, error) {
	result := make([]rune, length)
	runes := []rune(chars)
	x := int64(len(runes))
	for i := range result {
		// rand.Int requires a bigInt
		num, err := rand.Int(rand.Reader, big.NewInt(x))
		if err != nil {
			return "", errors.New("Error generating a random number.")
		}
		result[i] = runes[num.Int64()]
	}
	return string(result), nil
}

func GenerateDjangoSalt() (string, error) {
	// Taken from django docs https://github.com/django/django/blob/main/django/utils/crypto.py:
	// The bit length of the returned value can be calculated with the formula:
	//     log_2(len(allowed_chars)^length)
	// For example, with default `allowed_chars` (26+26+10), this gives:
	//   * length: 12, bit length =~ 71 bits
	//   * length: 22, bit length =~ 131 bits
	return GenerateRandomString(DJANGO_DEFAULT_SALT_SIZE, DJANGO_DEFAULT_ALLOWED_CHAR_SET)
}
