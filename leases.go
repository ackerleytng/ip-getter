package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// Takes the content of a lease block, and returns a pair.
// The second item in the pair determines whether this is a valid lease.
// If the lease is valid, the first item in the pair will be the mac address
//   associated with this lease.
func parseLeaseContent(content []string) (string, bool) {
	// We don't care about the expiry time (starts/ends) because
	//   the dhcp server has a grace period
	macAddressRegex := regexp.MustCompile(`((?:[0-9A-Fa-f]{2}[:-]){5}(?:[0-9A-Fa-f]{2}))`)
	var macAddress = ""

	// We will consider the lease invalid only if
	//   + The lease is marked abandoned
	//   + binding state is marked anything other than "active"
	for _, s := range content {
		if strings.HasPrefix(s, "abandoned") ||
			(strings.HasPrefix(s, "binding state") && !strings.Contains(s, "active")) {
			return "", false
		}

		if strings.HasPrefix(s, "hardware") {
			result := macAddressRegex.FindStringSubmatch(s)
			if len(result) > 1 {
				macAddress = result[1]
			}
		}
	}

	return macAddress, len(macAddress) > 0
}

func GetLeases(leasesFile string) map[string]string {
	var ipToMac = make(map[string]string)

	file, err := os.Open(leasesFile)
	if err == nil {
		defer file.Close()

		ipv4Regex := regexp.MustCompile(`((?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})`)

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)

		var leaseContent []string
		var inLease = false
		var ipAddress = ""

		for scanner.Scan() {
			line := scanner.Text()

			if inLease {
				if strings.HasPrefix(line, "}") {
					inLease = false
					macAddress, valid := parseLeaseContent(leaseContent)
					if valid {
						ipToMac[ipAddress] = macAddress
					}
					leaseContent = leaseContent[:0]
				} else {
					leaseContent = append(leaseContent, strings.Trim(line, " \t"))
				}
			} else if strings.HasPrefix(line, "lease") {
				result := ipv4Regex.FindStringSubmatch(line)
				if len(result) > 1 {
					ipAddress = result[1]
				}
				inLease = true
			}

		}

		// check for errors
		err = scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}

	return ipToMac
}
