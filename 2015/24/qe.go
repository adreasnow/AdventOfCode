package main

func calculateQE(ints []int) int {
	QE := 1
	for _, n := range ints {
		QE = QE * n
	}
	return QE
}

func smallestQE(groups [][]int) int {
	QE := uint64(0)
	QE--
	for _, group := range groups {
		if qe := uint64(calculateQE(group)); qe < QE {
			QE = qe
		}
	}

	return int(QE)
}
