package crypto

import (
	"os"
	"testing"
)

func TestSHA1String(t *testing.T) {
	result := SHA1String("hello")
	if len(result) != 40 {
		t.Errorf("Expected SHA1 hash length 40, got %d", len(result))
	}
}

func TestSHA256String(t *testing.T) {
	result := SHA256String("hello")
	if len(result) != 64 {
		t.Errorf("Expected SHA256 hash length 64, got %d", len(result))
	}
}

func TestSHA512String(t *testing.T) {
	result := SHA512String("hello")
	if len(result) != 128 {
		t.Errorf("Expected SHA512 hash length 128, got %d", len(result))
	}
}

func TestSHA1File(t *testing.T) {
	tmpFile := "test_sha1.txt"
	os.WriteFile(tmpFile, []byte("test"), 0644)
	defer os.Remove(tmpFile)
	hash, err := SHA1File(tmpFile)
	if err != nil {
		t.Fatalf("SHA1File failed: %v", err)
	}
	if len(hash) != 40 {
		t.Errorf("Expected SHA1 hash length 40, got %d", len(hash))
	}
}

func TestAESEncryptDecrypt(t *testing.T) {
	key := []byte("1234567890123456")
	plaintext := []byte("hello world")
	ciphertext, err := AESEncrypt(key, plaintext)
	if err != nil {
		t.Fatalf("AESEncrypt failed: %v", err)
	}
	decrypted, err := AESDecrypt(key, ciphertext)
	if err != nil {
		t.Fatalf("AESDecrypt failed: %v", err)
	}
	if string(decrypted) != string(plaintext) {
		t.Errorf("Expected %q, got %q", string(plaintext), string(decrypted))
	}
}

func TestAESEncryptDecryptString(t *testing.T) {
	key := "1234567890123456"
	plaintext := "hello world"
	ciphertext, err := AESEncryptString(key, plaintext)
	if err != nil {
		t.Fatalf("AESEncryptString failed: %v", err)
	}
	decrypted, err := AESDecryptString(key, ciphertext)
	if err != nil {
		t.Fatalf("AESDecryptString failed: %v", err)
	}
	if decrypted != plaintext {
		t.Errorf("Expected %q, got %q", plaintext, decrypted)
	}
}

func TestBase64EncodeDecode(t *testing.T) {
	data := []byte("hello world")
	encoded := Base64Encode(data)
	decoded, err := Base64Decode(encoded)
	if err != nil {
		t.Fatalf("Base64Decode failed: %v", err)
	}
	if string(decoded) != string(data) {
		t.Errorf("Expected %q, got %q", string(data), string(decoded))
	}
}

func TestBase64EncodeDecodeString(t *testing.T) {
	s := "hello world"
	encoded := Base64EncodeString(s)
	decoded, err := Base64DecodeString(encoded)
	if err != nil {
		t.Fatalf("Base64DecodeString failed: %v", err)
	}
	if decoded != s {
		t.Errorf("Expected %q, got %q", s, decoded)
	}
}
