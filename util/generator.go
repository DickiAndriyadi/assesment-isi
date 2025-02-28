package util

import (
	"crypto/rand"
	"fmt"
)

func GenerateNoRekening() string {
	b := make([]byte, 5)
	rand.Read(b)
	return fmt.Sprintf("99%X", b)
}
