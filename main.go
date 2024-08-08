package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
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

		// Process(Params{
		// 	Percent:     *Percent / 1000,
		// 	Sum:         *StartSum,
		// 	MonthIncome: *MonthIncome,
		// })
	} else {
		fmt.Println("Type -percent=<> -start=<> -month-income=<>\nType -start to start process")
	}
}

func main() {
	p := Params{
		Percent:     0,
		Sum:         0,
		MonthIncome: 0,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		flags := strings.Split(text, "-")

		for _, v := range flags {
			flag := (strings.Split(v, "="))
			if strings.ContainsAny(flag[0], "percent") {
				flagValue := ""

				if len(flag) == 2 {
					flagValue = strings.TrimSpace(flag[1])
				}

				value, _ := strconv.ParseFloat(flagValue, 64)

				// p.Percent = value
				fmt.Println(value)
			}

			if strings.ContainsAny(flag[0], "start") {
				flagValue := ""

				if len(flag) == 2 {
					flagValue = strings.TrimSpace(flag[1])
				}

				value, _ := strconv.ParseFloat(flagValue, 64)

				// p.Sum = value
				fmt.Println(value)
			}

			if strings.ContainsAny(flag[0], "income") {
				flagValue := ""

				if len(flag) == 2 {
					flagValue = strings.TrimSpace(flag[1])
				}

				value, _ := strconv.ParseFloat(flagValue, 64)

				// p.MonthIncome = value
				fmt.Println(value)
			}
		}

		if text == "start" {
			fmt.Println(p)
		}
	}
}

// func Parse() (bool, string) {
// 	if *StartSum == 0 {
// 		return false, "Start Sum not entered"
// 	}

// 	return true, ""
// }

func Process(p Params) {
	// fmt.Println(p)
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
