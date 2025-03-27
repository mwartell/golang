package weasel

import (
	"math/rand"
	"strings"
)

const Target = "METHINKS IT IS LIKE A WEASEL"
const Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "

func RandomString(length int) string {
	var sb strings.Builder
	for range length {
		sb.WriteByte(Charset[rand.Intn(len(Charset))])
	}
	return sb.String()
}
