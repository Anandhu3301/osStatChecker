package pkg

import "math"


func ByteToGiga(memoryValues []uint64)[]float64 {
	var memorySlice []float64;
	for _, mem := range memoryValues {
		memorySlice = append(memorySlice, operationInsideByteToGiga(mem))
	}
	return  memorySlice;
}

func roundFloatNumbers(number float64, precision uint) float64 {
	ratio := math.Pow10(int(precision))
	return math.Round(number*ratio) / ratio
}

func operationInsideByteToGiga(mem uint64) float64 {
	divideValue := math.Pow(1024,3)
	var gigaValue float64 = float64(mem) / divideValue
	var floatingGiga float64 = roundFloatNumbers(gigaValue, 2)
	return floatingGiga
}
