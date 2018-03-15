package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/StageAutoControl/controller/fixtures"
	"github.com/creasty/go-easing"
)

func main() {
	logger := logrus.New()

	ds := fixtures.DataStore()
	t := ds.DMXTransitions["a1a02b6c-12dd-4d7b-bc3e-24cc823adf21"]

	p := t.Params[0]
	blue1 := float64(p.From.Blue.Value)
	blue2 := float64(p.To.Blue.Value)
	diff := blue2 - blue1
	length := float64(100) // float64(t.Length)
	step := float64(1.0) / length

	logger.Infof("length=%v step=%v blue1=%v blue2=%v", length, step, blue1, blue2)

	file, err := os.Create("circular_result.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := float64(0); i <= length; i++ {
		value := math.Floor(blue1 + diff*easing.CircularEaseInOut(i*step))
		err := writer.Write([]string{fmt.Sprintf("%.6f", i), fmt.Sprintf("%.6f", value)})
		if err != nil {
			panic(err)
		}
		logger.Infof("Step=%v value=%v", i, value)
	}
}
