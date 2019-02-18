package utils

import "math/rand"

type Data struct {
	Vars   []interface{}
	Result interface{}
}

type DataSet struct {
	// TODO make this work with other kind of inputs
	From, To  float64
	NSamples  int
	NVars     int
	Evaluator func(ins ...interface{}) interface{}
}

func (d *DataSet) Generate() []*Data {
	var varSamples [][]float64
	for i := 0; i < d.NVars; i++ {
		varSamples = append(varSamples, randFloats(d.From, d.To, d.NSamples))
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

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}
