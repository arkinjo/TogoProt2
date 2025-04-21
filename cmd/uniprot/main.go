package main

import (
	"fmt"

	"github.com/arkinjo/TogoProt2/uniprot"
)

// Example usage (in your main package):
func main() {
	for entry, err := range UniProtEntries("uniprot.xml.gz") {
		if err != nil {
			log.Fatal("Reading a UniProt entry failed: ", err)
		}
		fmt.Printf("Found entry with accession(s): %v\n", entry.Accession)
		// Process the entry here
	}
	fmt.Println("Finished processing entries.")
}
