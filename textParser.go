/*
Text file parser template
*/

package main

import "fmt"
import "os"
import "log"
import "flag"
import "path/filepath"
import "bufio"

import "time"

var flagVerbose bool
var flagQuiet bool
var dateFormat=time.RFC1123

func parsefile(filePath string) ([]string, error) {

  inputFile, err := os.Open(filePath)
  if err != nil {
    return nil, err
  }
  defer inputFile.Close()	

	scanner := bufio.NewScanner(inputFile)
	var results []string
	for scanner.Scan() {
		if output, add := parser(scanner.Text()); add {
			results = append(results, output)
		}
	}
	return results, nil
}

func parser(text string) (string, bool) {
	return text, true
}

func main() {
	flag.BoolVar(&flagVerbose, "v", false, "Prints detailed operations")
	flag.BoolVar(&flagQuiet, "q", false, "No output apart from errors")
	flag.Parse()

	//items := []string{"."}  // default arguments to use if omitted

	if flag.NArg() == 0 {
		flag.PrintDefaults()
		return
	}

	start := time.Now()
	for _, i := range flag.Args() {
		items, err := filepath.Glob(i)
		if err != nil { log.Fatal(err) }
		for _, j := range items {
			output, err := parsefile(j)
			if err != nil { log.Fatal(err) }
			for _, l := range output {
				fmt.Println(l)
			}
		}
	}



	if !flagQuiet {
		elapsed := time.Since(start)
		fmt.Printf("Processed in %v\n" , elapsed)
	}
}


