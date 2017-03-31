package gomesh_test

import (
	"io"
	"os"
	"testing"

	"github.com/crhntr/gomesh"
)

func TestParseQualifierRecordSet(t *testing.T) {
	f, err := os.Open("testdata/qual2017.xml")
	if err != nil {
		t.Error(err)
	}

	qrc, errc := gomesh.ParseQualifierRecordSet(f)

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
