package main

import "os"
import "fmt"
import "encoding/csv"

const usage string = "prepare-csv <filename.csv>\n" +
	"    prepares file for parsing by gnuplot\n" +
	"    writes to stdout"

/* read all records from csv file */
func read_csv(filename string) (records [][]string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(f)
	records, read_err := reader.ReadAll()
	if read_err != nil {
		panic(read_err)
	}

	f.Close()

	return records
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	records := read_csv(args[0])
	for _, record := range records[1:] {
		if len(record[2]) != 0 {
			fmt.Println(record[0] + "," + record[2])
		}
	}
}
