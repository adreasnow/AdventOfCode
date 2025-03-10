package main

type data [][]rune

type index struct {
	i int
	j int
}

func (d data) iterateIndexes(part int) int {
	count := 0
	for i := range len(d) {
		for j := range len(d) {
			switch part {
			case 1:
				count += d.checkIndexP1(index{i, j})
			case 2:
				count += d.checkIndexP2(index{i, j})
			}
		}
	}
	return count
}

func (d data) checkIndexP1(i index) int {
	indicesToChceck := d.genIndexesP1(i)
	count := 0

	for _, index := range indicesToChceck {
		if d[index[0].i][index[0].j] == 'X' {
			if d[index[1].i][index[1].j] == 'M' {
				if d[index[2].i][index[2].j] == 'A' {
					if d[index[3].i][index[3].j] == 'S' {
						count++
					}
				}
			}
		}
	}
	return count
}

func (d data) checkIndexP2(i index) int {
	indicesToChceck := d.genIndexesP2(i)
	count := 0

	for _, index := range indicesToChceck {
		if d[index[0].i][index[0].j] == 'M' {
			if d[index[1].i][index[1].j] == 'M' {
				if d[index[2].i][index[2].j] == 'A' {
					if d[index[3].i][index[3].j] == 'S' {
						if d[index[4].i][index[4].j] == 'S' {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func (d data) genIndexesP1(i index) [][4]index {
	indexes := [][4]index{}

	if il, check := d.up(i); check {
		indexes = append(indexes, il)
	}
	if il, check := d.down(i); check {
		indexes = append(indexes, il)
	}
	if il, check := d.right(i); check {
		indexes = append(indexes, il)
	}
	if il, check := d.left(i); check {
		indexes = append(indexes, il)
	}
	if il, check := d.upLeft(i); check {
		indexes = append(indexes, il)
	}
	if il, check := d.upRight(i); check {
		indexes = append(indexes, il)
	}
	if il, check := d.downLeft(i); check {
		indexes = append(indexes, il)
	}
	if il, check := d.downRight(i); check {
		indexes = append(indexes, il)
	}
	return indexes
}

func (d data) genIndexesP2(i index) [4][5]index {
	indexes := [4][5]index{}

	if i.i-1 < 0 || i.j-1 < 0 || i.i+1 >= len(d) || i.j+1 >= len(d) {
		return [4][5]index{}
	}

	// M.M     00    02
	// .A.        11
	// S.S     20    22
	indexes[0][0] = index{i: i.i - 1, j: i.j - 1}
	indexes[0][1] = index{i: i.i - 1, j: i.j + 1}
	indexes[0][2] = index{i: i.i, j: i.j}
	indexes[0][3] = index{i: i.i + 1, j: i.j - 1}
	indexes[0][4] = index{i: i.i + 1, j: i.j + 1}

	// M.S     00    02
	// .A.        11
	// M.S     20    22
	indexes[1][0] = index{i: i.i + 1, j: i.j - 1}
	indexes[1][1] = index{i: i.i - 1, j: i.j - 1}
	indexes[1][2] = index{i: i.i, j: i.j}
	indexes[1][3] = index{i: i.i + 1, j: i.j + 1}
	indexes[1][4] = index{i: i.i - 1, j: i.j + 1}

	// S.M      00    02
	// .A.         11
	// S.M      20    22
	indexes[2][0] = index{i: i.i - 1, j: i.j + 1}
	indexes[2][1] = index{i: i.i + 1, j: i.j + 1}
	indexes[2][2] = index{i: i.i, j: i.j}
	indexes[2][3] = index{i: i.i - 1, j: i.j - 1}
	indexes[2][4] = index{i: i.i + 1, j: i.j - 1}

	// S.S     00    02
	// .A.        11
	// M.M     20    22
	indexes[3][0] = index{i: i.i + 1, j: i.j + 1}
	indexes[3][1] = index{i: i.i + 1, j: i.j - 1}
	indexes[3][2] = index{i: i.i, j: i.j}
	indexes[3][3] = index{i: i.i - 1, j: i.j + 1}
	indexes[3][4] = index{i: i.i - 1, j: i.j - 1}

	return indexes
}
