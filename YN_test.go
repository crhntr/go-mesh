package mesh_test

import (
	"encoding/xml"
	"testing"

	mesh "github.com/crhntr/go-mesh"
)

func TestYN_XMLUnmarshal(t *testing.T) {
	xmlData := [][]byte{[]byte(`<foo YesYN="Y" NoYN="N"></foo>`), []byte(`<foo YesYN="N" NoYN="Y"></foo>`)}

	type Foo struct {
		XMLName xml.Name `xml:"foo"`
		YesYN   mesh.YN  `xml:"YesYN,attr"`
		NoYN    mesh.YN  `xml:"NoYN,attr"`
	}

	v0 := Foo{}
	if err := xml.Unmarshal(xmlData[0], &v0); err != nil {
		t.Error(err)
	}

	if v0.NoYN {
		t.Error("NoYN should be false")
	}

	if !v0.YesYN {
		t.Error("YesYN should be true")
	}

	v1 := Foo{}
	if err := xml.Unmarshal(xmlData[1], &v1); err != nil {
		t.Error(err)
	}

	if !v1.NoYN {
		t.Error("NoYN should be true")
	}

	if v1.YesYN {
		t.Error("YesYN should be false")
	}
}

func TestYN_XMLMarshal(t *testing.T) {
	type Foo struct {
		XMLName xml.Name `xml:"foo"`
		YesYN   mesh.YN  `xml:"YesYN,attr"`
		NoYN    mesh.YN  `xml:"NoYN,attr"`
	}

	foo := Foo{
		YesYN: true,
		NoYN:  false,
	}

	res, err := xml.Marshal(foo)
	if err != nil {
		t.Error(err)
	}

	expected := "<foo YesYN=\"Y\" NoYN=\"N\"></foo>"
	if expected != string(res) {
		t.Errorf("expected %q got %q", expected, string(res))
	}
}
