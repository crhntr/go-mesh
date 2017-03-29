package gomesh

import "encoding/xml"

type YN bool

func (pf *YN) UnmarshalXMLAttr(attr xml.Attr) error {
	*pf = attr.Value == "Y"
	return nil
}

func (pf YN) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	val := "N"
	if pf {
		val = "Y"
	}
	return xml.Attr{
		Name:  name,
		Value: val,
	}, nil
}
