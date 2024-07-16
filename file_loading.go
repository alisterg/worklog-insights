package main

import (
	"io/fs"
	"os"
	"regexp"
)

func GetFilenames(dirname string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func ShouldParseFile(file fs.DirEntry, fromDate string, toDate string) (bool, error) {
	matchesFilePattern, err := matchesFilePattern(file)
	if err != nil {
		return false, err
	}

	isWithinDates := isWithinDates(file.Name(), fromDate, toDate)
	return matchesFilePattern && isWithinDates, nil
}

func matchesFilePattern(file fs.DirEntry) (bool, error) {
	pattern, err := regexp.Compile(`^\d{4}-\d{2}-\d{2}\.md$`)
	if err != nil {
		return false, err
	}

	return !file.IsDir() && pattern.MatchString(file.Name()), nil
}

func isWithinDates(filename string, fromDate string, toDate string) bool {
	return filename >= fromDate && filename < toDate
}
