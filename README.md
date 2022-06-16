# Go API client

Apartment aggregator

- API version: 1.0.0
- Build date: 2022-06-16T10:47:13.999319+03:00[Europe/Minsk]

This is an apartment aggregator that collects information about available apartments, cottages and plots for purchase or
rent

## Requirements

Building the API client library requires:

1. Go 1.12+
2. PostgreSQL database 14

## Installation

Install the following dependencies:

```shell
go get github.com/gin-gonic/gin v1.8.1
go get github.com/gocolly/colly v1.2.0
go get github.com/jackc/pgx/v4 v4.16.1
go get github.com/sirupsen/logrus v1.8.1
go get github.com/spf13/viper v1.12.0
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080/flats*

| Go files              | Method/Function                                                                  | HTTP request                                          | Description                                                        |
|-----------------------|----------------------------------------------------------------------------------|-------------------------------------------------------|--------------------------------------------------------------------|
| *cottagesHandler*     | [**getAllAreas**](handlers/cottagesHandler#getAllAreas)                          | **Get** /cottages/sale/areas/{page}                   | Returns all available areas to purchase.                           |
| *cottagesHandler*     | [**getAllHouses**](handlers/cottagesHandler#getAllHouses)                        | **Get** /cottages/sale/houses/{page}                  | Returns all available houses to purchase.                          |
| *cottagesHandler*     | [**getAllDacha**](handlers/cottagesHandler#getAllDacha)                          | **Get** /cottages/sale/vitebskiyi-rayion/dacha/{page} | Returns all available country houses to purchase.                  |
| *flatsHandler*        | [**getOneRoomFlats**](handlers/flatsHandler#getOneRoomFlats)                     | **Get** /flats/sale/one-room-flats/{page}             | Returns all available one room apartments to purchase.             |
| *flatsHandler*        | [**getAllFlats**](handlers/flatsHandler#getAllFlats)                             | **Get** /flats/sale/{page}                            | Returns all available apartments to purchase.                      |
| *flatsHandler*        | [**getMoreThanThreeRoomFlats**](handlers/flatsHandler#getMoreThanThreeRoomFlats) | **Get** /flats/sale/three-more-room-flats/{page}      | Returns all available more than three room apartments to purchase. |
| *flatsHandler*        | [**getThreeRoomFlats**](handlers/flatsHandler#getThreeRoomFlats)                 | **Get** /flats/sale/three-room-flats/{page}           | Returns all available three room apartments to purchase.           |
| *flatsHandler*        | [**getTwoRoomFlats**](handlers/flatsHandler#getTwoRoomFlats)                     | **Get** /flats/sale/two-room-flats/{page}             | Returns all available two room apartments to purchase.             |
| *flatsForRentHandler* | [**getCottagesForDay**](handlers/flatsForRentHandler#getCottagesForDay)          | **Get** /flats/rent/cottages-for-day/{page}           | Returns all available cottages for day rent.                       |
| *flatsForRentHandler* | [**getCottagesLongRent**](handlers/flatsForRentHandler#getCottagesLongRent)      | **Get** /flats/rent/cottages-for-long/{page}          | Returns all available cottages for long rent.                      |
| *flatsForRentHandler* | [**getFlatsForDay**](handlers/flatsForRentHandler#getFlatsForDay)                | **Get** /flats/rent/flats-for-day/{page}              | Returns all available flats for day/days rent.                     |
| *flatsForRentHandler* | [**getFlatsLongRent**](handlers/flatsForRentHandler#getFlatsLongRent)            | **Get** /flats/rent/flats-for-long/{page}             | Returns all available flats for long rent.                         |

## Documentation for Models

- [Areas](models/area)
- [Cottages](models/cottage)
- [Flats](models/flat)

## Documentation for Authorization

All endpoints do not require authorization.
Authentication schemes defined for the API:

## Author
yury.bliznets@innowise-group.com


