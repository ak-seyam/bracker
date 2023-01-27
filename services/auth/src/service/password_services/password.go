package password_services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Argon2Identifier struct {
	format  string
	version int
	time    uint32
	memory  uint32
	keyLen  uint32
	saltLen uint32
	threads uint8
}

func NewArgon2Identifier() Argon2Identifier {
	return Argon2Identifier{
		format:  "$2$v=%d$m=%d,t=%d,p=%d$%s$%s",
		version: argon2.Version,
		time:    1,
		memory:  64 * 1024,
		keyLen:  32,
		saltLen: 16,
		threads: 2,
	}
}

func (aid Argon2Identifier) HashPassword(plainText string) (string, error) {
	salt := make([]byte, aid.saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(plainText), salt, aid.time, aid.memory, aid.threads, aid.keyLen)
	return fmt.Sprintf(
		aid.format,
		aid.version,
		aid.memory,
		aid.time,
		aid.threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	), nil
}
