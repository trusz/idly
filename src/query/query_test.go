package query_test

import (
	"testing"

	"github.com/trusz/idly/src/test_helpers/assert"
)

// year := flag.String("y", "", "Year")
// month := flag.String("m", "", "Month")
// day := flag.String("d", "", "Day")
// isGroup := flag.Bool("group", false, "Groups by activity")
// app := flag.String("app", "", "Summ of time by given app name")

func TestFileChooser(t *testing.T) {
	year := "2000"
	month := "01"
	day := "01"

	expectedFilePath := "log-2000-01-01.txt"

	acualFilePath := filePath(year, month, day)

	assert.Equal(t, expectedFilePath, acualFilePath)

}
