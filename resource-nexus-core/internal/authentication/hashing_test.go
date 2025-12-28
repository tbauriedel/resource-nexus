package authentication

import (
	"strings"
	"testing"

	"github.com/tbauriedel/resource-nexus-core/internal/config"
)

func TestHashPasswordString(t *testing.T) {
	p := config.HashingParams{
		Iterations:   1,
		MemoryCost:   16 * 1024,
		ThreadsCount: 1,
		KeyLength:    32,
		SaltLength:   16,
	}

	hash1 := HashPasswordString("dummy", p)

	if !strings.HasPrefix(hash1, "$argon2id$v=19$m=16384,t=1,p=1") {
		t.Fatal("wrong hash format")
	}

	hash2 := HashPasswordString("dummy", p)

	if hash1 == hash2 {
		t.Fatal("hashes should be different because of the random salt!")
	}

}
