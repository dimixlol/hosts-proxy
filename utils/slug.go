package utils

import "strings"

func GetSlugFromHost(host string) string {
	//TODO : remove port
	// remove root dot if any
	if strings.HasSuffix(host, ".") {
		host = host[:len(host)-1]
	}

	subdomains := strings.Split(host, ".")
	domainNumber := len(subdomains)

	// domain.tld or domain or ""
	if domainNumber < 3 {
		panic("err durint checking the slug for host: " + host + "")
		// subdomain.domain.tld
	} else if domainNumber == 3 {
		return subdomains[0]
	}

	// subdomainN...subdomain1.subdomain0.domain.tld
	return subdomains[domainNumber-3]
}
