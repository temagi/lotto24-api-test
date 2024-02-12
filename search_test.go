package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	const baseURL = "https://api.wikimedia.org/core/v1/wikipedia/en"

	t.Run("Search for page containing 'furry rabbits' should returns empty results", func(t *testing.T) {
		uri := "search/page?q=\"furry%20rabbits\""
		
		resp := makeAPIRequest(baseURL, uri)

		assert.Equal(t, 200, resp.StatusCode())

		var result PagesResponse
		err := json.Unmarshal([]byte(resp.Body()), &result)
		log.Println(result)
		assert.Equal(t, PagesResponse{Pages: []Page{}}, result)
		require.NoError(t, err)
	})
}