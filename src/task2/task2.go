package task2

import (
	"fmt"
	"math"
)

// укоротить длинные куски кода

type T2 struct {
	E1 Envelope
	E2 Envelope
}

func (T *T2) Dotask2() {
	_, d, _ := Dotask(T.E1, T.E2)
	fmt.Println(d)
}

type Envelope struct {
	Side1, Side2 float64
}

func (e *Envelope) isBiggerThan(e2 *Envelope) (b bool) {
	if e.Side1*e.Side2 > e2.Side1*e2.Side2 {
		return true
	}
	return false
}

func (e *Envelope) biggerSide() (f float64) {
	if e.Side1 >= e.Side2 {
		return e.Side1

	} else if e.Side1 < e.Side2 {
		return e.Side2
	}

	return
}

func (e *Envelope) smallerSide() (f float64) {
	if e.Side1 < e.Side2 {
		return e.Side1

	} else if e.Side2 <= e.Side1 {
		return e.Side2
	}

	return
}

func (e *Envelope) findDiagonal() (f float64) {
	f = float64(math.Sqrt(e.Side1*e.Side1 + e.Side2*e.Side2))
	return
}

func Dotask(e1 Envelope, e2 Envelope) (ok bool, data interface{}, reason string) {
	A, B := e1.Side1, e1.Side2
	C, D := e2.Side1, e2.Side2
	if A <= 0 || B <= 0 || C <= 0 || D <= 0 {
		return false, nil, "Ошибка!\nСторона конверта не может быть меньше или равно 0"

	} else if (A == C || A == D) && (B == C || B == D) {
		return true, 0, reason

	} else if e1.isBiggerThan(&e2) && (e1.findDiagonal()-e2.biggerSide() >= e2.smallerSide()) {
		return true, 2, reason

	} else if e2.isBiggerThan(&e1) && (e2.findDiagonal()-e1.biggerSide() >= e1.smallerSide()) {
		return true, 1, reason
	}

	return true, 0, reason
}
