package dedup

import "testing"

func TestHaimingDistant(t *testing.T) {
	s1 := "0010110001111100010101101101001110011001011010111101010110011111"
	s2 := "1010110101011001010100000100000000110100000100101110100010000010"

	res, _ := hamming(s1, s2)

	if res != 30 {
		t.Errorf("expect haiming distant %s va %s eq 30, but got %d", s1, s2, res)
	}
}

func TestUniqSliceInt64(t *testing.T) {
	a := []int64{1, 2, 3, 2, 1}
	b := uniqSliceInt64(a)
	if len(b) != 3 {
		t.Errorf("uniqSliceInt64 fail %v", b)
	}
}

func TestCheckTitle(t *testing.T) {
	TitleMap[8450] = []int64{101}
	TitleMap[2584] = []int64{101}
	TitleMap[19884] = []int64{101, 102}
	TitleMap[50564] = []int64{101}

	TitleMap[41250] = []int64{102}
	TitleMap[23164] = []int64{102}
	TitleMap[52612] = []int64{102}

	NewsMap[101] = []uint16{0, 0, 0, 0, 8450, 2584, 19884, 50564}
	NewsMap[102] = []uint16{0, 0, 0, 0, 41250, 23164, 19884, 52612}


	sh := [8]uint16{0, 0, 0, 0, 8226, 19048, 19756, 52228}
	res := checkTitle(&sh)
	if res {
		t.Errorf("Test1: checkTitile fail")
	}

	sh = [8]uint16{0, 0, 0, 0, 41250, 23164, 19884, 52612}
	res = checkTitle(&sh)
	if !res {
		t.Errorf("Test2: checkTitile fail")
	}
}

func TestCheckContent(t *testing.T) {
	ContMap[8450] = []int64{101}
	ContMap[2584] = []int64{101}
	ContMap[19884] = []int64{101}
	ContMap[50564] = []int64{101}

	ContMap[41250] = []int64{102}
	ContMap[23164] = []int64{102}
	ContMap[52612] = []int64{102}

	NewsMap[101] = []uint16{8450, 2584, 19884, 50564, 0, 0, 0, 0}
	NewsMap[102] = []uint16{41250, 23164, 19884, 52612, 0, 0, 0, 0}

	sh := [8]uint16{8226, 19048, 19756, 52228, 0, 0, 0, 0}
	res := checkContent(&sh)
	if res {
		t.Errorf("Test1: checkContent fail")
	}

	sh = [8]uint16{41250, 23164, 19884, 52612, 0, 0, 0, 0}
	res = checkContent(&sh)
	if !res {
		t.Errorf("Test2: checkContent fail")
	}
}

