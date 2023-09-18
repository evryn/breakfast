package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"version-forge/config"
	"version-forge/modules"
)

func main() {
	config.MustLoadConfig()

	if len(os.Args) == 1 {
		modules.MustListen(modules.MustGetPrepareTemplates(config.Main.MockedVersion))
	}

	switch os.Args[1] {
	case "test":
		for _, v := range config.Main.Versions {
			td, lt, st := modules.MustGetPrepareTemplates(v.Version)

			if err := lt.Execute(io.Discard, td); err != nil {
				log.Fatalf("Unable to execute long template for version %s: %v", v.Version, err)
			}

			if err := st.Execute(io.Discard, td); err != nil {
				log.Fatalf("Unable to execute short template for version %s: %v", v.Version, err)
			}
		}
	case "build":
		os.RemoveAll("dist")
		os.Mkdir("dist", 0755)

		for _, v := range config.Main.Versions {
			td, lt, st := modules.MustGetPrepareTemplates(v.Version)

			longFile, _ := os.Create(fmt.Sprintf("dist/%s-long.txt", v.Version))
			shortFile, _ := os.Create(fmt.Sprintf("dist/%s-short.txt", v.Version))

			if err := lt.Execute(longFile, td); err != nil {
				log.Fatalf("Unable to execute long template for version %s: %v", v.Version, err)
			}

			if err := st.Execute(shortFile, td); err != nil {
				log.Fatalf("Unable to execute short template for version %s: %v", v.Version, err)
			}
		}
	default:
		log.Fatalf("Unknown argument: %s", os.Args[1])
	}
}
