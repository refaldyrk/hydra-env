package helper

import (
	"errors"

	"github.com/mervick/aes-everywhere/go/aes256"
)

// The `func (hc *HelperConfig) SetPasswordAES(password string) error` is a method of the `HelperConfig`
// struct. It takes a `password` string as a parameter and returns an error. This method is used to set
// the password in the `HelperConfig` struct. If the `password` is empty, it returns an error
// indicating that the password cannot be empty. Otherwise, it sets the `password` in the
// `HelperConfig` struct and returns `nil` to indicate that there was no error.
func (hc *HelperConfig) SetPasswordAES(password string) error {
	if password == "" {
		return errors.New("password can't be empty")
	}

	hc.Password = password
	return nil
}

// The `func (hc *HelperConfig) Encrypt(v string) (string, error) {` is a method of the `HelperConfig`
// struct. It takes a string `v` as a parameter and returns a string and an error. This method is used
// to encrypt the input string `v` using AES encryption with the password stored in the `HelperConfig`
// struct.
func (hc *HelperConfig) Encrypt(v string) (string, error) {
	if hc.Password == "" {
		return "", errors.New("password empty")
	}

	result := aes256.Encrypt(v, hc.Password)

	return result, nil
}

// The `func (hc *HelperConfig) Decrypt(v string) (string, error) {` is a method of the `HelperConfig`
// struct. It takes a string `v` as a parameter and returns a string and an error. This method is used
// to decrypt the input string `v` using AES decryption with the password stored in the `HelperConfig`
// The `ReadJSON` function is a method of the `HelperConfig` struct. It reads JSON data from a file
// specified by the `FilePath` field of the `HelperConfig` struct. It returns a map of type
// `map[string]interface{}` that represents the JSON data, and an error if any occurred during the
// reading process.
// struct.
func (hc *HelperConfig) Decrypt(v string) (string, error) {
	if hc.Password == "" {
		return "", errors.New("password empty")
	}

	result := aes256.Decrypt(v, hc.Password)

	return result, nil
}
