package server_test

import (
	"net/http/httptest"
	"testing"

	"github.com/foomo/gotsrpc/v2/tests/aliases/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultServiceGoTSRPCClient(t *testing.T) {
	t.Parallel()

	s := httptest.NewServer(server.NewDefaultServiceGoTSRPCProxy(&server.Handler{}))
	c := server.NewDefaultServiceGoTSRPCClient(s.URL)
	t.Cleanup(s.Close)

	t.Run("StatusValue", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.StatusValue(t.Context(), server.StatusActive)
		require.NoError(t, clientErr)
		assert.Equal(t, server.StatusActive, ret)
	})

	t.Run("CategoryValue", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.CategoryValue(t.Context(), server.CategoryA)
		require.NoError(t, clientErr)
		assert.Equal(t, server.CategoryA, ret)
	})

	t.Run("PriorityValue", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.PriorityValue(t.Context(), server.PriorityHigh)
		require.NoError(t, clientErr)
		assert.Equal(t, server.PriorityHigh, ret)
	})

	t.Run("RatingValue", func(t *testing.T) {
		t.Parallel()
		ret, clientErr := c.RatingValue(t.Context(), server.Rating(4.5))
		require.NoError(t, clientErr)
		assert.InDelta(t, 4.5, float64(ret), 1e-10)
	})

	t.Run("TagsValue", func(t *testing.T) {
		t.Parallel()
		v := server.Tags{"go", "typescript", "rpc"}
		ret, clientErr := c.TagsValue(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("EntriesValue", func(t *testing.T) {
		t.Parallel()
		e := &server.Entry{
			ID: "1", Status: server.StatusActive, Priority: server.PriorityHigh,
			Rating: 4.5, Tags: server.Tags{"a"},
		}
		v := server.Entries{e}
		ret, clientErr := c.EntriesValue(t.Context(), v)
		require.NoError(t, clientErr)
		require.Len(t, ret, 1)
		require.NotNil(t, ret[0])
		assert.Equal(t, *e, *ret[0])
	})

	t.Run("RegistryValue", func(t *testing.T) {
		t.Parallel()
		v := server.Registry{
			"first": {ID: "1", Status: server.StatusActive, Priority: server.PriorityLow, Rating: 1.0},
		}
		ret, clientErr := c.RegistryValue(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("IndexValue", func(t *testing.T) {
		t.Parallel()
		v := server.Index{
			server.CategoryA: {{ID: "1", Status: server.StatusActive, Priority: server.PriorityLow, Rating: 1.0}},
		}
		ret, clientErr := c.IndexValue(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("LabelMapValue", func(t *testing.T) {
		t.Parallel()
		v := server.LabelMap{"key": "val", "env": "prod"}
		ret, clientErr := c.LabelMapValue(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("EntryValue", func(t *testing.T) {
		t.Parallel()
		v := server.Entry{
			ID: "1", Status: server.StatusActive, Priority: server.PriorityHigh,
			Rating: 4.5, Tags: server.Tags{"a", "b"},
		}
		ret, clientErr := c.EntryValue(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("DetailValue", func(t *testing.T) {
		t.Parallel()
		v := server.Detail{
			Name:        "test",
			Description: "desc",
			Entry: server.Entry{
				ID: "1", Status: server.StatusActive, Priority: server.PriorityHigh, Rating: 4.5,
			},
			Labels: server.LabelMap{"key": "val"},
		}
		ret, clientErr := c.DetailValue(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("DataRecordValue", func(t *testing.T) {
		t.Parallel()
		note := "some note"
		v := server.DataRecord{
			ID:     "1",
			Title:  "test",
			Status: server.StatusPending,
			Amount: &server.Amount{Value: 100, Currency: "USD"},
			Items: server.Entries{
				{ID: "i1", Status: server.StatusActive, Priority: server.PriorityLow, Rating: 1.0},
			},
			Metadata:   &server.Metadata{CreatedBy: "admin", Note: &note},
			Categories: []server.Category{server.CategoryA, server.CategoryB},
		}
		ret, clientErr := c.DataRecordValue(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v.ID, ret.ID)
		assert.Equal(t, v.Title, ret.Title)
		assert.Equal(t, v.Status, ret.Status)
		require.NotNil(t, ret.Amount)
		assert.Equal(t, *v.Amount, *ret.Amount)
		require.Len(t, ret.Items, 1)
		require.NotNil(t, ret.Items[0])
		assert.Equal(t, *v.Items[0], *ret.Items[0])
		require.NotNil(t, ret.Metadata)
		assert.Equal(t, v.Metadata.CreatedBy, ret.Metadata.CreatedBy)
		require.NotNil(t, ret.Metadata.Note)
		assert.Equal(t, note, *ret.Metadata.Note)
		assert.Equal(t, v.Categories, ret.Categories)
	})

	t.Run("MapOfEntries", func(t *testing.T) {
		t.Parallel()
		v := map[server.Category][]server.Entry{
			server.CategoryA: {{ID: "1", Status: server.StatusActive, Priority: server.PriorityLow, Rating: 1.0}},
		}
		ret, clientErr := c.MapOfEntries(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v, ret)
	})

	t.Run("DataRecordNil", func(t *testing.T) {
		t.Parallel()
		v := server.DataRecord{
			ID:     "1",
			Title:  "minimal",
			Status: server.StatusClosed,
		}
		ret, clientErr := c.DataRecordNil(t.Context(), v)
		require.NoError(t, clientErr)
		assert.Equal(t, v.ID, ret.ID)
		assert.Equal(t, v.Status, ret.Status)
		assert.Nil(t, ret.Amount)
		assert.Nil(t, ret.Metadata)
		assert.Nil(t, ret.Items)
		assert.Nil(t, ret.Categories)
	})
}
