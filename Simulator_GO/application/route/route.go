package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	Id        string     `json:"routeId"`
	ClientId  string     `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type PartialRoutePosition struct {
	Id       string    `json:"routeId"`
	ClientId string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (r *Route) LoadPositions() error {
	if r.Id == "" {
		return errors.New("route id not informed")
	}
	f, err := os.Open("destinations/" + r.Id + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		Lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}
		Lng, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Position{
			Lat: Lat,
			Lng: Lng,
		})
	}
	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for k, v := range r.Positions {
		route.Id = r.Id
		route.ClientId = r.ClientId
		route.Position = []float64{v.Lat, v.Lng}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
