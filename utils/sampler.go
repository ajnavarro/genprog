package utils

import (
	"math/rand"
)

type Data struct {
	Vars   []interface{}
	Result interface{}
}

type DataSet struct {
	// TODO make this work with other kind of inputs
	From, To  int
	NSamples  int
	NVars     int
	Evaluator func(ins ...interface{}) interface{}
}

func (d *DataSet) Generate() []*Data {
	var varSamples [][]int
	for i := 0; i < d.NVars; i++ {
		varSamples = append(varSamples, randInts(d.From, d.To, d.NSamples))
	}

	var data []*Data
	for i := 0; i < d.NSamples; i++ {
		var in []interface{}
		for j := 0; j < d.NVars; j++ {
			in = append(in, varSamples[j][i])
		}

		out := d.Evaluator(in...)

		d := &Data{
			Vars:   in,
			Result: out,
		}
		data = append(data, d)
	}

	return data
}

func randInts(min, max int, n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(max-min) + min
	}
	return res
}
