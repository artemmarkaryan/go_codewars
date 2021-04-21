package hashing

import (
	"crypto/md5"
	"fmt"
)

func PassHash(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func main() {

}
