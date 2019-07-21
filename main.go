package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/etc/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resolve := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		fields := strings.Fields(line)
		for i := 1; i < len(fields); i++ {
			resolve[fields[i]] = fields[0]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for k, v := range resolve {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
