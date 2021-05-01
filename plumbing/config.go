package plumbing

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Common struct {
	Keys map[string]string
}

func Config() *Common {
	c := Common{}
	c.Keys = keysFromEnv()
	return &c
}

func keysFromEnv() map[string]string {
	f, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ln := scanner.Text()
		lns := strings.Split(ln, "=")
		k := lns[0]
		v := lns[1]
		m[k] = v
	}
	return m
}
