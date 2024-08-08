package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

var (
	Percent     = flag.Float64("percent", 0, "percent of acc")
	StartSum    = flag.Float64("startSum", 0, "start sum on acc")
	MonthIncome = flag.Float64("income", 0, "month income to acc")
)

type Params struct {
	Percent     float64
	Sum         float64
	MonthIncome float64
}

type Result struct {
	Sum     float64
	Percent float64
}

func init() {
	flag.Parse()

	fmt.Printf("Parsing flags:\nPerecent = %v\nStart Sum = %v\nMonth Income = %v\n", *Percent, *StartSum, *MonthIncome)
}

func main() {
	p := Params{
		Percent:     *Percent / 1000,
		Sum:         *StartSum,
		MonthIncome: *MonthIncome,
	}

	var result []Result
	for i := 0; i < 12; i++ {
		percentIncome := math.Round(p.Sum * p.Percent)

		p.Sum = p.Sum + percentIncome + p.MonthIncome

		result = append(result, Result{
			Sum:     p.Sum,
			Percent: percentIncome,
		})
	}

	ShowResult(result)
}

func ShowResult(r []Result) {
	time := time.Now()

	for i, v := range r {
		t := time.AddDate(0, i, 0).Format("January, 2006")
		fmt.Printf("Month: %v\nSum end of month: %v, Percent: %v\n\n", t, v.Sum, v.Percent)
	}
}
