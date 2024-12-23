package main

import (
	"encoding/json"
	"fmt"
	"os"

	v1alpha1 "github.com/project-copacetic/copacetic/pkg/types/v1alpha1"
)

type FakeParser struct{}

// parseFakeReport parses a fake report from a file
func parseFakeReport(file string) (*FakeReport, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var fake FakeReport
	if err = json.Unmarshal(data, &fake); err != nil {
		return nil, err
	}

	return &fake, nil
}

func newFakeParser() *FakeParser {
	return &FakeParser{}
}

func (k *FakeParser) parse(file string) (*v1alpha1.UpdateManifest, error) {
	// Parse the fake report
	//report, err := parseFakeReport(file)
	//if err != nil {
	//return nil, err
	//}

	// Create the standardized report
	updates := v1alpha1.UpdateManifest{
		APIVersion: v1alpha1.APIVersion,
		Metadata: v1alpha1.Metadata{
			OS: v1alpha1.OS{
				Type:    "debian",
				Version: "11",
			},
			Config: v1alpha1.Config{
				Arch: "amd64",
			},
		},
	}

	// Convert the fake report to the standardized report

	// Convert the fake report to the standardized report
	for i := range report.basic_plan_component_vulnerability_fixes {
		vulnerabilities := &report.basic_plan_component_vulnerability_fixes[i]
		    updates.Updates = append(updates.Updates, v1alpha1.UpdatePackage{
				current_component_purl: vulnerabilities.current_component_purl,
				target_component_purl: vulnerabilities.target_component_purl,
			})
		}
	return &updates, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <image report>\n", os.Args[0])
		os.Exit(1)
	}

	// Initialize the parser
	fakeParser := newFakeParser()

	// Get the image report from command line
	imageReport := os.Args[1]

	report, err := fakeParser.parse(imageReport)
	if err != nil {
		fmt.Printf("error parsing report: %v\n", err)
		os.Exit(1)
	}

	// Serialize the standardized report and print it to stdout
	reportBytes, err := json.Marshal(report)
	if err != nil {
		fmt.Printf("Error serializing report: %v\n", err)
		os.Exit(1)
	}

	os.Stdout.Write(reportBytes)
}
