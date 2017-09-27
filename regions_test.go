package main

import (
	"strings"
	"testing"
)

func TestNormalizeRegion(t *testing.T) {
	if normalizeRegionName("South Central US") != "southcentralus" {
		t.Error("Failed to normalize 'South Central US'")
	}

	if normalizeRegionName("southcentralus") != "southcentralus" {
		t.Error("Failed to normalize 'southcentralus'")
	}
}

func TestRegionIsExact(t *testing.T) {
	testRegions := []string{
		"South Central US",
	}

	regions := normalizeRegionList(testRegions)
	if len(regions) != 1 {
		t.Fatalf("There should be exactly one region.")
	}

	if strings.Compare("South Central US", regions[0]) != 0 {
		t.Fatalf("Expected the region to by \"South Central US\", but got %q", regions[0])
	}
}

func TestRegionMustNormalize(t *testing.T) {
	testRegions := []string{
		"southcentralus",
		"NORTH CENTRAL US",
		"uksouth",
	}

	regions := normalizeRegionList(testRegions)
	if len(regions) != 3 {
		t.Fatalf("There should be exactly three regions.")
	}

	if strings.Compare("South Central US", regions[0]) != 0 {
		t.Fatalf("Expected the region to by \"South Central US\", but got %q", regions[0])
	}
	if strings.Compare("North Central US", regions[1]) != 0 {
		t.Fatalf("Expected the region to by \"North Central US\", but got %q", regions[1])
	}
	if strings.Compare("UK South", regions[2]) != 0 {
		t.Fatalf("Expected the region to by \"UK South\", but got %q", regions[2])
	}
}

func TestUnrecognizedRegionIsPassed(t *testing.T) {
	testRegions := []string{
		"Candy Land East",
	}

	regions := normalizeRegionList(testRegions)
	if len(regions) != 1 {
		t.Fatalf("There should be exactly one region.")
	}

	if strings.Compare("Candy Land East", regions[0]) != 0 {
		t.Fatalf("Expected the region to by \"Candy Land East\", but got %q", regions[0])
	}
}
