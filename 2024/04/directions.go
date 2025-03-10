package main

func (d data) up(i index) ([4]index, bool) {
	il := [4]index{}
	for n := range 4 {
		if i.i-n < 0 {
			return [4]index{}, false
		}
		il[n] = index{i: i.i - n, j: i.j}
	}

	return il, true
}

func (d data) down(i index) ([4]index, bool) {
	il := [4]index{}
	for n := range 4 {
		if i.i+n >= len(d) {
			return [4]index{}, false
		}
		il[n] = index{i: i.i + n, j: i.j}
	}

	return il, true
}

func (d data) right(i index) ([4]index, bool) {
	il := [4]index{}
	for n := range 4 {
		if i.j+n >= len(d) {
			return [4]index{}, false
		}
		il[n] = index{i: i.i, j: i.j + n}
	}

	return il, true
}

func (d data) left(i index) ([4]index, bool) {
	il := [4]index{}
	for n := range 4 {
		if i.j-n < 0 {
			return [4]index{}, false
		}
		il[n] = index{i: i.i, j: i.j - n}
	}

	return il, true
}

func (d data) upLeft(i index) ([4]index, bool) {
	il := [4]index{}
	for n := range 4 {
		if i.i-n < 0 || i.j-n < 0 {
			return [4]index{}, false
		}
		il[n] = index{i: i.i - n, j: i.j - n}
	}

	return il, true
}

func (d data) upRight(i index) ([4]index, bool) {
	il := [4]index{}
	for n := range 4 {
		if i.i-n < 0 || i.j+n >= len(d) {
			return [4]index{}, false
		}
		il[n] = index{i: i.i - n, j: i.j + n}
	}

	return il, true
}

func (d data) downLeft(i index) ([4]index, bool) {
	il := [4]index{}
	for n := range 4 {
		if i.i+n >= len(d) || i.j-n < 0 {
			return [4]index{}, false
		}
		il[n] = index{i: i.i + n, j: i.j - n}
	}

	return il, true
}

func (d data) downRight(i index) ([4]index, bool) {
	il := [4]index{}
	for n := range 4 {
		if i.i+n >= len(d) || i.j+n >= len(d) {
			return [4]index{}, false
		}
		il[n] = index{i: i.i + n, j: i.j + n}
	}

	return il, true
}
