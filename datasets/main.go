package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jttait/go/datasets/dateparser"
)

func main() {
	//UKCPIURL := "https://www.ons.gov.uk/generator?format=csv&uri=/economy/inflationandpriceindices/timeseries/l522/mm23"
	//downloadCSV(UKCPIURL, "raw_data/ONS_UK_Consumer_Price_Index")
	//UKCPIs, err := convertToRecords("raw_data/ONS_UK_Consumer_Price_Index", 0, 1, 186)
	//if err != nil {
	//		log.Fatal(err)
	//	}
	//	UKCPIs = interpolateRecords(UKCPIs)
	//	writeRecordsToCSV(UKCPIs, "records/UK_Consumer_Price_Index")

	//downloadCSV(landRegistryURL("united-kingdom"), "raw_data/Land_Registry_Nominal_UK_Average_House_Prices")
	//nominalUKPrices, err := convertToRecords("raw_data/Land_Registry_Nominal_UK_Average_House_Prices", 3, 6, 1)
	//if err != nil {
	//		log.Fatal(err)
	//	}
	//	nominalUKPrices = interpolateRecords(nominalUKPrices)
	//	writeRecordsToCSV(nominalUKPrices, "records/Nominal_UK_Average_House_Prices")
	//	realUKPrices := adjustForInflation(nominalUKPrices, UKCPIs)
	//	writeRecordsToCSV(realUKPrices, "records/Real_UK_Average_House_Prices")

	//downloadCSV(landRegistryURL("city-of-aberdeen"), "raw_data/Land_Registry_Nominal_Aberdeen_Average_House_Prices")
	//nominalAberdeenPrices, err := convertToRecords("raw_data/Land_Registry_Nominal_Aberdeen_Average_House_Prices", 3, 6, 1)
	//if err != nil {
	//		log.Fatal(err)
	//	}
	//	nominalAberdeenPrices = interpolateRecords(nominalAberdeenPrices)
	//	writeRecordsToCSV(nominalAberdeenPrices, "records/Nominal_Aberdeen_Average_House_Prices")
	//	realAberdeenPrices := adjustForInflation(nominalAberdeenPrices, UKCPIs)
	//	writeRecordsToCSV(realAberdeenPrices, "records/Real_Aberdeen_Average_House_Prices")

	//	downloadCSV(landRegistryURL("shetland-islands"), "raw_data/Land_Registry_Nominal_Shetland_Average_House_Prices")
	//	nominalShetlandPrices, err := convertToRecords("raw_data/Land_Registry_Nominal_Shetland_Average_House_Prices", 3, 6, 1)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	nominalShetlandPrices = interpolateRecords(nominalShetlandPrices)
	//	writeRecordsToCSV(nominalShetlandPrices, "records/Nominal_Shetland_Average_House_Prices")
	//	realShetlandPrices := adjustForInflation(nominalShetlandPrices, UKCPIs)
	//	writeRecordsToCSV(realShetlandPrices, "records/Real_Shetland_Average_House_Prices")

	//	downloadCSV(landRegistryURL("london"), "raw_data/Land_Registry_Nominal_London_Average_House_Prices")
	//	nominalLondonPrices, err := convertToRecords("raw_data/Land_Registry_Nominal_London_Average_House_Prices", 3, 6, 1)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	nominalLondonPrices = interpolateRecords(nominalLondonPrices)
	//	writeRecordsToCSV(nominalLondonPrices, "records/Nominal_London_Average_House_Prices")
	//	realLondonPrices := adjustForInflation(nominalLondonPrices, UKCPIs)
	//	writeRecordsToCSV(realLondonPrices, "records/Real_London_Average_House_Prices")

	BOEBaseRates := scrapeBankOfEnglandRates()
	BOEBaseRates = interpolateRecords(BOEBaseRates)
	writeRecordsToCSV(BOEBaseRates, "records/BOEBaseRates")
}

type Record struct {
	Date  time.Time
	Value float64
}

type ByDate []Record

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date.Before(a[j].Date) }

func interpolateRecords(records []Record) []Record {
	sort.Sort(ByDate(records))
	var interpolatedRecords []Record
	for i := 0; i < len(records)-1; i++ {
		date := records[i].Date
		value := records[i].Value
		nextDate := records[i+1].Date
		for date.Before(nextDate) {
			interpolatedRecords = append(interpolatedRecords, Record{date, value})
			date = date.AddDate(0, 0, 1)
		}
	}
	return interpolatedRecords
}

func scrapeBankOfEnglandRates() []Record {
	resp, err := http.Get("https://www.bankofengland.co.uk/boeapps/database/Bank-Rate.asp#")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s\n", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var records []Record
	doc.Find("#stats-table").Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		var record Record
		s.Find("td").Each(func(j int, q *goquery.Selection) {
			align, _ := q.Attr("align")
			if align == "left" {
				date, err := dateparser.ParseDate(q.Text())
				if err != nil {
					log.Fatal(err)
				}
				record.Date = date
			} else if align == "right" {
				trimmed := strings.TrimSpace(q.Text())
				num, err := strconv.ParseFloat(trimmed, 64)
				if err != nil {
					log.Fatal(err)
				}
				record.Value = num
			}
		})
		records = append(records, record)
	})
	return records
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

func convertToRecords(filename string, dateColumn int, valueColumn int, numHeaderRows int) ([]Record, error) {
	f, err := os.Open(filename + ".csv")
	if err != nil {
		return []Record{}, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return []Record{}, err
	}
	var records []Record
	for i := numHeaderRows; i < len(data); i++ {
		date, err := dateparser.ParseDate(data[i][dateColumn])
		if err != nil {
			return records, err
		}
		value, err := strconv.ParseFloat(data[i][valueColumn], 64)
		if err != nil {
			return records, err
		}
		records = append(records, Record{date, value})
	}
	return records, nil
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

func earliestDateThatIsInBothRecords(a []Record, b []Record) time.Time {
	date := a[0].Date
	if b[0].Date.After(date) {
		date = b[0].Date
	}
	return date
}

func latestDateThatIsInBothRecords(a []Record, b []Record) time.Time {
	date := a[len(a)-1].Date
	if b[len(b)-1].Date.Before(date) {
		date = b[len(b)-1].Date
	}
	return date
}

func indexForDate(records []Record, date time.Time) (int, error) {
	for index, record := range records {
		if record.Date.Equal(date) {
			return index, nil
		}
	}
	return 0, fmt.Errorf("Date is not in records")
}

func adjustForInflation(nominalRecords []Record, CPIRecords []Record) []Record {
	startDate := earliestDateThatIsInBothRecords(nominalRecords, CPIRecords)
	endDate := latestDateThatIsInBothRecords(nominalRecords, CPIRecords)
	currentNominal, err := indexForDate(nominalRecords, startDate)
	if err != nil {
		log.Fatal(err)
	}
	endNominal, err := indexForDate(nominalRecords, endDate)
	if err != nil {
		log.Fatal(err)
	}
	currentCPI, err := indexForDate(CPIRecords, startDate)
	if err != nil {
		log.Fatal(err)
	}
	endCPI, err := indexForDate(CPIRecords, endDate)
	if err != nil {
		log.Fatal(err)
	}
	var realRecords []Record
	for currentNominal <= endNominal && currentCPI <= endCPI {
		realValue := nominalRecords[currentNominal].Value * (CPIRecords[len(CPIRecords)-1].Value / CPIRecords[currentCPI].Value)
		realRecords = append(realRecords, Record{nominalRecords[currentNominal].Date, realValue})
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
