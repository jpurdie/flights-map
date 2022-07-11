package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Route struct {
	airlineCode string
	origin      string
	destination string
}

type Airport struct {
	id           string
	name         string
	code         string
	coordinatesX string
	coordinatesY string
}

func createRoutesList(data [][]string) []Route {
	var routesList []Route
	for _, line := range data {
		a := Route{
			airlineCode: line[0],
			origin:      line[2],
			destination: line[4],
		}
		routesList = append(routesList, a)
	}
	return routesList
}

func persistRoutesData(routes []Route) {
	m := make(map[string]int)
	for _, route := range routes {
		m[route.origin+"-"+route.destination]++
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	csvFile, err := os.Create("my-flights-data.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	row := []string{"airline", "airport1", "airport2", "cnt"}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

	for _, k := range keys {
		s := strings.Split(k, "-")

		row := []string{"XX", s[0], s[1], strconv.Itoa(m[k])}

		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}

func createAirporstList(data [][]string) []Airport {
	var airportsList []Airport
	for _, line := range data {
		a := Airport{
			id:           line[0],
			name:         line[1],
			code:         line[4],
			coordinatesX: line[6],
			coordinatesY: line[7],
		}
		airportsList = append(airportsList, a)
	}
	return airportsList
}

func persistAirportsData(airportList []Airport) {
	csvFile, err := os.Create("my-aiports-data.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	row := []string{"iata", "airport", "city", "state", "country", "lat", "long"}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

	for _, airport := range airportList {
		row := []string{airport.code, airport.name, "", "", "", airport.coordinatesX, airport.coordinatesY}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}

func main() {
	// open file
	f, err := os.Open("src-airports.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	airportsList := createAirporstList(data)
	persistAirportsData(airportsList)

	f2, err2 := os.Open("src-routes.csv")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer f2.Close()

	csvReader2 := csv.NewReader(f2)
	data2, err2 := csvReader2.ReadAll()
	if err2 != nil {
		log.Fatal(err2)
	}

	routesList := createRoutesList(data2)
	persistRoutesData(routesList)

}
