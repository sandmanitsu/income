package main

import (
	"bufio"
	"flag"
	"fmt"
	"income/internal/parse"
	"math"
	"os"
	"strings"
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

	// if start from GO RUN with flags -percent=<> -startSum=<> -income=<>
	if *Percent != 0 || *StartSum != 0 || *MonthIncome != 0 {
		fmt.Printf("Parsing flags:\nPerecent = %v\nStart Sum = %v\nMonth Income = %v\n", *Percent, *StartSum, *MonthIncome)

		Process(Params{
			Percent:     *Percent / 1000,
			Sum:         *StartSum,
			MonthIncome: *MonthIncome,
		})
	} else {
		fmt.Println("Type -percent=<> -start=<> -month-income=<>\nType <result> to start process")
	}
}

func main() {
	var p Params

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		flags := strings.Split(text, "-")

		for _, v := range flags {
			flag, value := parse.Parse(v)

			switch flag {
			case "percent":
				p.Percent = value / 1000
			case "start":
				p.Sum = value
			case "income":
				p.MonthIncome = value
			case "result":
				Process(p)
			case "exit":
				os.Exit(1)
			case "help":
				fmt.Println("Type -percent=<value> -start=<value> -month-income=<value>\tType <result> to start process\n<exit> to exit the programm")
			}
		}
	}
}

func Process(p Params) {
	var result []Result

	for i := 0; i < 12; i++ {
		percentIncome := math.Round(p.Sum * p.Percent)

		p.Sum = p.Sum + percentIncome + p.MonthIncome

		result = append(result, Result{
			Sum:     math.Round(p.Sum),
			Percent: percentIncome,
		})
	}

	ShowResult(result)
}

func ShowResult(r []Result) {
	time := time.Now()

	for i, v := range r {
		t := time.AddDate(0, i, 0).Format("January, 2006")
		fmt.Printf("Month: %v\tSum end of month: %v, Percent: %v\n", t, v.Sum, v.Percent)
	}
}
