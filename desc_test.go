package gomesh_test

import (
	"io"
	"os"
	"testing"

	"github.com/crhntr/gomesh"
)

func TestParseDescriptorRecordSet(t *testing.T) {
	f, err := os.Open("testdata/desc2017.xml")
	if err != nil {
		t.Error(err)
	}

	drc, errc := gomesh.ParseDescriptorRecordSet(f)

	count := 0
	done := false
	for !done && count < 1000 {
		select {
		case <-drc:
			count++
		case err := <-errc:
			if err == io.EOF {
				close(errc)
				done = true
				break
			} else {
				t.Error(err)
				close(errc)
			}
		}
	}
	t.Log(count)
}
