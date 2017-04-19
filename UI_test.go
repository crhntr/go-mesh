package mesh_test

import (
	"encoding/xml"
	"testing"

	mesh "github.com/crhntr/go-mesh"
)

func TestUI_Unmarshal(t *testing.T) {
	testXML := []byte(`<foo><SomeUI>C114158</SomeUI></foo>`)
	type TStruct struct {
		XMLName xml.Name `xml:"foo"`
		UI      mesh.UI  `xml:"SomeUI"`
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

func TestUI_Type(t *testing.T) {
	testTable := []struct {
		Exp mesh.UIType
		UI  string
	}{
		{mesh.UnknownUI, ""},
		{mesh.UnknownUI, "*"},
		{mesh.DescriptorUI, "D000001"},
		{mesh.DescriptorUI, "*D000001"},
		{mesh.QualifierUI, "Q000031"},
		{mesh.QualifierUI, "*Q000031"},
		{mesh.ConceptUI, "M0030903"},
		{mesh.ConceptUI, "C114158"},
		{mesh.TermUI, "T324525"},
		{mesh.UnknownUI, "324525"},
	}

	for i, testRow := range testTable {
		got := mesh.UI(testRow.UI).Type()
		if got != testRow.Exp {
			t.Errorf("%d\t ui:%s, expected: %q, got: %q", i, testRow.UI, got, testRow.Exp)
		}
	}
}
