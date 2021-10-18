package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "net/http"
    "os"

    "golang.org/x/text/encoding/charmap"
)

// Structure for a PEP
type PEP struct {
    Name string
    Birthday string
}

var PEPs []PEP

// Read a csv file to a slice of slices
func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(f))
    csvReader.Comma = ';'
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

// Get all lines of csv file that contains a PEP
func collect(lines [][]string) []PEP {
    var PEPs []PEP
    for _, line := range lines {
        if line[1] != "" && line[2] != "" && line[4] != "" {
            PEPs = append(PEPs, PEP{line[2] + " " + line[1], line[4]})
        }
    }
    return PEPs
}

// Check if a PEP is in a slice of PEPs
func contains(PEPs []PEP, lookupPEP PEP) bool {
    for _, pep := range PEPs {
        if pep == lookupPEP {
            return true
        }
    }
    return false
}

// Handler for queries with name and birthday
func PEPHandler(w http.ResponseWriter, r *http.Request) {
    v := r.URL.Query()

    firstname := v.Get("firstname")
    lastname := v.Get("lastname")
    birthday := v.Get("birthday")

    lookupPEP := PEP{firstname + " " + lastname, birthday}
    fmt.Fprintf(w, "%t\n", contains(PEPs, lookupPEP))
}

func main() {
    // Read file
    records := readCsvFile("../PEP_listen.csv")
    // Collect PEPs
    PEPs = collect(records)
    // Handler for queries
    http.HandleFunc("/", PEPHandler)
    // Listen for http requests
    log.Fatal(http.ListenAndServe(":3000", nil))
}
