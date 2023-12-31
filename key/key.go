package key

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"log"

	"github.com/google/uuid"
	"github.com/refaldyrk/hydra-env/helper"
)

type Key struct {
	hc  *helper.HelperConfig
	key string
}

func (k *Key) ReadJSON() (map[string]interface{}, error) {
	file, err := os.Open("key.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func DefaultKey() *Key {
	hc := helper.NewHelperConfig(&helper.HelperConfig{FilePath: "key.json"})
	return &Key{key: "hydra", hc: hc}
}

func (k *Key) NewKey() {
	k.key = uuid.NewString()
	k.hc.Password = k.key
}

type KeyStruct struct {
	Key string `json:"key"`
}

func (k *Key) GetHC() *helper.HelperConfig {
	passwd, err := k.GetKey()
	if err != nil {
		return nil
	}
	k.hc.Password = passwd
	return k.hc
}
func (k *Key) CreateKeyFile() error {
	keyData := KeyStruct{Key: k.key}
	data := map[string]interface{}{"key": keyData.Key}

	err := k.hc.WriteJSON(data)
	if err != nil {
		return err
	}
	return nil
}

func (k *Key) GetKey() (string, error) {
	fileData, err := k.hc.ReadJSON()
	if err != nil {
		return "", err
	}

	var keyData KeyStruct
	if keyVal, ok := fileData["key"].(string); ok {
		keyData.Key = keyVal
		k.key = keyVal
	} else {
		return "", errors.New("key not found or invalid")
	}

	return k.key, nil
}

func (k *Key) PrintKey() error {
	fileData, err := k.hc.ReadJSON()
	if err != nil {
		return err
	}

	var keyData KeyStruct
	if keyVal, ok := fileData["key"].(string); ok {
		keyData.Key = keyVal
		log.Println("[HYDRA] New Key: ", keyVal)
	} else {
		return errors.New("key not found or invalid")
	}

	return nil
}

func (k *Key) CheckKeyMatch(checkKey string) (bool, error) {
	fileData, err := k.hc.ReadJSON()
	if err != nil {
		return false, err
	}

	var keyData KeyStruct
	if keyVal, ok := fileData["key"].(string); ok {
		keyData.Key = keyVal
	} else {
		return false, errors.New("key not found or invalid")
	}

	return checkKey == keyData.Key, nil
}
