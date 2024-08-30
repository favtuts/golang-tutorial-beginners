package main

import "testing"

func average(score1, score2, score3 int) int {
	return ((score1 + score2 + score3) / 3)
}
func Test_arerage(t *testing.T) {
	score1, score2, score3 := 10, 18, 41

	averageScore := average(score1, score2, score3)
	if averageScore == 0 {
		t.Fail()
	}

}
