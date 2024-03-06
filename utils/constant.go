package utils

import "time"

var Timer time.Time // Variable to store the timer for when the program starts

const VERSION = "v1" // Version of the service

const PORT = ":8080" // local port

const Language2CountriesAPI = "http://129.241.150.113:3000/language2countries/" // API for language2countries

const GutendexAPI = "http://129.241.150.113:8000/books" // 		API for gutendex

const RESTCountriesAPI = "http://129.241.150.113:8080/v3.1/name/" // API for restcountries

const BOOKCOUNTPATH = "/librarystats/v1/bookcount/" // Path for bookcount

const READERSHIPPATH = "/librarystats/v1/readership/" // Path for readership

const STATUSPATH = "/librarystats/v1/status/" // Path for status

const REGEXPOSETIVINTEGER = "^[1-9]\\d*$" // Regex for positive integers

const STATUSLANGUAGE2COUNTRIESQUERY = "no" // Query for language2countries to check if the service is up
const STATUSRESTCOUNTRIESQUERY = "norway"  // Query for restcountries to check if the service is up
