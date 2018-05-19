package targetprocess

import (
	"encoding/base64"
)

func encodeAuth(user string, pass string) string {
	input := []byte(user + ":" + pass)
	b64Val := base64.StdEncoding.EncodeToString(input)
	return b64Val
}
