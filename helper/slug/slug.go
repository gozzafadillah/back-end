package slug

import (
	"github.com/gosimple/slug"
)

func GenerateSlug(name string) string {
	text := slug.Make(name)
	return text
}
