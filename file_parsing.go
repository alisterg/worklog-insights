package main

import (
	"bufio"
	"os"
	"strings"
	"time"
)

type WorklogTable struct {
	HoursWorked float32
	WFH         bool
}

func ParseFile(filePath string) (*WorklogTable, error) {
	file, err := os.Open(DirPrefix + filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	table := &WorklogTable{
		HoursWorked: 0,
		WFH:         false,
	}

	var startTime string
	var finishTime string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Start") {
			startTime = getTableRowValue(line)
		}
		if strings.Contains(line, "Finish") {
			finishTime = getTableRowValue(line)
		}
		if strings.Contains(line, "WFH") {
			table.WFH = getTableRowValue(line) == "y"
		}
	}
	table.HoursWorked, err = getTimeWorked(startTime, finishTime)
	if err != nil {
		return nil, err
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return table, nil
}

func getTableRowValue(tableRow string) string {
	fields := strings.Split(tableRow, "|")
	return strings.TrimSpace(fields[2])
}

func getTimeWorked(start string, finish string) (float32, error) {
	startTime, err := time.Parse("15:04", start)
	if err != nil {
		return 0, err
	}

	endTime, err := time.Parse("15:04", finish)
	if err != nil {
		return 0, err
	}

	duration := endTime.Sub(startTime)
	hours := float32(duration.Hours())

	return hours, nil
}
