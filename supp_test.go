package gomesh_test

import (
	"io"
	"os"
	"testing"

	"github.com/crhntr/gomesh"
)

func TestSupplementalRecordSet(t *testing.T) {
	f, err := os.Open("testdata/supp2017.xml")
	if err != nil {
		t.Fatal(err)
	}

	var (
		UI, Name, Created, Revised, Note, Frequency, Concepts, Sources, MappedTo bool
	)

	src, errc := gomesh.ParseSupplementalRecordSet(f)
	count := 0
	for {
		if testing.Short() && count > 100 {
			return
		}
		if testing.Verbose() && count > 1000 {
			return
		}
		if !testing.Verbose() && !testing.Short() && count > 200 {
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
