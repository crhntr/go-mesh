package gomesh

import (
	"encoding/xml"
	"io"
)

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
