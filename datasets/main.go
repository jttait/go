package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jttait/go/datasets/dateparser"
)

func main() {
	UKCPIURL := "https://www.ons.gov.uk/generator?format=csv&uri=/economy/inflationandpriceindices/timeseries/l522/mm23"
	downloadCSV(UKCPIURL, "raw_data/ONS_UK_Consumer_Price_Index")
	UKCPIs := convertToRecords("raw_data/ONS_UK_Consumer_Price_Index", 0, 1, 186)
	writeRecordsToCSV(UKCPIs, "records/UK_Consumer_Price_Index")

	downloadCSV(landRegistryURL("united-kingdom"), "raw_data/Land_Registry_Nominal_UK_Average_House_Prices")
	nominalUKPrices := convertToRecords("raw_data/Land_Registry_Nominal_UK_Average_House_Prices", 3, 6, 1)
	writeRecordsToCSV(nominalUKPrices, "records/Nominal_UK_Average_House_Prices")
	realUKPrices := adjustForInflation(nominalUKPrices, UKCPIs)
	writeRecordsToCSV(realUKPrices, "records/Real_UK_Average_House_Prices")

	downloadCSV(landRegistryURL("city-of-aberdeen"), "raw_data/Land_Registry_Nominal_Aberdeen_Average_House_Prices")
	nominalAberdeenPrices := convertToRecords("raw_data/Land_Registry_Nominal_Aberdeen_Average_House_Prices", 3, 6, 1)
	writeRecordsToCSV(nominalAberdeenPrices, "records/Nominal_Aberdeen_Average_House_Prices")
	realAberdeenPrices := adjustForInflation(nominalAberdeenPrices, UKCPIs)
	writeRecordsToCSV(realAberdeenPrices, "records/Real_Aberdeen_Average_House_Prices")

	downloadCSV(landRegistryURL("shetland-islands"), "raw_data/Land_Registry_Nominal_Shetland_Average_House_Prices")
	nominalShetlandPrices := convertToRecords("raw_data/Land_Registry_Nominal_Shetland_Average_House_Prices", 3, 6, 1)
	writeRecordsToCSV(nominalShetlandPrices, "records/Nominal_Shetland_Average_House_Prices")
	realShetlandPrices := adjustForInflation(nominalShetlandPrices, UKCPIs)
	writeRecordsToCSV(realShetlandPrices, "records/Real_Shetland_Average_House_Prices")

	downloadCSV(landRegistryURL("london"), "raw_data/Land_Registry_Nominal_London_Average_House_Prices")
	nominalLondonPrices := convertToRecords("raw_data/Land_Registry_Nominal_London_Average_House_Prices", 3, 6, 1)
	writeRecordsToCSV(nominalLondonPrices, "records/Nominal_London_Average_House_Prices")
	realLondonPrices := adjustForInflation(nominalLondonPrices, UKCPIs)
	writeRecordsToCSV(realLondonPrices, "records/Real_London_Average_House_Prices")

}

type Record struct {
	Date  time.Time
	Value float64
}

func downloadCSV(URL string, filename string) {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(file, resp.Body)
	defer file.Close()
}

func convertToRecords(filename string, dateColumn int, valueColumn int, numHeaderRows int) []Record {
	f, err := os.Open(filename + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var records []Record
	for i := numHeaderRows; i < len(data); i++ {
		var start time.Time
		var num float64
		for j, field := range data[i] {
			if j == dateColumn {
				start, err = dateparser.ParseDate(field)
				if err != nil {
					log.Fatal(err)
				}
			} else if j == valueColumn {
				num, err = strconv.ParseFloat(field, 64)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		for d := start; d.Month() == start.Month(); d = d.AddDate(0, 0, 1) {
			record := Record{d, num}
			records = append(records, record)
		}
	}
	return records
}

func writeRecordsToCSV(records []Record, filename string) {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < len(records); i++ {
		dateString := records[i].Date.Format("2006-01-02")
		valueString := strconv.FormatFloat(records[i].Value, 'f', 2, 64)
		writer.Write([]string{dateString, valueString})
	}
}

func adjustForInflation(nominalRecords []Record, CPIRecords []Record) []Record {
	startDate := nominalRecords[0].Date
	if CPIRecords[0].Date.After(startDate) {
		startDate = CPIRecords[0].Date
	}
	endDate := nominalRecords[len(nominalRecords)-1].Date
	if CPIRecords[len(CPIRecords)-1].Date.Before(endDate) {
		endDate = CPIRecords[len(CPIRecords)-1].Date
	}
	var startNominal int
	var endNominal int
	for i := 0; i < len(nominalRecords); i++ {
		if nominalRecords[i].Date.Equal(startDate) {
			startNominal = i
		}
		if nominalRecords[i].Date.Equal(endDate) {
			endNominal = i
		}
	}
	var startCPI int
	var endCPI int
	for i := 0; i < len(CPIRecords); i++ {
		if CPIRecords[i].Date.Equal(startDate) {
			startCPI = i
		}
		if CPIRecords[i].Date.Equal(endDate) {
			endCPI = i
		}
	}

	var realRecords []Record
	currentNominal := startNominal
	currentCPI := startCPI
	for currentNominal <= endNominal && currentCPI <= endCPI {
		if nominalRecords[currentNominal].Date.After(startDate) && nominalRecords[currentNominal].Date.Before(endDate) {
			realValue := nominalRecords[currentNominal].Value * (CPIRecords[len(CPIRecords)-1].Value / CPIRecords[currentCPI].Value)
			realRecords = append(realRecords, Record{nominalRecords[currentNominal].Date, realValue})
		}
		currentNominal++
		currentCPI++
	}
	return realRecords
}

func landRegistryURL(region string) string {
	return "https://landregistry.data.gov.uk/app/ukhpi/download/new.csv?" +
		"from=1900-01-01&to=2100-01-01&location=http%3A%2F%2Flandregistry.data.gov.uk%2Fid" +
		"%2Fregion%2F" + region + "&thm%5B%5D=property_type&in%5B%5D=avg&lang=en"
}
