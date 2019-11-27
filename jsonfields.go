package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	// We assume that the input JSON is a single document containing an array of objects.
	var data []map[string]interface{}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("can't read input: %v", err)
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatalf("can't unmarshal JSON: %v", err)
	}

	fields := map[string]int{}

	for _, obj := range data {

		for k := range obj {
			fields[k]++
		}
	}

	fieldNames := mapKeys(fields)

	if len(os.Args) > 1 && os.Args[1] == "-c" {
		for _, name := range fieldNames {
			fmt.Printf("%s: %d\n", name, fields[name])
		}
	} else {
		fmt.Printf("%s\n", strings.Join(fieldNames, " "))
	}
}

func mapKeys(m map[string]int) []string {

	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	return keys
}
