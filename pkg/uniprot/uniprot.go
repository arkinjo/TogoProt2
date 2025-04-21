package uniprot

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
)

// Define the structure for a single UniProt entry
type Entry struct {
	XMLName          xml.Name         `xml:"entry"`
	Dataset          string           `xml:"dataset,attr"`
	Created          string           `xml:"created,attr"`
	Modified         string           `xml:"modified,attr"`
	Version          int              `xml:"version,attr"`
	Accession        []string         `xml:"accession"`
	Name             []Name           `xml:"name"`
	Protein          Protein          `xml:"protein"`
	Gene             []Gene           `xml:"gene"`
	Organism         Organism         `xml:"organism"`
	OrganismHost     []OrganismHost   `xml:"organismHost"`
	GeneLocation     []GeneLocation   `xml:"geneLocation"`
	Reference        []Reference      `xml:"reference"`
	Comment          []Comment        `xml:"comment"`
	DbReference      []DbReference    `xml:"dbReference"`
	ProteinExistence ProteinExistence `xml:"proteinExistence"`
	Keyword          []Keyword        `xml:"keyword"`
	Feature          []Feature        `xml:"feature"`
	Sequence         Sequence         `xml:"sequence"`
}

// Define other necessary structs (Name, Protein, Gene, Organism, Sequence, etc.)
// Make sure these structs match the relevant parts of your UniProt XML.

