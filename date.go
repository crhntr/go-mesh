package gomesh

import (
	"encoding/xml"
	"time"
)

type Date struct {
	time.Time
}

func (date *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type dateXML struct {
		Year  int
		Month int
		Day   int
	}
	var (
		err  error
		dxml dateXML
	)

	err = d.DecodeElement(&dxml, &start)
	if err != nil {
		return err
	}

	*date = Date{time.Date(dxml.Year, time.Month(dxml.Month), dxml.Day, 0, 0, 0, 0, time.UTC)}
	return nil
}

func (date Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type dateXML struct {
		Year  int
		Month int
		Day   int
	}
	return e.EncodeElement(&dateXML{
		Year:  date.Year(),
		Month: int(date.Month()),
		Day:   date.Day(),
	}, start)
}
