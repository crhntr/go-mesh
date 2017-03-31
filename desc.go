package gomesh

import (
	"encoding/xml"
	"io"
)

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
