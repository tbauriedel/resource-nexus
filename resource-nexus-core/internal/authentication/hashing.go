package authentication

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/tbauriedel/resource-nexus-core/internal/config"
	"golang.org/x/crypto/argon2"
)

// HashPasswordString hashes a password using argon2id.
// Salt and cost parameters are based on the recommendations from https://github.com/P-H-C/phc-winner-argon2
//
// Returns the hash as a byte slice.
func HashPasswordString(password string, params config.HashingParams) string {
	// generate salt
	salt := make([]byte, params.SaltLength)

	// Read random bytes from the crypto/rand package and fills the slice.
	// never generates an error. panics in case of failure
	_, _ = rand.Read(salt)

	// generate hash
	hash := argon2.IDKey([]byte(password),
		salt,
		params.Iterations,
		params.MemoryCost,
		params.ThreadsCount,
		params.KeyLength,
	)

	// encode salt and hash to base64
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// format to phc format: $argon2id$v=19$m=MEM,t=ITER,p=PAR$salt$hash
	encoded := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		params.MemoryCost, params.Iterations, params.ThreadsCount, b64Salt, b64Hash)

	return encoded
}
