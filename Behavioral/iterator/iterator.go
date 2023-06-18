package main

import "fmt"

type RadioStation struct {
	frequency float32
}

func (rs *RadioStation) GetFrequency() float32 {
	return rs.frequency
}

type StationList struct {
	stations []*RadioStation
	counter  int
}

func (sl *StationList) AddStation(station RadioStation) {
	sl.stations = append(sl.stations, &station)
}

func (sl *StationList) RemoveStation(station RadioStation) {
	for i, val := range sl.stations {
		if station == *val {
			copy(sl.stations[i:], sl.stations[i+1:])
			sl.stations[len(sl.stations)-1] = nil
			sl.stations = sl.stations[:len(sl.stations)-1]
			break
		}
	}
}

func (sl *StationList) Count() int {
	return len(sl.stations)
}

func (sl *StationList) Current() *RadioStation {
	return sl.stations[sl.counter]
}

func (sl *StationList) Key() int {
	return sl.counter
}

func (sl *StationList) Next() {
	if (sl.counter + 1) < len(sl.stations) {
		sl.counter++
	}
}

func (sl *StationList) Prev() {
	if (sl.counter - 1) >= 0 {
		sl.counter--
	}
}

func (sl *StationList) Reset() {
	sl.counter = 0
}

func (sl *StationList) HasNext() bool {
	return (sl.counter + 1) < len(sl.stations)
}

func main() {
	stationsList := StationList{}

	stationsList.AddStation(RadioStation{123})
	stationsList.AddStation(RadioStation{45.7})
	stationsList.AddStation(RadioStation{67.2})
	stationsList.AddStation(RadioStation{89.4})

	fmt.Println(stationsList.Current().GetFrequency())
	stationsList.Next()
	fmt.Println(stationsList.Current().GetFrequency())
	stationsList.RemoveStation(RadioStation{45.7})
	fmt.Println(stationsList.Current().GetFrequency())

}
