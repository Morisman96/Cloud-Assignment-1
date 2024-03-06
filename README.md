# Cloud-Assignment-1

## Description
This repository contains the source code and documentation for a REST web application in Go that provides information about books available in a given language based on the Gutenberg library. The service also determines the number of potential readers presumed to be able to read books in that language. It integrates with three external APIs: Gutendex API, Language2Countries API, and REST Countries API.

### Table of Contents
1. [Usage](Usage) 
2. [Endpoints](Endpoints)


## Usage
The API is deployed and accessible at the following URL: https://cloud-assignment-1-9yeu.onrender.com/.

Endpoints
The API provides the following endpoints:

### Bookcount
`/librarystats/v1/bookcount/?language={:two_letter_language_code+}/`<br> 
Returns book count for a given language(s) code.

it will return the following JSON response:
```json
{
 "languages": "<two letter language code>",
 "count": "<number of books in that language>",
 "authors": "<number of authors for the books in that language>",
 "fraction": "<fraction that indicates how much of the books in the given language contribute to the total number of books available in Gutendex>"
}
```
Example:
```json
{
 "languages": "no",
 "count": 21,
 "authors": 16,
 "fraction": 0.0002877974
},
{
"languages": "sv",
"count": 230,
"authors": 139,
"fraction": 0.0031520666
}
```
### Readership
`/librarystats/v1/readership/{:two_letter_language_code}/{?limit={:number}}/`<br>
 Returns the number of potential readers for a given language code and an optional limit for how many countries to return.

it will return the following JSON response:
```json
[
  {
    "Official_Name": "<official name of the country>",
    "ISO3166_1_Alpha_2": "<two letter country code>",
    "count": "<number of books in that language>",
    "authors": "<number of authors for the books in that language>",
    "population": "<population of the country>"
  }
]
``` 

Example:
```json
[
 {
  "Official_Name": "Iceland",
  "ISO3166_1_Alpha_2": "IS",
  "count": 21,
  "authors": 16,
  "population": 366425
 },
 {
  "Official_Name": "Norway",
  "ISO3166_1_Alpha_2": "NO",
  "count": 21,
  "authors": 16,
  "population": 5379475
 },
 {
  "Official_Name": "Svalbard and Jan Mayen Islands",
  "ISO3166_1_Alpha_2": "SJ",
  "count": 21,
  "authors": 16,
  "population": 2562
 }
]
```
### Diagnostics
`/librarystats/v1/status/`<br>
Provides a status overview of services.

it will return the following JSON response:
```json
{
  "gutendexapi": "<http status code for gutendex API>",
  "languageapi": "<http status code for language2countries API>" , 
  "countriesapi": "<http status code for restcountries API>",
  "version": "version of the api",
  "uptime": ",<time in seconds from the last service restart>"
}
```
Example:
```json
{
  "gutendexapi": "200",
  "languageapi": "200",
  "countriesapi": "200",
  "version": "v1",
  "uptime": 60
}
```
