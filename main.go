package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
			panic(err)
		}
	} else {
		fmt.Println("[INFO] Env No Set")
	}

	if *genKeyFlag {
		key.NewKey()

		err := key.CreateKeyFile()
		if err != nil {
			panic(err)
			return
		}

		err = key.PrintKey()
		if err != nil {
			panic(err)
			return
		}

		return
	}

	if *addKeyFlag != "" {
		if *envFlag == "" {
			fmt.Println("Error: --env flag is required for -add-key")
			return
		}

		parts := strings.Split(*addKeyFlag, "|")
		if len(parts) != 2 {
			fmt.Println("Invalid format for add-key")
			return
		}

		keys := parts[0]
		value := parts[1]

		err := envis.AddKeyToFile(keys, value)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Key '%s' added successfully.\n", keys)
	}

	if *getKeyFlag != "" {
		if *envFlag == "" {
			fmt.Println("Error: --env flag is required for -get-key")
			return
		}

		value, err := envis.GetKey(*getKeyFlag)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Value of key '%s': %s\n", *getKeyFlag, value)
	}

	if *listKeysFlag {
		if *envFlag == "" {
			fmt.Println("Error: --env flag is required for -list-keys")
			return
		}

		data, err := envis.GetExistingData()
		if err != nil {
			panic(err)
		}

		fmt.Println("List of keys:")
		for key := range data {
			fmt.Println(key)
		}
	}

	if *delKeyFlag != "" {
		if *envFlag == "" {
			fmt.Println("Error: --env flag is required for -add-key")
			return
		}

		err := envis.DeleteKey(*delKeyFlag)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Key '%s' deleted successfully.\n", *delKeyFlag)
	}

	if *loadEnvFlag != "" {
		files, err := os.Open(*loadEnvFlag)
		if err != nil {
			panic(err)
		}

		defer files.Close()
		data, err := ioutil.ReadAll(files)
		if err != nil {
			panic(err)
		}
		dataString := string(data)
		arrNewLine := strings.Split(dataString, "\n")
		for _, v := range arrNewLine {
			arrEnv := strings.Split(v, "=")
			if v != "" {
				err := envis.AddKeyToFile(arrEnv[0], arrEnv[1])
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println("Add Key: ", arrEnv[0])
			}
		}
		fmt.Println("[Hydra] Everything Is Load...")
	}

}
