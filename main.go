package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/alecthomas/kong"
	"github.com/mileusna/useragent"
	w3c "github.com/stephane-martin/w3c-extendedlog-parser"
)

var cli struct {
	CpuProfile string `flag:"cpu-profile" help:"Enable CPU profiling."`
	File       string `arg:"" name:"file" help:"W3C log file to read." type:"string"`
}

func main() {
	kong.Parse(&cli)

	if cli.CpuProfile != "" {
		f, err := os.Create(cli.CpuProfile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	err := readLines(cli.File)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
}

func readLines(path string) error {
	// read line by line through a file
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	m := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, vals, err := w3c.ExtractStrings([]byte(fmt.Sprintf("%s\n", line)))
		if err != nil {
			return err
		}
		// fmt.Println(rest, vals[9])

		ua := useragent.Parse(vals[9])
		switch {
		case ua.Mobile:
			m["mobile"] += 1
		case ua.Bot:
			m["bot"] += 1
		case ua.Desktop:
			m["desktop"] += 1
		case ua.Tablet:
			m["tablet"] += 1
		}
	}

	fmt.Println(m)

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
