package mesh_test

import (
	"io"
	"os"
	"testing"

	mesh "github.com/crhntr/go-mesh"
)

func TestParseDescriptorRecordSet(t *testing.T) {
	f, err := os.Open("testdata/desc2017.xml")
	if err != nil {
		t.Error(err)
	}

	drc, errc := mesh.ParseDescriptorRecordSet(f)

	count := 0
	for {
		if (testing.Short() && count > 100) ||
			(testing.Verbose() && count > 1000) ||
			(!testing.Verbose() && !testing.Short() && count > 200) {
			return
		}

		select {
		case dr := <-drc:
			t.Log(dr)
			count++
		case err := <-errc:
			if err == io.EOF {
				close(errc)
				break
			} else {
				t.Error(err)
				close(errc)
			}
		}
	}
}

func TestParsePharmacologicalActionSet(t *testing.T) {
	f, err := os.Open("testdata/pa2017.xml")
	if err != nil {
		t.Error(err)
	}

	pac, errc := mesh.ParsePharmacologicalActionSet(f)

	count := 0
	for {
		if (testing.Short() && count > 100) ||
			(testing.Verbose() && count > 1000) ||
			(!testing.Verbose() && !testing.Short() && count > 200) {
			return
		}

		select {
		case <-pac:
			count++
		case err := <-errc:
			t.Logf("count: %d", count)
			if err == io.EOF {
				close(errc)
				return
			}
			t.Error(err)
			close(errc)
			return
		}
	}
}

func TestSupplementalRecordSet(t *testing.T) {
	f, err := os.Open("testdata/supp2017.xml")
	if err != nil {
		t.Fatal(err)
	}

	var (
		UI, Name, Created, Revised, Note, Frequency, Concepts, Sources, MappedTo bool
	)

	src, errc := mesh.ParseSupplementalRecordSet(f)
	count := 0
	for {
		if (testing.Short() && count > 100) ||
			(testing.Verbose() && count > 1000) ||
			(!testing.Verbose() && !testing.Short() && count > 200) {
			return
		}

		select {
		case sr := <-src:
			count++
			if len(sr.UI) > 0 {
				UI = true
			}
			if len(sr.Name) > 0 {
				Name = true
			}
			if !sr.Created.IsZero() {
				Created = true
			}
			if !sr.Revised.IsZero() {
				Revised = true
			}
			if len(sr.Note) > 0 {
				Note = true
			}
			if sr.Frequency > 0 {
				Frequency = true
			}
			if len(sr.Concepts) > 0 {
				Concepts = true
			}
			if len(sr.Sources) > 0 {
				Sources = true
			}
			if len(sr.MappedTo) > 0 {
				MappedTo = true
			}
			if UI && Name && Created && Revised && Note && Frequency && Concepts && Sources && MappedTo {
				return
			}
			if count > 100 {
				if !UI {
					t.Error("UI does not have at least one none zero value")
				}
				if !Name {
					t.Error("Name does not have at least one none zero value")
				}
				if !Created {
					t.Error("Created does not have at least one none zero value")
				}
				if !Revised {
					t.Error("Revised does not have at least one none zero value")
				}
				if !Note {
					t.Error("Note does not have at least one none zero value")
				}
				if !Frequency {
					t.Error("Frequency does not have at least one none zero value")
				}
				if !Concepts {
					t.Error("Concepts does not have at least one none zero value")
				}
				if !Sources {
					t.Error("Sources does not have at least one none zero value")
				}
				if !MappedTo {
					t.Error("MappedTo does not have at least one none zero value")
				}
				return
			}
		case err := <-errc:
			if err == io.EOF {
				close(errc)
				return
			}
			t.Error(err)
			close(errc)
			return
		}
	}
}

func TestParseQualifierRecordSet(t *testing.T) {
	f, err := os.Open("testdata/qual2017.xml")
	if err != nil {
		t.Error(err)
	}

	qrc, errc := mesh.ParseQualifierRecordSet(f)

	count := 0
	for {
		if (testing.Short() && count > 100) ||
			(testing.Verbose() && count > 1000) ||
			(!testing.Verbose() && !testing.Short() && count > 200) {
			return
		}

		select {
		case <-qrc:
			count++
		case err := <-errc:
			t.Logf("count: %d", count)
			if err == io.EOF {
				close(errc)
				return
			}
			t.Error(err)
			close(errc)
			return
		}
	}
}
