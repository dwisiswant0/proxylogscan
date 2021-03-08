package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora/v3"
	log "github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

func init() {
	flag.StringVar(&baseURL, "u", "", "Target URL/list to scan")
	flag.StringVar(&method, "m", "GET", "Request method")
	flag.StringVar(&proxyURL, "p", "", "Proxy URL (HTTP/SOCKSv5)")
	flag.BoolVar(&silent, "s", false, "Silent mode (Only display vulnerable/suppress errors)")
	flag.Parse()

	if silent {
		log.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	}
}

func main() {
	if strings.HasPrefix(baseURL, "http") {
		if !isURL(baseURL) {
			log.Fatal().Msg("invalid URL: " + baseURL)
		}
		target = []string{baseURL}
	} else {
		list, err := readList(baseURL)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		target = list
	}

	for _, URL := range target {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			req, err := request(u, method, proxyURL)
			if err != nil {
				log.Error().Msg(err.Error())
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Error().Msgf("%s\n", err.Error())
				return
			}
			defer resp.Body.Close()

			if isVuln(resp) {
				fmt.Printf("[%s] %s\n", aurora.Green("VLN"), u)
			} else {
				log.Error().Msg(u)
			}
		}(URL)
	}

	wg.Wait()
}
