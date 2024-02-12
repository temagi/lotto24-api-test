package tests

import (
	"encoding/json"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	t.Run("Search for page containing exactly 'furry rabbits' should returns empty results", func(t *testing.T) {
		uri := "search/page?q=\"furry%20rabbits\""

		resp := makeAPIRequest(uri)

		assert.Equal(t, OK, resp.StatusCode())

		var result SearchContentResponse
		err := json.Unmarshal([]byte(resp.Body()), &result)
		require.NoError(t, err)

		assert.Equal(t, SearchContentResponse{Pages: []SearchResultObject{}}, result)
	})

	t.Run("Search for page containing 'furry rabbits' should returns results containing page with title 'Sesame Street'", func(t *testing.T) {
		uri := "search/page?q=furry%20rabbits&limit=100"

		resp := makeAPIRequest(uri)

		assert.Equal(t, OK, resp.StatusCode())

		var result SearchContentResponse
		err := json.Unmarshal([]byte(resp.Body()), &result)
		require.NoError(t, err)

		assert.Equal(t, 100, len(result.Pages))
		var found bool
		for _, page := range result.Pages {
			if page.Title == "Sesame Street" {
				found = true
				break
			}
		}
		assert.True(t, found, "Page not found in results")
	})

	t.Run("'Sesame Street' page timestamp should be > 2023-08-17", func(t *testing.T) {
		uri := "search/page?q=furry%20rabbits&limit=100"

		resp := makeAPIRequest(uri)

		assert.Equal(t, OK, resp.StatusCode())

		var searchResults SearchContentResponse
		err := json.Unmarshal([]byte(resp.Body()), &searchResults)
		require.NoError(t, err)
		// Just to demonstrate more modern go functional approach
		resultPosition := slices.IndexFunc(searchResults.Pages, func(p SearchResultObject) bool { return p.Title == "Sesame Street" })
		assert.NotEqual(t, -1, resultPosition, "Page not found in results")

		pageKey := searchResults.Pages[resultPosition].Key

		uri = fmt.Sprintf("page/%s/bare", pageKey)
		resp = makeAPIRequest(uri)

		assert.Equal(t, OK, resp.StatusCode())

		var pageDetails Page
		err = json.Unmarshal([]byte(resp.Body()), &pageDetails)
		require.NoError(t, err)

		latest, err := time.Parse(time.RFC3339, pageDetails.Latest.Timestamp)
		require.NoError(t, err)
		
		expected, err := time.Parse("2006-01-02", "2023-08-17")
		require.NoError(t, err)
		
		assert.Greater(t, latest, expected)
	})
}
