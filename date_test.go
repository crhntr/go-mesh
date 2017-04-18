package mesh_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	mesh "github.com/crhntr/go-mesh"
)

func TestDate_UnmarshalXML(t *testing.T) {
	testXML := []byte(`<foo><Name>Here and Now</Name><SomeDate><Year>1999</Year><Month>07</Month><Day>01</Day></SomeDate></foo>`)
	type TStruct struct {
		XMLName  xml.Name `xml:"foo"`
		Name     string
		SomeDate mesh.Date
	}
	v := TStruct{}
	err := xml.Unmarshal(testXML, &v)
	if err != nil {
		t.Error(err)
	}

	if v.Name != "Here and Now" {
		t.Error("could not parse name")
	}

	if v.SomeDate.Time.Year() != 1999 {
		t.Errorf("did not parse Year properly: expected %d, got: %d", 1999, v.SomeDate.Time.Year())
	}

	if v.SomeDate.Time.Month() != 7 {
		t.Errorf("did not parse Month properly: expected %d, got: %d", 7, v.SomeDate.Time.Month())
	}

	if v.SomeDate.Time.Day() != 1 {
		t.Errorf("did not parse Day properly: expected %d, got: %d", 1, v.SomeDate.Time.Day())
	}
}

func TestDate_MarshalXML(t *testing.T) {
	testXML := []byte(`<foo><SomeDate><Year>1999</Year><Month>07</Month><Day>01</Day></SomeDate></foo>`)
	type Foo struct {
		XMLName  xml.Name `xml:"foo"`
		SomeDate mesh.Date
	}
	v := Foo{}
	if err := xml.Unmarshal(testXML, &v); err != nil {
		t.Error(err)
	}

	resultXML, err := xml.Marshal(v)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(resultXML, testXML) {
		t.Error("xml results should be equal")
		t.Logf("Expected: %q", string(testXML))
		t.Logf("Got: %q", string(resultXML))
	}
}
