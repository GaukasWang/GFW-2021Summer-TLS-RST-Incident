package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	tls "github.com/refraction-networking/utls"
	"gopkg.in/yaml.v3"
)

// ClientHello Specs
var ClientHelloSpecsGen []func() tls.ClientHelloSpec

type Conf struct {
	Port      uint16   `yaml:"port"`
	Domain    string   `yaml:"domain"`
	Subdomain []string `yaml:"subdomain"`
}

func usage(CustomInfo string) {
	fmt.Println("Usage: verifier config.yaml ITERATION SLEEP_MS")
	panic(CustomInfo)
}

func main() {
	argc := len(os.Args)
	if argc != 4 {
		usage("Incorrect arguments.")
	}

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		usage("Can't open config file.")
	}
	myConf := &Conf{}
	err = yaml.Unmarshal(buf, myConf)
	if err != nil {
		usage("Config file doesn't match expected format.")
	}

	var iter uint = 1 // Default iteration: 1 test per hostname
	iter64, err := strconv.ParseUint(os.Args[2], 10, 32)
	if err != nil {
		usage("Incorrect arguments.")
	}
	iter = uint(iter64)

	var sleep_ms int = 5
	sleep_ms64, err := strconv.ParseUint(os.Args[3], 10, 32)
	if err != nil {
		usage("Incorrect arguments.")
	}
	sleep_ms = int(sleep_ms64)

	ClientHelloSpecsGen = append(ClientHelloSpecsGen, CH_a91c0644c199823d)
	ClientHelloSpecsGen = append(ClientHelloSpecsGen, CH_6bfedc5d5c740d58)
	ClientHelloSpecsGen = append(ClientHelloSpecsGen, CH_f6c7540db365dd4c)
	ClientHelloSpecsGen = append(ClientHelloSpecsGen, CH_8466c4390d4bc355)
	// ClientHelloSpecsGen = append(ClientHelloSpecsGen, CH_2aaf12c5eb0cb798) // This won't work :(

	for chs_i, chspecgen := range ClientHelloSpecsGen {
		fmt.Printf("=== ClientHello Spec #%d ===\n", chs_i)
		for i := uint(0); i < iter; i++ {
			for _, sub := range myConf.Subdomain {
				chspec := chspecgen()
				hostname := fmt.Sprintf("%s.%s", sub, myConf.Domain)
				addr := fmt.Sprintf("%s:%d", hostname, myConf.Port)
				start_t := time.Now()
				_, err := TestHTTPSHandshake(hostname, addr, chspec)
				elapse_ms := time.Since(start_t) / (1000 * 1000)
				if err != nil {
					fmt.Printf("%s %s %d\n", sub, err, elapse_ms)
				} else {
					fmt.Printf("%s SUCCESS %d\n", sub, elapse_ms)
				}
				time.Sleep(time.Duration(sleep_ms) * time.Millisecond)
			}
		}
	}
}
