package mesh

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
)

// Descriptors is a var
var Descriptors []DescriptorRecord

// DescriptorRecord is a type
type DescriptorRecord struct {
	Name string   `xml:"DescriptorName>String" json:"name"`
	UI   UniqueID `xml:"DescriptorUI" json:"ui"`

	AllowableQualifiers []struct {
		UI           string `xml:"QualifierReferredTo>QualifierUI" json:"ui"`
		Name         string `xml:"QualifierReferredTo>QualifierName>String" json:"name"`
		Abbreviation string `xml:"Abbreviation" json:"abbreviation"`
	} `xml:"AllowableQualifiersList>AllowableQualifier" json:"allowable_qualifiers,omitempty"`

	TreeNumbers []TreeNumber `xml:"TreeNumberList>TreeNumber" json:"tree_numbers,omitempty"`

	Concepts []struct {
		UI        UniqueID `xml:"ConceptUI" json:"ui"`
		Prefered  YN       `xml:"PreferredConceptYN,attr" json:"prefered,omitempty"`
		Name      string   `xml:"ConceptName>String" json:"name"`
		ScopeNote string   `json:"scope_note"`
		CASN1Name string   `xml:"CASN1Name" json:"casn1_name,omitempty"`
		Terms     []struct {
			UI UniqueID `xml:"TermUI" json:"ui"`

			ConceptPreferedTerm YN `xml:"ConceptPreferredTermYN,attr" json:"prefered,omitempty"`
			Permuted            YN `xml:"IsPermutedTermYN,attr" json:"permuted,omitempty"`
			RecordPreferedTerm  YN `xml:"RecordPreferredTermYN,attr" json:"permuted,omitempty"`

			Name       string `xml:"String" json:"name"`
			LexicalTag string `xml:"LexicalTag,attr" json:"lexical_tag"`
		} `xml:"TermList>Term" json:"terms"`
	} `xml:"ConceptList>Concept" json:"concepts,omitempty"`

	PharmActions []struct {
		UI   UniqueID `xml:"DescriptorReferredTo>DescriptorUI" json:"ui"`
		Name string   `xml:"DescriptorReferredTo>DescriptorName>String" json:"name"`
	} `xml:"PharmacologicalActionList>PharmacologicalAction" json:"pharm_actions,omitempty"`
}

func (record DescriptorRecord) String() string {
	b, _ := json.MarshalIndent(record, "", "  ")
	return string(b)
}

// TreeNumber is a type
type TreeNumber string

// LoadDescriptors is a func
func LoadDescriptors(r io.ReadCloser) error {
	defer r.Close()
	descSet := struct {
		DescriptorRecords []DescriptorRecord `xml:"DescriptorRecord"`
	}{}

	dec := xml.NewDecoder(r)
	err := dec.Decode(&descSet)
	Descriptors = descSet.DescriptorRecords
	fmt.Printf("len(Descriptors): %d\n", len(Descriptors))
	return err
}
