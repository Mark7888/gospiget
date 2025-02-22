package gospiget

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"math/rand"

	"github.com/go-resty/resty/v2"
)

const baseURL = "https://api.spiget.org/v2"

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15",
}

func getRandomUserAgent() string {
	return userAgents[rand.Intn(len(userAgents))]
}

type Client struct {
	restyClient *resty.Client
}

func NewClient() *Client {
	client := resty.New()
	client.SetBaseURL(baseURL)
	client.SetTimeout(10 * time.Second)
	client.SetHeader("User-Agent", getRandomUserAgent())
	return &Client{restyClient: client}
}

func (c *Client) GetStatus() (map[string]interface{}, error) {
	resp, err := c.restyClient.R().Get("/status")
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return result, nil
}

func (c *Client) GetResources(params map[string]string) ([]Resource, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get("/resources")
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var resources []Resource
	if err := json.Unmarshal(resp.Body(), &resources); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return resources, nil
}

func (c *Client) GetResourceByID(resourceID int) (*Resource, error) {
	resp, err := c.restyClient.R().Get(fmt.Sprintf("/resources/%d", resourceID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() == http.StatusNotFound {
		return nil, &NotFoundError{Message: "resource not found"}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var resource Resource
	if err := json.Unmarshal(resp.Body(), &resource); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return &resource, nil
}

func (c *Client) GetResourceAuthor(resourceID int) (*Author, error) {
	resp, err := c.restyClient.R().Get(fmt.Sprintf("/resources/%d/author", resourceID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() == http.StatusNotFound {
		return nil, &NotFoundError{Message: "resource author not found"}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var author Author
	if err := json.Unmarshal(resp.Body(), &author); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return &author, nil
}

func (c *Client) GetResourceVersions(resourceID int, params map[string]string) ([]ResourceVersion, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get(fmt.Sprintf("/resources/%d/versions", resourceID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var versions []ResourceVersion
	if err := json.Unmarshal(resp.Body(), &versions); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return versions, nil
}

func (c *Client) GetResourceVersionByID(resourceID, versionID int) (*ResourceVersion, error) {
	resp, err := c.restyClient.R().Get(fmt.Sprintf("/resources/%d/versions/%d", resourceID, versionID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() == http.StatusNotFound {
		return nil, &NotFoundError{Message: "resource version not found"}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var version ResourceVersion
	if err := json.Unmarshal(resp.Body(), &version); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return &version, nil
}

func (c *Client) GetLatestResourceVersion(resourceID int) (*ResourceVersion, error) {
	resp, err := c.restyClient.R().Get(fmt.Sprintf("/resources/%d/versions/latest", resourceID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() == http.StatusNotFound {
		return nil, &NotFoundError{Message: "latest resource version not found"}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var version ResourceVersion
	if err := json.Unmarshal(resp.Body(), &version); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return &version, nil
}

func (c *Client) GetResourceUpdates(resourceID int, params map[string]string) ([]ResourceUpdate, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get(fmt.Sprintf("/resources/%d/updates", resourceID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var updates []ResourceUpdate
	if err := json.Unmarshal(resp.Body(), &updates); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return updates, nil
}

func (c *Client) GetLatestResourceUpdate(resourceID int) (*ResourceUpdate, error) {
	resp, err := c.restyClient.R().Get(fmt.Sprintf("/resources/%d/updates/latest", resourceID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() == http.StatusNotFound {
		return nil, &NotFoundError{Message: "latest resource update not found"}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var update ResourceUpdate
	if err := json.Unmarshal(resp.Body(), &update); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return &update, nil
}

func (c *Client) GetResourceReviews(resourceID int, params map[string]string) ([]ResourceReview, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get(fmt.Sprintf("/resources/%d/reviews", resourceID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var reviews []ResourceReview
	if err := json.Unmarshal(resp.Body(), &reviews); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return reviews, nil
}

func (c *Client) GetAuthors(params map[string]string) ([]Author, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get("/authors")
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var authors []Author
	if err := json.Unmarshal(resp.Body(), &authors); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return authors, nil
}

func (c *Client) GetAuthorByID(authorID int) (*Author, error) {
	resp, err := c.restyClient.R().Get(fmt.Sprintf("/authors/%d", authorID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() == http.StatusNotFound {
		return nil, &NotFoundError{Message: "author not found"}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var author Author
	if err := json.Unmarshal(resp.Body(), &author); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return &author, nil
}

func (c *Client) GetAuthorResources(authorID int, params map[string]string) ([]Resource, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get(fmt.Sprintf("/authors/%d/resources", authorID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var resources []Resource
	if err := json.Unmarshal(resp.Body(), &resources); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return resources, nil
}

func (c *Client) GetAuthorReviews(authorID int, params map[string]string) ([]ResourceReview, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get(fmt.Sprintf("/authors/%d/reviews", authorID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var reviews []ResourceReview
	if err := json.Unmarshal(resp.Body(), &reviews); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return reviews, nil
}

func (c *Client) GetCategories(params map[string]string) ([]Category, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get("/categories")
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var categories []Category
	if err := json.Unmarshal(resp.Body(), &categories); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return categories, nil
}

func (c *Client) GetCategoryByID(categoryID int) (*Category, error) {
	resp, err := c.restyClient.R().Get(fmt.Sprintf("/categories/%d", categoryID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() == http.StatusNotFound {
		return nil, &NotFoundError{Message: "category not found"}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var category Category
	if err := json.Unmarshal(resp.Body(), &category); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return &category, nil
}

func (c *Client) GetCategoryResources(categoryID int, params map[string]string) ([]Resource, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get(fmt.Sprintf("/categories/%d/resources", categoryID))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var resources []Resource
	if err := json.Unmarshal(resp.Body(), &resources); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return resources, nil
}

func (c *Client) SearchResources(query string, params map[string]string) ([]Resource, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get(fmt.Sprintf("/search/resources/%s", query))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var resources []Resource
	if err := json.Unmarshal(resp.Body(), &resources); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return resources, nil
}

func (c *Client) SearchAuthors(query string, params map[string]string) ([]Author, error) {
	resp, err := c.restyClient.R().SetQueryParams(params).Get(fmt.Sprintf("/search/authors/%s", query))
	if err != nil {
		return nil, &RequestError{Message: err.Error()}
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}
	var authors []Author
	if err := json.Unmarshal(resp.Body(), &authors); err != nil {
		return nil, &UnmarshalError{Message: err.Error()}
	}
	return authors, nil
}

func (c *Client) DownloadResourceVersion(resource ResourceVersion, path string, proxy bool) error {
	var url string
	if proxy {
		url = fmt.Sprintf("/resources/%d/versions/%d/download/proxy", resource.ResourceId, resource.ID)
	} else {
		url = fmt.Sprintf("/resources/%d/versions/%d/download", resource.ResourceId, resource.ID)
	}

	resp, err := c.restyClient.R().Get(url)
	if err != nil {
		return &RequestError{Message: err.Error()}
	}

	if resp.StatusCode() != http.StatusOK {
		return &UnexpectedStatusCodeError{StatusCode: resp.StatusCode()}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(resp.Body())
	if err != nil {
		return err
	}

	return nil
}
