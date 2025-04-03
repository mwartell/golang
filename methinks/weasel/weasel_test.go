package weasel

import "testing"

func BenchmarkNewExperiment(b *testing.B) {
	for b.Loop() {
		_, err := NewExperiment("methinks it is like a weasel", 100, 0.05, CharsetDefault)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkNewPool(b *testing.B) {
	for b.Loop() {
		_ = newPool(100, 30, CharsetDefault)
	}
}

func BenchmarkMutantChild(b *testing.B) {
	exp, err := NewExperiment("methinks it is like a weasel", 100, 0.05, CharsetDefault)
	if err != nil {
		b.Error(err)
	}
	for b.Loop() {
		_ = exp.mutantChild("methinks it is like a weasel")
	}
}

func BenchmarkRatchetSolve(b *testing.B) {
	var generations int

	for b.Loop() {
		exp, err := NewExperiment("methinks it is like a weasel", 100, 0.05, CharsetDefault)
		if err != nil {
			b.Error(err)
		}
		generations = exp.RatchetSolve()
	}
	b.Logf("Generations: %d\n", generations)
}

func BenchmarkPoolSolve(b *testing.B) {
	exp, err := NewExperiment("methinks it is like a weasel", 100, 0.05, CharsetDefault)
	if err != nil {
		b.Error(err)
	}
	generations := 0
	for b.Loop() {
		generations = PoolSolve(exp)
	}
	b.Logf("Generations: %d\n", generations)
}
