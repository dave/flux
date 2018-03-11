package flux_test

import (
	"testing"

	"fmt"

	"github.com/dave/flux"

	"github.com/dave/ktest/assert"
)

type Flights struct {
	Dispatcher *flux.Dispatcher
	Country    *CountryStore
	City       *CityStore
	Price      *PriceStore
}

func TestFlights(t *testing.T) {

	a := &Flights{}
	a.Country = &CountryStore{app: a}
	a.City = &CityStore{app: a}
	a.Price = &PriceStore{app: a}
	a.Dispatcher = flux.NewDispatcher(nil, a.Country, a.City, a.Price)

	done := a.Dispatcher.Dispatch(&UpdateCountryAction{Country: "France"})
	<-done

	assert.Equal(t, "France", a.Country.GetCountry())
	assert.Equal(t, "Paris", a.City.GetCity())
	assert.Equal(t, "€100", a.Price.GetPrice())
}

type UpdateCity struct {
	City string
}

type CityStore struct {
	app  *Flights
	city string
}

func (m *CityStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *UpdateCity:
		m.city = action.City
	case *UpdateCountryAction:
		payload.Wait(m.app.Country)
		m.city = getCapital(m.app.Country.GetCountry())
	}
	return true
}

func getCapital(country string) string {
	switch country {
	case "UK":
		return "London"
	case "France":
		return "Paris"
	}
	return fmt.Sprint("Capital of ", country)
}

func (m *CityStore) GetCity() string {
	return m.city
}

type UpdateCountryAction struct {
	Country string
}

type CountryStore struct {
	app     *Flights
	country string
}

func (m *CountryStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *UpdateCountryAction:
		m.country = action.Country
	}
	return true
}

func (m *CountryStore) GetCountry() string {
	return m.country
}

type PriceStore struct {
	app   *Flights
	price string
}

func (m *PriceStore) Handle(payload *flux.Payload) (finished bool) {
	switch payload.Action.(type) {
	case *UpdateCountryAction, *UpdateCity:
		payload.Wait(m.app.City)
		m.price = calculatePrice(m.app.Country.GetCountry(), m.app.City.GetCity())
	}
	return true
}

func (m *PriceStore) GetPrice() string {
	return m.price
}

func calculatePrice(country string, city string) string {
	switch country {
	case "UK":
		return "£100"
	case "France":
		return "€100"
	}
	return "$100"
}
