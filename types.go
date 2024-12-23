// Type definitions for fake scanner report
package main

// FakeReport contains OS, Arch, and Package information
type LineajeReport struct {
	meta_data LineajeVulnerability
}

type LineajeVulnerability struct {
	basic_plan_component_vulnerability_fixes []Vulnerability
}

// FakePackage contains package and vulnerability information
type Vulnerability struct {
	current_component_purl string
	target_component_purl  string
}
