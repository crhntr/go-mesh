package mesh_test

import (
	"io"
	"os"
	"testing"

	mesh "github.com/crhntr/go-mesh"
)

func TestParseDescriptorRecordSet(t *testing.T) {
	t.Parallel()
	f, err := os.OpenFile("testdata/desc2017.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Error(err)
	}

	count := 0
	if err := mesh.ScanDescriptorRecordSet(f, func(r *mesh.DescriptorRecord) error {
		if testing.Short() && count > 100 {
			return io.EOF
		}
		count++
		return nil
	}); err != nil {
		if err != io.EOF {
			t.Error(err)
		}
	}
	t.Logf("count: %d", count-1)
}

func TestParsePharmacologicalActionSet(t *testing.T) {
	t.Parallel()
	f, err := os.OpenFile("testdata/pa2017.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Error(err)
	}

	count := 0
	if err := mesh.ScanPharmacologicalActionSet(f, func(r *mesh.PharmacologicalAction) error {
		if testing.Short() && count > 100 {
			return io.EOF
		}
		count++
		return nil
	}); err != nil {
		if err != io.EOF {
			t.Error(err)
		}
	}
	t.Logf("count: %d", count-1)
}

func TestSupplementalRecordSet(t *testing.T) {
	t.Parallel()
	f, err := os.OpenFile("testdata/supp2017.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Fatal(err)
	}

	var (
		UI, Name, Created, Revised, Note, Frequency, Concepts, Sources, MappedTo bool
	)

	count := 0

	if err := mesh.ScanSupplementalRecordSet(f, func(r *mesh.SupplementalRecord) error {
		if testing.Short() && count > 100 {
			return io.EOF
		}
		count++
		if string(r.UI) == "" {
			t.Errorf("supp record %d is missing UI", count)
		}
		if len(r.UI) > 0 {
			UI = true
		}
		if len(r.Name) > 0 {
			Name = true
		}
		if !r.Created.IsZero() {
			Created = true
		}
		if !r.Revised.IsZero() {
			Revised = true
		}
		if len(r.Note) > 0 {
			Note = true
		}
		if r.Frequency > 0 {
			Frequency = true
		}
		if len(r.Concepts) > 0 {
			Concepts = true
		}
		if len(r.Sources) > 0 {
			Sources = true
		}
		if len(r.MappedTo) > 0 {
			MappedTo = true
		}
		return nil
	}); err != nil {
		if err != io.EOF {
			t.Error(err)
		}
	}

	if !UI {
		t.Error("UI does not have at least one non-zero value")
	}
	if !Name {
		t.Error("Name does not have at least one non-zero value")
	}
	if !Created {
		t.Error("Created does not have at least one non-zero value")
	}
	if !Revised {
		t.Error("Revised does not have at least one non-zero value")
	}
	if !Note {
		t.Error("Note does not have at least one non-zero value")
	}
	if !Frequency {
		t.Error("Frequency does not have at least one non-zero value")
	}
	if !Concepts {
		t.Error("Concepts does not have at least one non-zero value")
	}
	if !Sources {
		t.Error("Sources does not have at least one non-zero value")
	}
	if !MappedTo {
		t.Error("MappedTo does not have at least one non-zero value")
	}
	t.Logf("count: %d", count-1)
}

func TestParseQualifierRecordSet(t *testing.T) {
	t.Parallel()
	f, err := os.OpenFile("testdata/qual2017.xml", os.O_RDONLY, 0755)
	if err != nil {
		t.Error(err)
	}

	count := 0
	if err := mesh.ScanQualifierRecordSet(f, func(r *mesh.QualifierRecord) error {
		if testing.Short() && count > 100 {
			return io.EOF
		}
		count++
		return nil
	}); err != nil {
		if err != io.EOF {
			t.Error(err)
		}
	}
	t.Logf("count: %d", count-1)
}
