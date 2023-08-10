package present

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/refaldyrk/hydra-env/env"
	"github.com/refaldyrk/hydra-env/key"
)

type Present struct {
	envis *env.Env
	key   *key.Key
}

func NewPresent(env *env.Env, keys *key.Key) *Present {
	return &Present{envis: env, key: keys}
}

func (p *Present) EnvFlagPresent(flag string) {
	p.envis.SetHelperConfig(flag, p.key.GetHC())
	err := p.envis.CreateKeyFile(map[string]interface{}{})
	if err != nil {
		log.Fatal("[HYDRA] Failed To Create Key File Env", err.Error())
		return
	}
}

func (p *Present) GenKeyFlag() {
	p.key.NewKey()

	err := p.key.CreateKeyFile()
	if err != nil {
		log.Fatal("[HYDRA] Error Create Key File: ", err.Error())
		return
	}

	err = p.key.PrintKey()
	if err != nil {
		log.Fatal("[HYDRA] Error Print Key", err)
		return
	}

	return
}

func (p *Present) AddKeyFlag(envFlag string, addKeyFlag string) {
	if envFlag == "" {
		log.Fatal("[HYDRA] Require flag --env")
		return
	}

	parts := strings.Split(addKeyFlag, "|")
	if len(parts) != 2 {
		log.Fatal("[HYDRA] Invalid Command For Add Key: ", addKeyFlag)
		return
	}

	keys := parts[0]
	value := parts[1]

	err := p.envis.AddKeyToFile(keys, value)
	if err != nil {
		log.Fatal("[HYDRA] Error When Add Key: ", err)
		return
	}

	log.Println("[HYDRA] Success Add Key: ", keys)
	return
}

func (p Present) GetKeyFlag(envFlag string, getKeyFlag string) {
	if envFlag == "" {
		log.Fatal("[HYDRA] Require flag --env")
		return
	}

	value, err := p.envis.GetKey(getKeyFlag)
	if err != nil {
		log.Fatal("[HYDRA] ERROR: ", err)
		return
	}

	log.Printf("[HYDRA] Value of key '%s': %s\n", getKeyFlag, value)
	return
}

func (p *Present) ListKeyFlag(envFlag string) {
	if envFlag == "" {
		log.Fatal("[HYDRA] Require flag --env")
		return
	}

	data, err := p.envis.GetExistingData()
	if err != nil {
		log.Fatal("[HYDRA] ERROR: ", err)
		return
	}

	log.Println("[HYDRA] List of keys:")
	for key := range data {
		fmt.Println(key)
	}
	log.Println("[HYDRA] Key: ", len(data))
	return
}

func (p *Present) DelKeyFlag(envFlag string, delKeyFlag string) {
	if envFlag == "" {
		log.Fatal("[HYDRA] Require flag --env")
		return
	}

	err := p.envis.DeleteKey(delKeyFlag)
	if err != nil {
		log.Fatal("[HYDRA] ERROR: ", err)
	}

	log.Printf("Key '%s' deleted successfully.\n", delKeyFlag)
	return
}

func (p *Present) LoadEnvFlag(loadEnvFlag string) {
	files, err := os.Open(loadEnvFlag)
	if err != nil {
		log.Fatal("[HYDRA] ERROR: ", err)
	}

	defer files.Close()
	data, err := ioutil.ReadAll(files)
	if err != nil {
		log.Fatal("[HYDRA] ERROR: ", err)
	}
	dataString := string(data)
	arrNewLine := strings.Split(dataString, "\n")
	for _, v := range arrNewLine {
		if v == "" || v == "\n" {
			continue
		}
		if v != "" || v != "\n" {
			arrEnv := strings.Split(v, "=")
			if len(arrEnv) <= 1 {
				continue
			}
			err := p.envis.AddKeyToFile(arrEnv[0], arrEnv[1])
			if err != nil {
				log.Fatal("[HYDRA] ERROR: ", err)
				continue
			}
			log.Println("[HYDRA] Add Key: ", arrEnv[0])
		}
	}
	log.Println("[HYDRA] Success Load Env ", loadEnvFlag)
}
