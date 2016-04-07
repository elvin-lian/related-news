package dedup

import (
	"testing"
)

func TestAppendToNewsMap(t *testing.T) {
	CleanMap()
	news := News{
		Id:  102,
		Sh1: 41250,
		Sh2: 23164,
		Sh3: 19884,
		Sh4: 52612,
	}
	if NewsMapLen() != 0 {
		t.Error("AppendToNewsMap should be eq 0")
	}
	AppendToNewsMap(&news)
	if NewsMapLen() != 1 {
		t.Error("AppendToNewsMap should be eq 1")
	}
}