type Name struct {
	XMLName xml.Name `xml:"name"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type Protein struct {
	XMLName         xml.Name          `xml:"protein"`
	RecommendedName RecommendedName   `xml:"recommendedName"`
	AlternativeName []AlternativeName `xml:"alternativeName"`
	SubmittedName   []SubmittedName   `xml:"submittedName"`
}

type RecommendedName struct {
	XMLName   xml.Name    `xml:"recommendedName"`
	FullName  FullName    `xml:"fullName"`
	ShortName []ShortName `xml:"shortName"`
}

type AlternativeName struct {
	XMLName   xml.Name    `xml:"alternativeName"`
	FullName  FullName    `xml:"fullName"`
	ShortName []ShortName `xml:"shortName"`
}

type SubmittedName struct {
	XMLName   xml.Name    `xml:"submittedName"`
	FullName  FullName    `xml:"fullName"`
	ShortName []ShortName `xml:"shortName"`
}

type FullName struct {
	XMLName  xml.Name   `xml:"fullName"`
	Evidence []Evidence `xml:"evidence"`
	Value    string     `xml:",chardata"`
}

type ShortName struct {
	XMLName  xml.Name   `xml:"shortName"`
	Evidence []Evidence `xml:"evidence"`
	Value    string     `xml:",chardata"`
}

type Gene struct {
	XMLName xml.Name   `xml:"gene"`
	Name    []GeneName `xml:"name"`
}

type GeneName struct {
	XMLName xml.Name `xml:"name"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type Organism struct {
	XMLName        xml.Name       `xml:"organism"`
	Name           []OrganismName `xml:"name"`
	DbReference    []DbReference  `xml:"dbReference"`
	Lineage        Lineage        `xml:"lineage"`
	Classification []string       `xml:"classification"`
}

type OrganismName struct {
	XMLName xml.Name `xml:"name"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type DbReference struct {
	XMLName  xml.Name   `xml:"dbReference"`
	Type     string     `xml:"type,attr"`
	ID       string     `xml:"id,attr"`
	Evidence []Evidence `xml:"evidence"`
}

type Lineage struct {
	XMLName xml.Name `xml:"lineage"`
	Taxon   []Taxon  `xml:"taxon"`
}

type Taxon struct {
	XMLName  xml.Name   `xml:"taxon"`
	Evidence []Evidence `xml:"evidence"`
	Value    string     `xml:",chardata"`
}

type Sequence struct {
	XMLName  xml.Name `xml:"sequence"`
	Length   int      `xml:"length,attr"`
	Mass     int      `xml:"mass,attr"`
	Version  int      `xml:"version,attr"`
	Modified string   `xml:"modified,attr"`
	Checksum string   `xml:"checksum,attr"`
	Value    string   `xml:",chardata"`
}

type Feature struct {
	XMLName     xml.Name    `xml:"feature"`
	Type        string      `xml:"type,attr"`
	Id          string      `xml:"id,attr"`
	Description string      `xml:"description,attr"`
	Evidence    []Evidence  `xml:"evidence"`
	Location    Location    `xml:"location"`
	Ref         string      `xml:"ref,attr"`
	Original    string      `xml:"original"`
	Variation   []Variation `xml:"variation"`
}

type Location struct {
	XMLName  xml.Name `xml:"location"`
	Position Position `xml:"position"`
	Begin    Begin    `xml:"begin"`
	End      End      `xml:"end"`
}

type Position struct {
	XMLName xml.Name `xml:"position"`
	Status  string   `xml:"status,attr"`
	Value   int      `xml:",chardata"`
}

type Begin struct {
	XMLName  xml.Name `xml:"begin"`
	Status   string   `xml:"status,attr"`
	Position int      `xml:",chardata"`
}

type End struct {
	XMLName  xml.Name `xml:"end"`
	Status   string   `xml:"status,attr"`
	Position int      `xml:",chardata"`
}

type Variation struct {
	XMLName  xml.Name `xml:"variation"`
	Original string   `xml:"original"`
	Sequence string   `xml:",chardata"`
}

type Evidence struct {
	XMLName xml.Name `xml:"evidence"`
	Type    string   `xml:"type,attr"`
	Key     string   `xml:"key,attr"`
}

type OrganismHost struct {
	XMLName     xml.Name       `xml:"organismHost"`
	Name        []OrganismName `xml:"name"`
	DbReference []DbReference  `xml:"dbReference"`
	Lineage     Lineage        `xml:"lineage"`
}

type GeneLocation struct {
	XMLName     xml.Name         `xml:"geneLocation"`
	Gene        string           `xml:"gene,attr"`
	Evidence    []Evidence       `xml:"evidence"`
	Name        GeneLocationName `xml:"name"`
	Chromosome  string           `xml:"chromosome"`
	MapPosition string           `xml:"mapPosition"`
}

type GeneLocationName struct {
	XMLName xml.Name `xml:"name"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type Reference struct {
	XMLName     xml.Name        `xml:"reference"`
	Key         string          `xml:"key,attr"`
	Citation    Citation        `xml:"citation"`
	Scope       []string        `xml:"scope"`
	Source      Source          `xml:"source"`
	Protein     ProteinSection  `xml:"protein"`
	Gene        GeneSection     `xml:"gene"`
	Organism    OrganismSection `xml:"organism"`
	DbReference []DbReference   `xml:"dbReference"`
}

type Citation struct {
	XMLName     xml.Name      `xml:"citation"`
	Type        string        `xml:"type,attr"`
	Date        string        `xml:"date"`
	Title       string        `xml:"title"`
	Journal     Journal       `xml:"journal"`
	AuthorList  AuthorList    `xml:"authorList"`
	DbReference []DbReference `xml:"dbReference"`
}

type Journal struct {
	XMLName xml.Name `xml:"journal"`
	Value   string   `xml:",chardata"`
}

type AuthorList struct {
	XMLName xml.Name `xml:"authorList"`
	Person  []Person `xml:"person"`
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name,attr"`
}

type Source struct {
	XMLName     xml.Name      `xml:"source"`
	Organism    Organism      `xml:"organism"`
	DbReference []DbReference `xml:"dbReference"`
	Strain      []string      `xml:"strain"`
}

type ProteinSection struct {
	XMLName xml.Name `xml:"protein"`
	Name    []Name   `xml:"name"`
}

type GeneSection struct {
	XMLName xml.Name   `xml:"gene"`
	Name    []GeneName `xml:"name"`
}

type OrganismSection struct {
	XMLName xml.Name       `xml:"organism"`
	Name    []OrganismName `xml:"name"`
}

type Comment struct {
	XMLName           xml.Name          `xml:"comment"`
	Type              string            `xml:"type,attr"`
	Evidence          []Evidence        `xml:"evidence"`
	Text              []Text            `xml:"text"`
	Molecule          string            `xml:"molecule,attr"`
	Location          Location          `xml:"location"`
	Reaction          Reaction          `xml:"reaction"`
	Enzyme            Enzyme            `xml:"enzyme"`
	Ph                Ph                `xml:"ph"`
	Temperature       Temperature       `xml:"temperature"`
	KineticParameters KineticParameters `xml:"kineticParameters"`
}

type Text struct {
	XMLName  xml.Name   `xml:"text"`
	Evidence []Evidence `xml:"evidence"`
	Value    string     `xml:",chardata"`
}

type Reaction struct {
	XMLName     xml.Name      `xml:"reaction"`
	Name        []string      `xml:"name"`
	DbReference []DbReference `xml:"dbReference"`
	EC          string        `xml:"ec"`
}

type Enzyme struct {
	XMLName xml.Name `xml:"enzyme"`
	EC      []string `xml:"ec"`
}

type Ph struct {
	XMLName xml.Name `xml:"ph"`
	Value   string   `xml:",chardata"`
}

type Temperature struct {
	XMLName xml.Name `xml:"temperature"`
	Value   string   `xml:",chardata"`
}

type KineticParameters struct {
	XMLName xml.Name `xml:"kineticParameters"`
	Km      []Km     `xml:"km"`
	Vmax    []Vmax   `xml:"vmax"`
}

type Km struct {
	XMLName xml.Name `xml:"km"`
	Value   string   `xml:",chardata"`
	Unit    string   `xml:"unit,attr"`
}

type Vmax struct {
	XMLName xml.Name `xml:"vmax"`
	Value   string   `xml:",chardata"`
	Unit    string   `xml:"unit,attr"`
}

type ProteinExistence struct {
	XMLName xml.Name `xml:"proteinExistence"`
	Type    string   `xml:"type,attr"`
}

type Keyword struct {
	XMLName  xml.Name   `xml:"keyword"`
	Evidence []Evidence `xml:"evidence"`
	Value    string     `xml:",chardata"`
}

// UniProtEntries returns an iterator over UniProt entries from a gzipped XML file.
func UniProtEntries(filePath string) iter.Seq2[Entry, error] {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(filePath)
	}

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		file.Close()
		log.Fatal("gzip error")
	}

	decoder := xml.NewDecoder(gzipReader)
	yieldedRoot := false

	return func(yield func(Entry, error) bool) {
		defer file.Close()
		defer gzipReader.Close()
		for {
			token, err := decoder.Token()
			if err != nil {
				if err == io.EOF {
					return
				}
				log.Fatalf("Error decoding XML: %v\n", err)
				return
			}

			if start, ok := token.(xml.StartElement); ok {
				if start.Name.Local == "uniprot" {
					yieldedRoot = true
					continue // Move to the next token
				}
				if start.Name.Local == "entry" && yieldedRoot {
					var entry Entry
					err := decoder.DecodeElement(&entry, &start)
					if !yield(entry, err) {
						return // Stop if the consumer doesn't want more
					}
				}
			}
		}
	}
}
