package task4

import (
	"testing"
)

func Test_reverseString(t *testing.T) {

}

func Benchmark_reverseString(b *testing.B) {

}

func Test_isPal(t *testing.T) {

}

func Benchmark_isPal(b *testing.B) {

}

func Test_findLargest(t *testing.T) {

}

func Benchmark_findLargest(b *testing.B) {

}

func Test_slicetonum(t *testing.T) {

}

func Benchmark_slicetonum(b *testing.B) {

}

func Test_numtoslice(t *testing.T) {

}

func Benchmark_numtoslice(b *testing.B) {

}

func Test_dotask(t *testing.T) {

}


func Benchmark_dotask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			dotask(j)
		}
	}
}