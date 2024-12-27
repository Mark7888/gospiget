# Golang Spiget API Wrapper

This is a Golang API wrapper for the [Spiget API](https://spiget.org/), providing easy access to Spigot resources, authors, and categories.

## Installation
To add this package to your project, run:
```sh
go get github.com/Mark7888/gospiget
```

## Models
The following models are used in the API:

### BaseModel
Contains common functionality for models.
```go
type BaseModel struct {
	ID int `json:"id"`
}
```

### ResourceFile
Represents a resource's file information.
```go
type ResourceFile struct {
	Type        string  `json:"type"`
	Size        float64 `json:"size"`
	SizeUnit    string  `json:"sizeUnit"`
	URL         string  `json:"url"`
	ExternalURL string  `json:"externalUrl"`
}
```

### Icon
Represents a resource icon or author avatar.
```go
type Icon struct {
	URL  string `json:"url"`
	Data string `json:"data"`
}
```

### ResourceRating
Represents rating information.
```go
type ResourceRating struct {
	Count   int     `json:"count"`
	Average float64 `json:"average"`
}
```

### ResourceVersion
Represents a version of a resource.
```go
type ResourceVersion struct {
	BaseModel
	UUID        string         `json:"uuid"`
	Name        string         `json:"name"`
	ReleaseDate int64          `json:"releaseDate"`
	Downloads   int            `json:"downloads"`
	Rating      ResourceRating `json:"rating"`
}
```

### ResourceUpdate
Represents an update to a resource.
```go
type ResourceUpdate struct {
	BaseModel
	Resource    int    `json:"resource"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        int64  `json:"date"`
	Likes       int    `json:"likes"`
}
```

### Author
Represents a resource author.
```go
type Author struct {
	BaseModel
	Name string `json:"name"`
	Icon *Icon  `json:"icon"`
}
```

### Category
Represents a resource category.
```go
type Category struct {
	BaseModel
	Name string `json:"name"`
}
```

### ResourceReview
Represents a review of a resource.
```go
type ResourceReview struct {
	Author          Author         `json:"author"`
	Rating          ResourceRating `json:"rating"`
	Message         string         `json:"message"`
	ResponseMessage string         `json:"responseMessage"`
	Version         string         `json:"version"`
	Date            int64          `json:"date"`
}
```

### Resource
Represents the main resource model.
```go
type Resource struct {
	BaseModel
	Name           string            `json:"name"`
	Tag            string            `json:"tag"`
	Contributors   string            `json:"contributors"`
	Likes          int               `json:"likes"`
	File           *ResourceFile     `json:"file"`
	TestedVersions []string          `json:"testedVersions"`
	Links          map[string]string `json:"links"`
	Rating         *ResourceRating   `json:"rating"`
	ReleaseDate    int64             `json:"releaseDate"`
	UpdateDate     int64             `json:"updateDate"`
	Downloads      int               `json:"downloads"`
	External       bool              `json:"external"`
	Icon           *Icon             `json:"icon"`
	Premium        bool              `json:"premium"`
	Price          float64           `json:"price"`
	Currency       string            `json:"currency"`
	Description    string            `json:"description"`
	Documentation  string            `json:"documentation"`
	SourceCodeLink string            `json:"sourceCodeLink"`
	DonationLink   string            `json:"donationLink"`
}
```

## Client
To use the client, import the package and create a new client instance.

### Importing the Client
```go
import "github.com/Mark7888/gospiget"
```

### Creating a Client
```go
client := gospiget.NewClient()
```

### Client Functions
#### GetStatus
Retrieves the status of the Spiget API.
```go
status, err := client.GetStatus()
```

#### GetResources
Retrieves a list of resources with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
resources, err := client.GetResources(params)
```

#### GetResourceByID
Retrieves a resource by its ID.
```go
resource, err := client.GetResourceByID(123)
```

#### GetResourceAuthor
Retrieves the author of a resource by the resource ID.
```go
author, err := client.GetResourceAuthor(123)
```

#### GetResourceVersions
Retrieves the versions of a resource with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
versions, err := client.GetResourceVersions(123, params)
```

#### GetResourceVersionByID
Retrieves a specific version of a resource by the resource ID and version ID.
```go
version, err := client.GetResourceVersionByID(123, 1)
```

#### GetLatestResourceVersion
Retrieves the latest version of a resource by the resource ID.
```go
latestVersion, err := client.GetLatestResourceVersion(123)
```

#### GetResourceUpdates
Retrieves the updates of a resource with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
updates, err := client.GetResourceUpdates(123, params)
```

#### GetLatestResourceUpdate
Retrieves the latest update of a resource by the resource ID.
```go
latestUpdate, err := client.GetLatestResourceUpdate(123)
```

#### GetResourceReviews
Retrieves the reviews of a resource with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
reviews, err := client.GetResourceReviews(123, params)
```

#### GetAuthors
Retrieves a list of authors with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
authors, err := client.GetAuthors(params)
```

#### GetAuthorByID
Retrieves an author by their ID.
```go
author, err := client.GetAuthorByID(123)
```

#### GetAuthorResources
Retrieves the resources of an author with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
resources, err := client.GetAuthorResources(123, params)
```

#### GetAuthorReviews
Retrieves the reviews of an author with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
reviews, err := client.GetAuthorReviews(123, params)
```

#### GetCategories
Retrieves a list of categories with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
categories, err := client.GetCategories(params)
```

#### GetCategoryByID
Retrieves a category by its ID.
```go
category, err := client.GetCategoryByID(123)
```

#### GetCategoryResources
Retrieves the resources of a category with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
resources, err := client.GetCategoryResources(123, params)
```

#### SearchResources
Searches for resources by a query string with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
resources, err := client.SearchResources("query", params)
```

#### SearchAuthors
Searches for authors by a query string with optional query parameters.
```go
params := map[string]string{"size": "10"} // Available parameters are listed below
authors, err := client.SearchAuthors("query", params)
```

### Query Parameters
The following query parameters can be used with the client functions:

- `size`: Size of the returned array (e.g., `size=10`)
- `page`: Page index (e.g., `page=1`)
- `sort`: Field to sort by. Use a `+` or `-` prefix for ascending or descending order (e.g., `sort=+name`)
- `fields`: Fields to return, separated by commas (e.g., `fields=id,name`)

## Client Error Types
The following error types are used in the client:

### NotFoundError
Represents a 404 Not Found error.
```go
type NotFoundError struct {
	Message string
}
```
Thrown when a resource or author is not found.

### UnexpectedStatusCodeError
Represents an unexpected status code error.
```go
type UnexpectedStatusCodeError struct {
	StatusCode int
}
```
Thrown when the API returns an unexpected status code.

### UnmarshalError
Represents an error during unmarshalling.
```go
type UnmarshalError struct {
	Message string
}
```
Thrown when there is an error unmarshalling the response body.

### RequestError
Represents an error during the request.
```go
type RequestError struct {
	Message string
}
```
Thrown when there is an error making the request.

## Credits
This wrapper is built for the [Spiget API](https://spiget.org/) created by the [SpiGetOrg team](https://github.com/SpiGetOrg).
