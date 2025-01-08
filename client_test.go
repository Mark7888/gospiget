package gospiget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c := NewClient()

	// Test GetStatus
	status, err := c.GetStatus()
	assert.NoError(t, err)
	assert.NotNil(t, status)
	t.Log("API Status:", status)

	// Test GetResources
	params := map[string]string{"size": "20"}
	resources, err := c.GetResources(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, resources)
	t.Log("Resources:", resources)

	if len(resources) > 0 {
		resourceID := resources[len(resources)-1].ID

		// Test GetResourceByID
		resource, err := c.GetResourceByID(resourceID)
		assert.NoError(t, err)
		assert.NotNil(t, resource)
		t.Log("Resource:", resource)

		// Test GetResourceVersions
		versions, err := c.GetResourceVersions(resourceID, params)
		assert.NoError(t, err)
		assert.NotEmpty(t, versions)
		t.Log("Resource Versions:", versions)

		if len(versions) > 0 {
			versionID := versions[0].ID

			// Test GetResourceVersionByID
			version, err := c.GetResourceVersionByID(resourceID, versionID)
			assert.NoError(t, err)
			assert.NotNil(t, version)
			t.Log("Resource Version:", version)

			// Test GetLatestResourceVersion
			latestVersion, err := c.GetLatestResourceVersion(resourceID)
			assert.NoError(t, err)
			assert.NotNil(t, latestVersion)
			t.Log("Latest Resource Version:", latestVersion)

			// Test DownloadResourceVersion
			err = c.DownloadResourceVersion(*latestVersion, "./misc/test_download.jar", true)
			assert.NoError(t, err)
		}

		// Test GetResourceUpdates
		updates, err := c.GetResourceUpdates(resourceID, params)
		assert.NoError(t, err)
		assert.NotEmpty(t, updates)
		t.Log("Resource Updates:", updates)

		// Test GetResourceReviews
		reviews, err := c.GetResourceReviews(resourceID, params)
		assert.NoError(t, err)
		assert.NotEmpty(t, reviews)
		t.Log("Resource Reviews:", reviews)

		// Test GetResourceAuthor
		author, err := c.GetResourceAuthor(resourceID)
		assert.NoError(t, err)
		assert.NotNil(t, author)
		t.Log("Resource Author:", author)
	}

	// Test GetAuthors
	authors, err := c.GetAuthors(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, authors)
	t.Log("Authors:", authors)

	if len(authors) > 1 {
		authorID := authors[1].ID // The first author might be empty, so we use the second one

		// Test GetAuthorByID
		author, err := c.GetAuthorByID(authorID)
		assert.NoError(t, err)
		assert.NotNil(t, author)
		t.Log("Author:", author)

		// Test GetAuthorResources
		authorResources, err := c.GetAuthorResources(authorID, params)
		assert.NoError(t, err)
		assert.NotEmpty(t, authorResources)
		t.Log("Author Resources:", authorResources)

		// Test GetAuthorReviews
		authorReviews, err := c.GetAuthorReviews(authorID, params)
		assert.NoError(t, err)
		assert.NotEmpty(t, authorReviews)
		t.Log("Author Reviews:", authorReviews)
	}

	// Test GetCategories
	categories, err := c.GetCategories(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, categories)
	t.Log("Categories:", categories)

	if len(categories) > 0 {
		categoryID := categories[0].ID

		// Test GetCategoryByID
		category, err := c.GetCategoryByID(categoryID)
		assert.NoError(t, err)
		assert.NotNil(t, category)
		t.Log("Category:", category)

		// Test GetCategoryResources
		categoryResources, err := c.GetCategoryResources(categoryID, params)
		assert.NoError(t, err)
		assert.NotEmpty(t, categoryResources)
		t.Log("Category Resources:", categoryResources)
	}

	// Test SearchResources
	searchResources, err := c.SearchResources("plugin", params)
	assert.NoError(t, err)
	assert.NotEmpty(t, searchResources)
	t.Log("Search Resources:", searchResources)

	// Test SearchAuthors
	searchAuthors, err := c.SearchAuthors("author", params)
	assert.NoError(t, err)
	assert.NotEmpty(t, searchAuthors)
	t.Log("Search Authors:", searchAuthors)
}
