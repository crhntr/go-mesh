package gomesh

import (
	"encoding/xml"
	"io"
)

// ParseDescriptorRecordSet is designed to parse desc####.xml files
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

// ParsePharmacologicalActionSet is designed to parse pa####.xml files
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

// ParseQualifierRecordSet is designed to parse qual####.xml files
func ParseQualifierRecordSet(r io.Reader) (<-chan QualifierRecord, chan error) {
	dec := xml.NewDecoder(r)
	qrc := make(chan QualifierRecord)
	errc := make(chan error)
	go func() {
		for {
			tok, err := dec.Token()
			if err != nil {
				errc <- err
				close(qrc)
				return
			}

			switch t := tok.(type) {
			case xml.StartElement:
				if t.Name.Local == "QualifierRecord" {
					qr := QualifierRecord{}

					if err := dec.DecodeElement(&qr, &t); err != nil {
						errc <- err
						close(qrc)
						return
					}

					qrc <- qr
				}
			}
		}
	}()
	return qrc, errc
}

// ParseSupplementalRecordSet is designed to parse supp####.xml files
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
