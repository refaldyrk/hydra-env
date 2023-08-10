package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"log"

	"github.com/refaldyrk/hydra-env/env"
	"github.com/refaldyrk/hydra-env/key"
)

func main() {
	key := key.DefaultKey()
	envis := env.DefaultEnv()

	genKeyFlag := flag.Bool("gen-key", false, "Generate and print a new key")
	envFlag := flag.String("env", "", "Specify the environment file path")
	addKeyFlag := flag.String("add-key", "", "Add a new key in the format 'key|value'")
	getKeyFlag := flag.String("get-key", "", "Get the value of a key by specifying the key name")
	loadEnvFlag := flag.String("load-env", "", "Load environment variables from a file")
	listKeysFlag := flag.Bool("list-keys", false, "List all keys in the environment file")
	delKeyFlag := flag.String("del-key", "", "Delete a key by specifying the key name")

	flag.Parse()

	if *envFlag != "" {
		envis.SetHelperConfig(*envFlag, key.GetHC())
		err := envis.CreateKeyFile(map[string]interface{}{})
		if err != nil {
			log.Fatal("[HYDRA] Failed To Create Key File Env", err.Error())
		}
	}

	if *genKeyFlag {
		key.NewKey()

		err := key.CreateKeyFile()
		if err != nil {
			log.Fatal("[HYDRA] Error Create Key File: ", err.Error())
			return
		}

		err = key.PrintKey()
		if err != nil {
			log.Fatal("[HYDRA] Error Print Key", err)
			return
		}

		return
	}

	if *addKeyFlag != "" {
		if *envFlag == "" {
			log.Fatal("[HYDRA] Require flag --env")
			return
		}

		parts := strings.Split(*addKeyFlag, "|")
		if len(parts) != 2 {
			log.Fatal("[HYDRA] Invalid Command For Add Key: ", addKeyFlag)
			return
		}

		keys := parts[0]
		value := parts[1]

		err := envis.AddKeyToFile(keys, value)
		if err != nil {
			log.Fatal("[HYDRA] Error When Add Key: ", err)
			return
		}

		log.Println("[HYDRA] Success Add Key: ", keys)
	}

	if *getKeyFlag != "" {
		if *envFlag == "" {
			log.Fatal("[HYDRA] Require flag --env")
			return
		}

		value, err := envis.GetKey(*getKeyFlag)
		if err != nil {
			log.Fatal("[HYDRA] ERROR: ", err)
		}

		log.Printf("[HYDRA] Value of key '%s': %s\n", *getKeyFlag, value)
	}

	if *listKeysFlag {
		if *envFlag == "" {
			log.Fatal("[HYDRA] Require flag --env")
			return
		}

		data, err := envis.GetExistingData()
		if err != nil {
			log.Fatal("[HYDRA] ERROR: ", err)
		}

		log.Println("[HYDRA] List of keys:")
		for key := range data {
			fmt.Println(key)
		}
		log.Println("[HYDRA] Key: ", len(data))
	}

	if *delKeyFlag != "" {
		if *envFlag == "" {
			log.Fatal("[HYDRA] Require flag --env")
			return
		}

		err := envis.DeleteKey(*delKeyFlag)
		if err != nil {
			log.Fatal("[HYDRA] ERROR: ", err)
		}

		log.Printf("Key '%s' deleted successfully.\n", *delKeyFlag)
	}

	if *loadEnvFlag != "" {
		files, err := os.Open(*loadEnvFlag)
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
				err := envis.AddKeyToFile(arrEnv[0], arrEnv[1])
				if err != nil {
					log.Fatal("[HYDRA] ERROR: ", err)
					continue
				}
				log.Println("[HYDRA] Add Key: ", arrEnv[0])
			}
		}
		log.Println("[HYDRA] Success Load Env ", loadEnvFlag)
	}

}
