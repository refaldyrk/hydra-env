package helper

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// The `SetFilePath` function is a method of the `HelperConfig` struct. It takes a `path` parameter of
// type string and sets the `FilePath` field of the `HelperConfig` struct to the provided path. This
// function is used to set the file path that will be used for reading and writing JSON data.
func (hc *HelperConfig) SetFilePath(path string) {
	hc.FilePath = path
}

func (hc *HelperConfig) ReadOrCreateJSON() (map[string]interface{}, error) {
	file, err := os.OpenFile(hc.FilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return make(map[string]interface{}), nil
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// The `ReadJSON` function is a method of the `HelperConfig` struct. It reads JSON data from a file
// specified by the `FilePath` field of the `HelperConfig` struct. It returns a map of type
// `map[string]interface{}` that represents the JSON data, and an error if any occurred during the
// reading process.
func (hc *HelperConfig) ReadJSON() (map[string]interface{}, error) {
	file, err := os.Open(hc.FilePath)
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

// The `WriteJSON` function is a method of the `HelperConfig` struct. It takes a `data` parameter of
// type `map[string]interface{}` which represents the JSON data to be written to a file.
func (hc *HelperConfig) WriteJSON(data map[string]interface{}) error {
	file, err := os.Create(hc.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// The `UpdateJSON` function is a method of the `HelperConfig` struct. It takes two parameters: `key`
// of type string and `value` of type interface{}.
func (hc *HelperConfig) UpdateJSON(key string, value interface{}) error {
	jsonData, err := hc.ReadJSON()
	if err != nil {
		return err
	}

	jsonData[key] = value

	err = hc.WriteJSON(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// The `DeleteJSON` function is a method of the `HelperConfig` struct. It takes a `key` parameter of
// type string and deletes the corresponding key-value pair from the JSON data stored in the file
// specified by the `FilePath` field of the `HelperConfig` struct. It returns an error if any occurred
// during the deletion process.
func (hc *HelperConfig) DeleteJSON(key string) error {
	jsonData, err := hc.ReadJSON()
	if err != nil {
		return err
	}

	delete(jsonData, key)

	err = hc.WriteJSON(jsonData)
	if err != nil {
		return err
	}

	return nil
}
