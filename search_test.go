package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	const baseURL = "https://api.wikimedia.org/core/v1/wikipedia/en"

	t.Run("Search for page containing exactly 'furry rabbits' should returns empty results", func(t *testing.T) {
		uri := "search/page?q=\"furry%20rabbits\""

		resp := makeAPIRequest(baseURL, uri)

		assert.Equal(t, 200, resp.StatusCode())

		var result SearchContentResponse
		err := json.Unmarshal([]byte(resp.Body()), &result)
		require.NoError(t, err)

		assert.Equal(t, SearchContentResponse{Pages: []SearchResultObject{}}, result)
	})

	t.Run("Search for page containing 'furry rabbits' should returns results containing page with title 'Sesame Street'", func(t *testing.T) {
		uri := "search/page?q=furry%20rabbits&limit=100"

		resp := makeAPIRequest(baseURL, uri)

		assert.Equal(t, 200, resp.StatusCode())

		var result SearchContentResponse
		err := json.Unmarshal([]byte(resp.Body()), &result)
		require.NoError(t, err)

		assert.Equal(t, 100, len(result.Pages))
		resultPosition := slices.IndexFunc(result.Pages, func(p SearchResultObject) bool { return p.Title == "Sesame Street" })
		assert.NotEqual(t, -1, resultPosition, "Page not found in results")
	})

	t.Run("'Sesame Street' page timestamp should be > 2023-08-17", func(t *testing.T) {
		uri := "search/page?q=furry%20rabbits&limit=100"

		resp := makeAPIRequest(baseURL, uri)

		assert.Equal(t, 200, resp.StatusCode())

		var searchResults SearchContentResponse
		err := json.Unmarshal([]byte(resp.Body()), &searchResults)
		require.NoError(t, err)
		resultPosition := slices.IndexFunc(searchResults.Pages, func(p SearchResultObject) bool { return p.Title == "Sesame Street" })
		assert.NotEqual(t, -1, resultPosition, "Page not found in results")

		pageKey := searchResults.Pages[resultPosition].Key

		uri = fmt.Sprintf("page/%s/bare", pageKey)
		resp = makeAPIRequest(baseURL, uri)

		assert.Equal(t, 200, resp.StatusCode())

		var pageDetails Page
		err = json.Unmarshal([]byte(resp.Body()), &pageDetails)
		require.NoError(t, err)

		latest, err := time.Parse(time.RFC3339, pageDetails.Latest.Timestamp)
		require.NoError(t, err)
		log.Println(latest)
		expected, err := time.Parse("2006-01-02", "2023-08-17")
		require.NoError(t, err)
		log.Println(expected)
		require.NoError(t, err)
		assert.Greater(t, latest, expected)
	})
}
