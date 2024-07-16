package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DirPrefix = "./"
)

type Report struct {
	WfhDays  int
	AvgHours float32
}

func main() {
	fromDate, toDate, err := parseCliFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	report, err := GenerateReport(DirPrefix, fromDate, toDate)
	if err != nil {
		fmt.Println("Couldn't get files: ", err)
		os.Exit(1)
	}

	fmt.Println("WFH days          : ", report.WfhDays)
	fmt.Println("Average work hours: ", report.AvgHours)
}

func parseCliFlags() (string, string, error) {
	fromDate := flag.String("from", "", "A date in the format YYYY-MM-DD")
	toDate := flag.String("to", "", "A date in the format YYYY-MM-DD")
	flag.Parse()

	if *fromDate == "" || *toDate == "" {
		flag.Usage()
		return "", "", fmt.Errorf("`-from` and `-to` arguments are required")
	}

	return *fromDate, *toDate, nil
}

// fromDate is inclusive, toDate is exclusive
func GenerateReport(dirname string, fromDate string, toDate string) (*Report, error) {
	result := &Report{
		WfhDays:  0,
		AvgHours: 0,
	}

	dayCount := 0
	var totalHoursWorked float32 = 0

	files, err := GetFilenames(dirname)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		shouldParse, err := ShouldParseFile(file, fromDate, toDate)
		if err != nil {
			return nil, err
		}

		if shouldParse {
			worklog, err := ParseFile(file.Name())
			if err != nil {
				return nil, fmt.Errorf("couldn't parse file %v: %v", file.Name(), err)
			}

			dayCount++
			totalHoursWorked += worklog.HoursWorked

			if worklog.WFH {
				result.WfhDays++
			}
		}
	}

	result.AvgHours = totalHoursWorked / float32(dayCount)

	return result, nil
}
