package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/trusz/idly/src/tracking"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Idly tracks activity and provides queries to it.\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()
	}

	// year := flag.String("y", "", "Year")
	// month := flag.String("m", "", "Month")
	// day := flag.String("d", "", "Day")
	// isGroup := flag.Bool("group", false, "Groups by activity")
	// app := flag.String("app", "", "Summ of time by given app name")

	flag.Parse()

	// log.Println()
	// log.Println(*year)
	// log.Println(*month)
	// log.Println(*day)
	// log.Println(*isGroup)
	// log.Println(*app)

	if !isInQueryMode() {
		tracking.Start()
	}

}

// ./idly -y 2019 -m 6 -group
// ./idly -y 2019 -m 6 -d 5 -group
// ./idly -y 2019 -m 6 -d 5 -app Code

func isInQueryMode() bool {
	return flag.NFlag() > 0
}
