package mesh

import (
	"encoding/xml"
	"io"
)

// ScanDescriptorRecordSet is designed to parse desc####.xml files
func ScanDescriptorRecordSet(r io.Reader, f func(dr *DescriptorRecord) error) error {
	nm := &xml.StartElement{Name: xml.Name{Local: "DescriptorRecord"}}
	return scanXMLElement(r, nm.Name.Local, func(dec *xml.Decoder) error {
		dr := &DescriptorRecord{}
		if err := dec.DecodeElement(dr, nm); err != nil {
			return err
		}
		return f(dr)
	})
}

// ScanPharmacologicalActionSet is designed to parse pa####.xml files
func ScanPharmacologicalActionSet(r io.Reader, f func(dr *PharmacologicalAction) error) error {
	nm := &xml.StartElement{Name: xml.Name{Local: "PharmacologicalAction"}}
	return scanXMLElement(r, nm.Name.Local, func(dec *xml.Decoder) error {
		dr := &PharmacologicalAction{}
		if err := dec.DecodeElement(dr, nm); err != nil {
			return err
		}
		return f(dr)
	})
}

// ScanQualifierRecordSet is designed to parse qual####.xml files
func ScanQualifierRecordSet(r io.Reader, f func(dr *QualifierRecord) error) error {
	nm := &xml.StartElement{Name: xml.Name{Local: "QualifierRecord"}}
	return scanXMLElement(r, nm.Name.Local, func(dec *xml.Decoder) error {
		dr := &QualifierRecord{}
		if err := dec.DecodeElement(dr, nm); err != nil {
			return err
		}
		return f(dr)
	})
}

// ScanSupplementalRecordSet is designed to parse supp####.xml files
func ScanSupplementalRecordSet(r io.Reader, f func(dr *SupplementalRecord) error) error {
	nm := &xml.StartElement{Name: xml.Name{Local: "SupplementalRecord"}}
	return scanXMLElement(r, nm.Name.Local, func(dec *xml.Decoder) error {
		dr := &SupplementalRecord{}
		if err := dec.DecodeElement(dr, nm); err != nil {
			return err
		}
		return f(dr)
	})
}

func scanXMLElement(r io.Reader, elemName string, f func(dec *xml.Decoder) error) error {
	dec := xml.NewDecoder(r)
	for {
		tok, err := dec.Token()
		if err != nil {
			return err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == elemName {
				if err := f(dec); err != nil {
					return err
				}
			}
		}
	}
}
