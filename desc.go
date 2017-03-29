package gomesh

import (
	"encoding/xml"
	"io"
)

type DescriptorRecordSet struct {
	XMLName xml.Name `xml:"DescriptorRecordSet"`
}

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
}

type Concept struct {
	PreferredConceptYN     YN
	UI                     string
	Name                   string `xml:"ConceptName>String"`
	CASN1Name              string
	RegistryNumber         string
	ScopeNote              string
	RelatedRegistryNumbers []string `xml:"RelatedRegistryNumberList>RelatedRegistryNumber"`
}

func ParseDescriptorRecordSet(r io.Reader) (<-chan DescriptorRecord, chan error) {
	dec := xml.NewDecoder(r)
	drc := make(chan DescriptorRecord)
	errc := make(chan error, 2)
	go func() {
		for {
			tok, err := dec.Token()
			if err != nil {
				errc <- err
				close(drc)
				return
			}

			switch t := tok.(type) {
			case xml.StartElement:
				if t.Name.Local == "DescriptorRecord" {
					dr := DescriptorRecord{}

					if err := dec.DecodeElement(&dr, &t); err != nil {
						errc <- err
						close(drc)
						return
					}

					drc <- dr
				}
			}
		}
	}()
	return drc, errc
}
