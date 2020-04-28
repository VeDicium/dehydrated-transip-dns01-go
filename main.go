package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/domain"
	"github.com/transip/gotransip/v6/product"
	"github.com/transip/gotransip/v6/repository"

	"github.com/bobesa/go-domain-util/domainutil"
)

var TransIPClient repository.Client
var DNSName string

func main() {
	// Set default DNS name
	DNSName = "_acme-challenge"

	// Create a new TransIP API client
	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:    os.Getenv("TRANSIP_ACCOUNT_NAME"),
		PrivateKeyPath: os.Getenv("TRANSIP_KEY_PATH"),
		TestMode:       os.Getenv("TEST_MODE") == "1",
		Mode:           gotransip.APIModeReadWrite,
	})
	if err != nil {
		log.Fatal(err)
	}
	TransIPClient = client

	// Check handler (command)
	handler := ""
	if len(os.Args) >= 2 {
		handler = os.Args[1]
	}

	// Test
	if handler == "test" {
		productRepo := product.Repository{Client: TransIPClient}
		_, err := productRepo.GetAll()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Successfully tested API credentials")

		return
	}

	// Deploy challenge
	if handler == "deploy_challenge" {
		// Execute the deploy challenge
		err := deployChallenge()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Sleep for 30 seconds")
		time.Sleep(30 * time.Second)

		return
	}

	// Clean challenge
	if handler == "clean_challenge" {
		// Execute the clean challenge
		err := cleanChallenge()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Sleep for 30 seconds")
		time.Sleep(30 * time.Second)

		return
	}
}

// Challenges
func deployChallenge() error {
	// Get domain name of challenge
	if len(os.Args) < 3 {
		return errors.New("No domain name given")
	}
	domainName := os.Args[2]
	topLevelDomain := domainutil.Domain(domainName)

	// Create domain repo
	domainRepo := domain.Repository{Client: TransIPClient}

	// Check if domain exists
	log.Printf("Check for domain '%s'\n", domainName)
	_, err := domainRepo.GetByDomainName(topLevelDomain)
	if err != nil {
		return err
	}

	// Get token
	if len(os.Args) < 5 {
		return errors.New("No token given")
	}
	token := os.Args[4]

	// Add DNS
	dnsEntry := domain.DNSEntry{
		Name:    DNSName,
		Expire:  60,
		Type:    "TXT",
		Content: token,
	}

	// Set name when subdomain is found
	subDomain := domainutil.Subdomain(domainName)
	if len(subDomain) > 0 {
		dnsEntry.Name = DNSName + "." + subDomain
	}

	// Set DNS record
	log.Printf("Adding DNS Record: %s - %s - %s (TTL %d) for domain '%s'...\n", dnsEntry.Name, dnsEntry.Type, dnsEntry.Content, dnsEntry.Expire, topLevelDomain)
	err = domainRepo.AddDNSEntry(topLevelDomain, dnsEntry)
	if err != nil {
		return err
	}
	log.Printf("DNS record added")

	return nil
}

func cleanChallenge() error {
	// Get domain name of challenge
	if len(os.Args) < 3 {
		return errors.New("No domain name given")
	}
	domainName := os.Args[2]
	topLevelDomain := domainutil.Domain(domainName)

	// Create domain repo
	domainRepo := domain.Repository{Client: TransIPClient}

	// Check if domain exists
	log.Printf("Check for domain '%s'\n", domainName)
	_, err := domainRepo.GetByDomainName(topLevelDomain)
	if err != nil {
		return err
	}

	// Get token
	if len(os.Args) < 5 {
		return errors.New("No token given")
	}
	token := os.Args[4]

	// Add DNS
	dnsEntry := domain.DNSEntry{
		Name:    DNSName,
		Expire:  60,
		Type:    "TXT",
		Content: token,
	}

	// Set name when subdomain is found
	subDomain := domainutil.Subdomain(domainName)
	if len(subDomain) > 0 {
		dnsEntry.Name = DNSName + "." + subDomain
	}

	// Set DNS record
	log.Printf("Removing DNS Record: %s - %s - %s (TTL %d) for domain '%s'...\n", dnsEntry.Name, dnsEntry.Type, dnsEntry.Content, dnsEntry.Expire, topLevelDomain)
	err = domainRepo.RemoveDNSEntry(topLevelDomain, dnsEntry)
	if err != nil {
		return err
	}
	log.Printf("DNS record removed")

	return nil
}
