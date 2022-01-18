package adresses

import "strings"

// AddressType cheks if an address has a valid type and returns
func AddressType(address string) string {

	validTypes := []string{"Street", "Avenue", "Road", "highway"}

	lowercaseAddress := strings.ToLower(address)
	firstWordAddress := strings.Split(lowercaseAddress, " ")[0]

	addressHasValidType := false
	for _, typeAddress := range validTypes {
		if typeAddress == firstWordAddress {
			addressHasValidType = true
		}
	}

	if addressHasValidType {
		return firstWordAddress
	}

	return "Invalid Type"
}
