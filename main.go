package main

import (
	"flag"
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/caddyserver/certmagic"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var whitelistCSV, target, email string
	var agreed bool
	flag.StringVar(&whitelistCSV, "whitelist", "", "the csv of whitelist domains to get ssl certs for")
	flag.StringVar(&target, "target", "", "the target url to proxy for")
	flag.StringVar(&email, "email", "test@example.com", "the acme email to declare")
	flag.BoolVar(&agreed, "agree", false, "whether you agree to the acme terms")
	flag.Parse()
	if len(whitelistCSV) == 0 {
		return fmt.Errorf("whitelist is required")
	}
	if len(target) == 0 {
		return fmt.Errorf("target is required")
	}
	if !agreed {
		return fmt.Errorf("you must agree to the acme terms")
	}
	whitelist := strings.Split(whitelistCSV, ",")
	log.Printf("running with whitelist %q and target %q\n", whitelist, target)
	u, err := url.Parse(target)
	if err != nil {
		return err
	}
	h := httputil.NewSingleHostReverseProxy(u)
	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = email
	return certmagic.HTTPS(whitelist, h)
}
