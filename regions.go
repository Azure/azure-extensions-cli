package main

import "strings"

var (
	// Region names differ between Service Management and Resource Manager.
	// This is an opportunity for customers to make mistakes, and not know
	// how to fix them.  This is bad. This is a brute force attempt to
	// reconcile the possible mistakes customer could make and fix them.
	//
	// The customer's input is always normalized by
	//
	//  1. region.ToLower()
	//  2. region.Replace(' ', '')
	//
	// e.g. South Central US -> southcentralus
	//
	// The value southcentralus is then looked up in a dictionary, and the
	// appropriate Service Management region is returned (South Central US).
	// If the customer's region does not exist in the dictionary, then the
	// customer's region is used.  The number of Azure regions is growing
	// over time, so this code may become out of date with reality.
	regionMap = map[string]string{
		// Simplest translation (remove whitespace, and lower case)
		//  - This is sufficient to map between SM and RM.
		"australiaeast":      "Australia East",
		"australiasoutheast": "Australia Southeast",
		"brazilsouth":        "Brazil South",
		"canadacentral":      "Canada Central",
		"canadaeast":         "Canada East",
		"centralindia":       "Central India",
		"centralus":          "Central US",
		"centraluseuap":      "Central US EUAP",
		"eastasia":           "East Asia",
		"eastus":             "East US",
		"eastus2":            "East US 2",
		"eastus2euap":        "East US 2 EUAP",
		"japaneast":          "Japan East",
		"japanwest":          "Japan West",
		"koreacentral":       "Korea Central",
		"koreasouth":         "Korea South",
		"northcentralus":     "North Central US",
		"northeurope":        "North Europe",
		"southcentralus":     "South Central US",
		"southeastasia":      "Southeast Asia",
		"southindia":         "South India",
		"uknorth":            "UK North",
		"uksouth":            "UK South",
		"uksouth2":           "UK South 2",
		"ukwest":             "UK West",
		"westcentralus":      "West Central US",
		"westeurope":         "West Europe",
		"westindia":          "West India",
		"westus":             "West US",
		"westus2":            "West US 2",

		// Swap direction and country, e.g. UK South -> South UK
		"asiaeast":           "East Asia",
		"asiasoutheast":      "Southeast Asia",
		"centralcanada":      "Canada Central",
		"centralkorea":       "Korea Central",
		"eastaustralia":      "Australia East",
		"eastcanada":         "Canada East",
		"eastjapan":          "Japan East",
		"indiacentral":       "Central India",
		"indiasouth":         "South India",
		"indiawest":          "West India",
		"northuk":            "UK North",
		"southbrazil":        "Brazil South",
		"southeastaustralia": "Australia Southeast",
		"southkorea":         "Korea South",
		"southuk":            "UK South",
		"southuk2":           "UK South 2",
		"westjapan":          "Japan West",
		"westuk":             "UK West",

		// Weirdness where direction/country is already reversed, so reverse
		// the other way.
		"europenorth ":   "North Europe",
		"europewest":     "West Europe",
		"uscentral":      "Central US",
		"uscentraleuap":  "Central US EUAP",
		"useast":         "East US",
		"useast2":        "East US 2",
		"useast2euap":    "East US 2 EUAP",
		"usnorthcentral": "North Central US",
		"ussouthcentral": "South Central US",
		"uswest":         "West US",
		"uswest2":        "West US 2",
		"uswestcentral":  "West Central US",
	}
)

func normalizeRegionList(regions []string) []string {
	normalizedRegions := make([]string, len(regions))
	for i := range regions {
		n := normalizeRegionName(regions[i])
		if x, ok := regionMap[n]; ok {
			normalizedRegions[i] = x
		} else {
			normalizedRegions[i] = regions[i]
		}
	}

	return normalizedRegions
}

func normalizeRegionName(region string) string {
	lowered := strings.ToLower(region)
	return strings.Replace(lowered, " ", "", -1)
}
