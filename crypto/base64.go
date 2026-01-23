package crypto

import "encoding/base64"

// Base64Encode Base64编码 / Base64 encoding
func Base64Encode(data []byte) []byte {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, data)
	return encoded
}

// Base64Decode Base64解码 / Base64 decoding
func Base64Decode(data []byte) ([]byte, error) {
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(decoded, data)
	if err != nil {
		return nil, err
	}
	return decoded[:n], nil
}

// Base64EncodeString Base64编码（字符串）/ Base64 encoding (string)
func Base64EncodeString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64DecodeString Base64解码（字符串）/ Base64 decoding (string)
func Base64DecodeString(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
