package gomesh

import (
	"encoding/xml"
	"io"
)

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

func ParseSupplementalRecordSet(r io.Reader) (<-chan SupplementalRecord, chan error) {
	dec := xml.NewDecoder(r)
	drc := make(chan SupplementalRecord)
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
				if t.Name.Local == "SupplementalRecord" {
					dr := SupplementalRecord{}

					if err := dec.DecodeElement(&dr, &t); err != nil {
						errc <- err
						close(drc)
						return
					}

					drc <- dr
				}
			default:

			}
		}
	}()
	return drc, errc
}
