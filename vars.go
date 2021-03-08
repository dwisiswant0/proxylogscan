package main

import (
	"bufio"
	"net/http"
	"sync"
)

var (
	baseURL  string
	method   string
	proxyURL string
	silent   bool

	target  []string
	scanner *bufio.Scanner
	client  http.Client
	wg      sync.WaitGroup
)
