package gomesh

import (
	"encoding/xml"
	"io"
)

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
