package mesh

import "encoding/xml"

// UniqueID is a type
type UniqueID string

// Type returns a string description of a UniqueID type
// the result can be
func (id UniqueID) Type() string {
	if !(len(id) >= 1) {
		return "invalid"
	}
	switch id[0] {
	case 'D':
		return "descriptor"
	case 'Q':
		return "qualifier"
	case 'M':
		return "concept"
	case 'T':
		return "term"
	default:
		return "unknown"
	}
}

// YN helps with Y/N attributes
type YN bool

// UnmarshalXMLAttr implements xml.Unmarshaler for YN
func (yn *YN) UnmarshalXMLAttr(attr xml.Attr) error {
	(*yn) = attr.Value == "Y"
	return nil
}

// MarshalXMLAttr implements xml.Unmarshaler for YN
func (yn YN) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if yn {
		return xml.Attr{
			Name:  name,
			Value: "Y",
		}, nil
	}
	return xml.Attr{
		Name:  name,
		Value: "N",
	}, nil
}
