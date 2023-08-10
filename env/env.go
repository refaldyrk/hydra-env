package env

import (
	"fmt"

	"github.com/refaldyrk/hydra-env/helper"
)

type Env struct {
	hc *helper.HelperConfig
}

func DefaultEnv() *Env {
	return &Env{}
}

func (e *Env) SetHelperConfig(filePath string, hcc *helper.HelperConfig) {
	if hcc == nil {
		fmt.Println("[HYDRA] Error When Set Helper Config: Myb You Key.JSON Doesn't Exists")
		return
	}
	hcc.FilePath = filePath
	e.hc = hcc
}

func (e *Env) CreateKeyFile(data map[string]interface{}) error {
	existingData, err := e.hc.ReadOrCreateJSON()
	if err != nil {
		return err
	}

	for key, value := range data {
		existingData[key] = value
	}

	return e.hc.WriteJSON(existingData)
}

func (e *Env) AddKey(key, value string) error {
	encryptedValue, err := e.hc.Encrypt(value)
	if err != nil {
		return err
	}

	data, err := e.hc.ReadJSON()
	if err != nil {
		return err
	}

	data[key] = encryptedValue

	return e.hc.WriteJSON(data)
}
func (e *Env) GetExistingData() (map[string]interface{}, error) {
	return e.hc.ReadJSON()
}
func (e *Env) AddKeyToFile(key, value string) error {
	data, err := e.hc.ReadJSON()
	if err != nil {
		return err
	}

	encryptedValue, err := e.hc.Encrypt(value)
	if err != nil {
		return err
	}

	data[key] = encryptedValue

	return e.hc.WriteJSON(data)
}

func (e *Env) GetKey(key string) (string, error) {
	data, err := e.hc.ReadJSON()
	if err != nil {
		return "", err
	}

	if encryptedValue, ok := data[key].(string); ok {
		decryptedValue, err := e.hc.Decrypt(encryptedValue)
		if err != nil {
			return "", err
		}
		return decryptedValue, nil
	}

	return "", fmt.Errorf("key '%s' not found", key)
}

func (e *Env) UpdateKey(key, value string) error {
	encryptedValue, err := e.hc.Encrypt(value)
	if err != nil {
		return err
	}

	data, err := e.hc.ReadJSON()
	if err != nil {
		return err
	}

	if _, exists := data[key]; exists {
		data[key] = encryptedValue
		return e.hc.WriteJSON(data)
	}

	return fmt.Errorf("key '%s' not found", key)
}

func (e *Env) DeleteKey(key string) error {
	data, err := e.hc.ReadJSON()
	if err != nil {
		return err
	}

	if _, exists := data[key]; exists {
		delete(data, key)
		return e.hc.WriteJSON(data)
	}

	return fmt.Errorf("key '%s' not found", key)
}
