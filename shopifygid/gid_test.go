package shopifygid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToGID(t *testing.T) {
	tests := []struct {
		id       int64
		domain   string
		resource string
		expected string
	}{
		{123123, "shopify", "Collection", "gid://shopify/Collection/123123"},
		{0, "shopify", "Collection", ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			assert.Equal(t, test.expected, ToGID(test.id, test.domain, test.resource))
		})
	}
}

func TestToShopifyGID(t *testing.T) {
	tests := []struct {
		id       int64
		resource string
		expected string
	}{
		{123123, "Collection", "gid://shopify/Collection/123123"},
		{0, "Collection", ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			assert.Equal(t, test.expected, ToShopifyGID(test.id, test.resource))
		})
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		gid      string
		expected int64
	}{
		{"", 0},
		{"gid://shopify/Collection/", 0},
		{"gid://shopify/Collection/123123", 123123},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			assert.Equal(t, test.expected, ToInt64(test.gid))
		})
	}
}
