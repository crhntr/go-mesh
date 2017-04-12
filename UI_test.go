package mesh_test

import (
	"encoding/xml"
	"testing"

	"github.com/crhntr/gomesh"
)

func TestUI(t *testing.T) {
	testXML := []byte(`<foo><SomeUI>C114158</SomeUI></foo>`)
	type TStruct struct {
		XMLName xml.Name  `xml:"foo"`
		UI      gomesh.UI `xml:"SomeUI"`
	}
	v := TStruct{}
	err := xml.Unmarshal(testXML, &v)
	if err != nil {
		t.Error(err)
	}

	if v.UI == "" {
		t.Error("could not parse UI")
	}
}
