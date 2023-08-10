package main

import (
	"flag"

	"github.com/refaldyrk/hydra-env/env"
	"github.com/refaldyrk/hydra-env/key"
	"github.com/refaldyrk/hydra-env/present"
)

func main() {
	initFlag()
}

func initFlag() {
	key := key.DefaultKey()
	envis := env.DefaultEnv()
	presents := present.NewPresent(envis, key)

	genKeyFlag := flag.Bool("gen-key", false, "Generate and print a new key")
	envFlag := flag.String("env", "", "Specify the environment file path")
	addKeyFlag := flag.String("add-key", "", "Add a new key in the format 'key|value'")
	getKeyFlag := flag.String("get-key", "", "Get the value of a key by specifying the key name")
	loadEnvFlag := flag.String("load-env", "", "Load environment variables from a file")
	listKeysFlag := flag.Bool("list-keys", false, "List all keys in the environment file")
	delKeyFlag := flag.String("del-key", "", "Delete a key by specifying the key name")

	flag.Parse()

	if *envFlag != "" {
		presents.EnvFlagPresent(*envFlag)
	}

	if *genKeyFlag {
		presents.GenKeyFlag()
		return
	}

	if *addKeyFlag != "" {
		presents.AddKeyFlag(*envFlag, *addKeyFlag)
	}

	if *getKeyFlag != "" {
		presents.GetKeyFlag(*envFlag, *getKeyFlag)
	}

	if *listKeysFlag {
		presents.ListKeyFlag(*envFlag)
	}

	if *delKeyFlag != "" {
		presents.DelKeyFlag(*envFlag, *delKeyFlag)
	}

	if *loadEnvFlag != "" {
		presents.LoadEnvFlag(*loadEnvFlag)
	}

}
