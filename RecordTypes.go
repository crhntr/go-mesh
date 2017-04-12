package mesh

// SupplementalRecords are parsed from supp####.xml files
type SupplementalRecord struct {
	UI        UI     `xml:"SupplementalRecordUI"`
	Name      string `xml:"SupplementalRecordName>String"`
	Created   Date   `xml:"DateCreated"`
	Revised   Date   `xml:"DateRevised"`
	Note      string `xml:"Note"`
	Frequency int
	Concepts  []Concept `xml:"ConceptList>Concept"`
	Sources   []string  `xml:"SourceList>Source"`
	MappedTo  []struct {
		UI   UI     `xml:"DescriptorReferredTo>DescriptorUI"`
		Name string `xml:"DescriptorReferredTo>DescriptorName"`
	} `xml:"HeadingMappedToList"`
}

// QualifierRecords are parsed from qual####.xml files
type QualifierRecord struct {
	UI          UI     `xml:"QualifierUI"`
	Name        string `xml:"QualifierName>String"`
	Created     Date   `xml:"DateCreated"`
	Revised     Date   `xml:"DateRevised"`
	Established Date   `xml:"DateEstablished"`
	Annotation  string
	HistoryNote string
	TreeNumbers []string `xml:"TreeNumberList>TreeNumber"`
}

// DescriptorRecord are parsed from desc####.xml files
type DescriptorRecord struct {
	UI                string `xml:"DescriptorUI"`
	Name              string `xml:"DescriptorName"`
	Created           Date   `xml:"DateCreated"`
	Revised           Date   `xml:"DateRevised"`
	Established       Date   `xml:"DateEstablished"`
	HistoryNote       string
	OnlineNote        string
	PublicMeSHNote    string
	PreviousIndexings []string  `xml:"PreviousIndexingList>PreviousIndexing"`
	TreeNumbers       []string  `xml:"TreeNumberList>TreeNumber"`
	Concepts          []Concept `xml:"ConceptList>Concept"`
	Terms             []Term    `xml:"TermList"`
}

// PharmacologicalAction are parsed from pa####.xml files
type PharmacologicalAction struct {
	UI   UI     `xml:"DescriptorReferredTo>DescriptorUI"`
	Name string `xml:"DescriptorReferredTo>DescriptorName>String"`

	Substances []Substance `xml:"PharmacologicalActionSubstanceList>Substance"`
}

// Substances are emebeded structs in PharmacologicalAction
type Substance struct {
	UI   UI     `xml:"RecordUI"`
	Name string `xml:"RecordName>String"`
}
