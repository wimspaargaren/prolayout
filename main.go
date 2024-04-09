// Package main bootstrap the analyzer to validate the project structure.
package main

import (
	"log"
	"os"

	"golang.org/x/tools/go/analysis/singlechecker"
	"gopkg.in/yaml.v3"

	"github.com/wimspaargaren/prolayout/internal/analyzer"
	"github.com/wimspaargaren/prolayout/internal/model"
)

func main() {
	data, err := os.ReadFile(".prolayout.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	t := model.Root{}
	err = yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	singlechecker.Main(analyzer.New(t))
}
