package otpauth

import (
	"crypto"
	"fmt"
	"testing"
)

func TestHMAC(t *testing.T) {
	t.Run("UnavailableHashAlgorithm", func(t *testing.T) {
		t.Parallel()
		_, err := HMAC(crypto.MD5, []byte("12345678901234567890"), Itob(0))
		if err == nil || err.Error() != "Unsupported hash algorithm" {
			t.Fail()
			return
		}
	})
	for i, h := range []string{
		"cc93cf18508d94934c64b65d8ba7667fb7cde4b0",
		"75a48a19d4cbe100644e8ac1397eea747a2d33ab",
		"0bacb7fa082fef30782211938bc1c5e70416ff44",
		"66c28227d03a2d5529262ff016a1e6ef76557ece",
		"a904c900a64b35909874b33e61c5938a8e15ed1c",
		"a37e783d7b7233c083d4f62926c7a25f238d0316",
		"bc9cd28561042c83f219324d3c607256c03272ae",
		"a4fb960c0bc06e1eabb804e5b397cdc4b45596fa",
		"1b3c89f65e6c9e883012052823443f048b4332db",
		"1637409809a679dc698207310c8c7fc07290d9e5",
	} {
		i, h := i, h
		t.Run(fmt.Sprintf("%d to %s", i, h), func(t *testing.T) {
			t.Parallel()
			if b, _ := HMAC(crypto.SHA1, []byte("12345678901234567890"), Itob(int64(i))); fmt.Sprintf("%x", b) != h {
				t.Errorf("HMAC(crypto.SHA1, []byte(\"12345678901234567890\"), Itob(%d)) = %x; want %s", i, b, h)
			}
		})
	}
}
