package utils

import (
	"crypto/rand"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
	"testing"
)

func TestEncryptKey(t *testing.T) {
	key := make([]byte, 16)
	rand.Read(key)
	fmt.Printf("%x\n", string(key))
}
func TestGetUUid(t *testing.T) {
	key := uuid.NewV4().String()
	token := strings.ReplaceAll(key, "-", "")
	fmt.Println(key)
	fmt.Printf("%x\n", token)
}
