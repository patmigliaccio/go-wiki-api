# go-wiki-api

A simple Go service that retrieves Wikipedia article information.

## Development

### Local Environment

Run the Go server locally

```bash
$ go run *.go
```

Accessible at `http://localhost:3000`

### Local Deployment

Docker Build

```bash
$ docker build -t go-wiki-api .
```

Docker Run

```bash
$ docker run --publish 7200:80 --name go-wiki-api --rm -d go-wiki-api
```

Accessible at `http://localhost:7200`

### Production Deployment

Prerequisites

```bash
$ yarn global add now
```

Now Deployment

```bash
$ now
```

## Usage

Prerequisites (*Optional*)

[HTTPie](https://httpie.org/) - command line HTTP client

```bash
$ apt-get install httpie
# --or--
$ brew install httpie
```

### API

#### GET Extracts

Extracts the content for a list of titles delimited by `|` and then URI encoded (e.g. `Jimmy Wales|Steve Wozniak`)

```bash
$ http http://localhost:3000/api/v1.0/extracts/Jimmy%20Wales
```

##### Response

```json
[
 {
   "metadata": {
     "id": 3703446,
     "title": "Jimmy Wales",
     "url": "https://en.wikipedia.org/wiki/Jimmy_Wales"
   },
   "extract": "Jimmy Donal Wales (born August 7, 1966), also known by the online moniker \"Jimbo\", is an American..."
 }
]
```

#### GET Search

Retrieves a list of pages based on a search value. Specifically tailored for typing in the prefix of a word for autocomplete functionality.

* `limit=50` - Limits the amount of search results returned from the API (*optional*) 

```bash
$ http http://localhost:3000/api/v1.0/search/Jimmy?limit=10
```

##### Response

```json
[
  {
    "id": 2954486,
    "title": "Jimmy Lin",
    "url": "https://en.wikipedia.org/wiki/Jimmy_Lin"
  },
  {
    "id": 3703446,
    "title": "Jimmy Wales",
    "url": "https://en.wikipedia.org/wiki/Jimmy_Wales"
  }
  ...
]
```

#### GET Categories

Retrieves the categories associated with an article based on a specified Wikipedia `pageid`.

```bash
$ http http://localhost:3000/api/v1.0/categories/3703446
```

##### Response

```json
{
  "metadata": {
    "id": 3703446,
    "title": "Jimmy Wales",
    "url": "https://en.wikipedia.org/wiki/Jimmy_Wales"
  },
  "extract": "",
  "categories": [
    "1966_births",
    "Living_people",
    "American_bloggers",
    "American_expatriates_in_the_United_Kingdom",
    "American_libertarians",
    "American_technology_company_founders",
    "Auburn_University_alumni",
    "Recipients_of_the_Gottlieb_Duttweiler_Prize"
  ]
}
```

#### GET Sections

Retrieves the sections within an article based on a specified Wikipedia `pageid`.

```bash
$ http http://localhost:3000/api/v1.0/sections/3703446
```

##### Response

```json
{
  "metadata": {
    "id": 3703446,
    "title": "Jimmy Wales",
    "url": "https://en.wikipedia.org/wiki/Jimmy_Wales"
  },
  "extract": "",
  "categories": null,
  "sections": [
    "Early_life",
    "Career",
    "Chicago_Options_Associates_and_Bomis",
    "Nupedia_and_the_origins_of_Wikipedia",
    "Wikipedia",
    "Controversy_regarding_Wales's_status_as_co-founder",
    "Role",
    ...
  ]
}
```