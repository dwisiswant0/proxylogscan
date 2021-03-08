package main

import (
	"bufio"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

func isURL(s string) bool {
	_, e := url.ParseRequestURI(s)
	if e != nil {
		return false
	}

	u, e := url.Parse(s)
	if e != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func isVuln(r *http.Response) bool {
	if r.StatusCode != http.StatusInternalServerError {
		return false
	}

	for key, header := range r.Header {
		if key == "X-Calculatedbetarget" {
			for _, h := range header {
				if strings.Contains(h, "localhost") {
					return true
				}
			}
		}
	}
	return false
}

func readList(p string) ([]string, error) {
	keys := make(map[string]bool)

	if isStdin() {
		scanner = bufio.NewScanner(os.Stdin)
	} else if p != "" {
		file, err := os.Open(p)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	} else {
		return target, errors.New("no target provided")
	}

	for scanner.Scan() {
		URL := scanner.Text()
		if _, value := keys[URL]; !value {
			if isURL(URL) {
				keys[URL] = true
				target = append(target, URL)
			}
		}
	}

	if len(target) < 1 {
		return target, errors.New("no valid target URLs")
	}

	return target, scanner.Err()
}
