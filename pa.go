package gomesh

import (
	"encoding/xml"
	"io"
)

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
