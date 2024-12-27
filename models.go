package gospiget

// BaseModel contains common functionality for models
type BaseModel struct {
	ID int `json:"id"`
}

// ResourceFile represents a resource's file information
type ResourceFile struct {
	Type        string  `json:"type"`
	Size        float64 `json:"size"`
	SizeUnit    string  `json:"sizeUnit"`
	URL         string  `json:"url"`
	ExternalURL string  `json:"externalUrl"`
}

// Icon represents a resource icon or author avatar
type Icon struct {
	URL  string `json:"url"`
	Data string `json:"data"`
}

// ResourceRating represents rating information
type ResourceRating struct {
	Count   int     `json:"count"`
	Average float64 `json:"average"`
}

// ResourceVersion represents a version of a resource
type ResourceVersion struct {
	BaseModel
	UUID        string         `json:"uuid"`
	Name        string         `json:"name"`
	ReleaseDate int64          `json:"releaseDate"`
	Downloads   int            `json:"downloads"`
	Rating      ResourceRating `json:"rating"`
}

// ResourceUpdate represents an update to a resource
type ResourceUpdate struct {
	BaseModel
	Resource    int    `json:"resource"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        int64  `json:"date"`
	Likes       int    `json:"likes"`
}

// Author represents a resource author
type Author struct {
	BaseModel
	Name string `json:"name"`
	Icon *Icon  `json:"icon"`
}

// Category represents a resource category
type Category struct {
	BaseModel
	Name string `json:"name"`
}

// ResourceReview represents a review of a resource
type ResourceReview struct {
	Author          Author         `json:"author"`
	Rating          ResourceRating `json:"rating"`
	Message         string         `json:"message"`
	ResponseMessage string         `json:"responseMessage"`
	Version         string         `json:"version"`
	Date            int64          `json:"date"`
}

// Resource represents the main resource model
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
