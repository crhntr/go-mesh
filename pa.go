package gomesh

import (
	"encoding/xml"
	"io"
)

type PharmacologicalAction struct {
	UI   UI     `xml:"DescriptorReferredTo>DescriptorUI"`
	Name string `xml:"DescriptorReferredTo>DescriptorName>String"`

	Substances []Substance `xml:"PharmacologicalActionSubstanceList>Substance"`
}

type Substance struct {
	UI   UI     `xml:"RecordUI"`
	Name string `xml:"RecordName>String"`
}

func ParsePharmacologicalActionSet(r io.Reader) (<-chan PharmacologicalAction, chan error) {
	dec := xml.NewDecoder(r)
	drc := make(chan PharmacologicalAction)
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
				if t.Name.Local == "PharmacologicalAction" {
					dr := PharmacologicalAction{}

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
