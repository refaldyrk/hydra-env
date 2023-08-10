package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/refaldyrk/hydra-env/env"
	"github.com/refaldyrk/hydra-env/key"
	"github.com/refaldyrk/hydra-env/present"
)

func main() {
	initFlag()
}

func initFlag() {
	//Get Environment HYDRA_MONGO_SERVER
	mongoURL := os.Getenv("HYDRA_MONGO_SERVER")
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

	//Server Flag
	serverFlag := flag.String("server", "", "run your server command, if u wanna see list command please send --server=help")

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

	if *serverFlag != "" {
		if mongoURL == "" {
			fmt.Println("[HYDRA] Not Found Your HYDRA_MONGO_SERVER: Set Your Environment And Try Again")
			return
		}
		presents.ServerFlag(*envFlag, *serverFlag, mongoURL)
	}
}
