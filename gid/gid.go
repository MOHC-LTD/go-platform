package gid

import (
	"fmt"
	"strconv"
	"strings"
)

// ToGID converts an integer to a GID resource ID
func ToGID(id int64, domain string, resourceName string) string {
	if id == 0 {
		return ""
	}
	return fmt.Sprintf("gid://%s/%s/%v", domain, resourceName, id)
}

// ToShopifyGID converts an integer to a Shopify GID resource ID
func ToShopifyGID(id int64, resourceName string) string {
	return ToGID(id, "shopify", resourceName)
}

// ToInt64 converts a Shopify GQL resource ID to an integer
func ToInt64(gid string) int64 {
	components := strings.Split(gid, "/")
	idString := components[len(components)-1]

	id, _ := strconv.ParseInt(idString, 10, 64)
	return id
}
