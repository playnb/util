package mathex

func MaxInt(v ...int) int {
	r := 0
	for _, i := range v {
		if i > r {
			r = i
		}
	}
	return r
}

func MaxUint32(v ...uint32) uint32 {
	r := uint32(0)
	for _, i := range v {
		if i > r {
			r = i
		}
	}
	return r
}

func AbsInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
