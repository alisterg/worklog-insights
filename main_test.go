package main

import "testing"

const (
	ExpectedAvgHours = 7.683333
	ExpectedWfhDays  = 2
)

func TestGenerateReport(t *testing.T) {
	result, err := GenerateReport("./samples/", "2024-01-01", "2024-01-10")
	if err != nil {
		t.Error("Couldn't generate report", err)
	}

	if result.AvgHours != ExpectedAvgHours {
		t.Errorf("incorrect average hours: expected %v, got %v", ExpectedAvgHours, result.AvgHours)
	}

	if result.WfhDays != ExpectedWfhDays {
		t.Errorf("incorrect wfh days: expected %v, got %v", ExpectedWfhDays, result.WfhDays)
	}
}
