package ez_log

import (
	"os"
	"strings"
)

type CategoryFilter interface {
	Enable(category *Category) bool
}

var filter CategoryFilter

type categoryFilter Category
type nopFilter struct {}

func (n nopFilter) Enable(_ *Category) bool {
	return true
}

func init() {
	debug := os.Getenv("DEBUG")
	if debug == "" {
		filter = &nopFilter{}
	}
	debugs := strings.SplitN(debug, ":", 5)
	for i, v := range debugs {
		if v == "*" {
			debugs[i] = ""
		}
	}
	debugs = debugs[0:5]
	filter = &categoryFilter{debugs[0], debugs[1], debugs[2], debugs[3], debugs[4]}
}

func (filter *categoryFilter) Enable(category *Category) bool {
	if filter.Module != "" && filter.Module != category.Module {
		return false
	}
	if filter.Category != "" && filter.Category != category.Category {
		return false
	}
	if filter.SubCategory != "" && filter.SubCategory != category.SubCategory {
		return false
	}
	if filter.Filter1 != "" && filter.Filter1 != category.Filter1 {
		return false
	}
	if filter.Filter2 != "" && filter.Filter2 != category.Filter2 {
		return false
	}
	return true
}
