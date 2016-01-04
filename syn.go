package main

import (
	"./syn"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	env := syn.Env{}
	for _, arg := range os.Args[1:] {
		scanner := new(syn.Scanner)
		body, err := ioutil.ReadFile(arg)
		if err != nil {
			log.Fatal(err)
		}
		scanner.Init(string(body))
		for _, statement := range syn.Parse(scanner) {
			_, err := syn.Evaluate(statement, env)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
