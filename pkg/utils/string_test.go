package utils

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractContent(t *testing.T) {
	t1 := `{
		"media":{
			"image":{
				"source":"voc/image.jpg",
				"timestamp":"125695373218523",
				"file_key":"c6991f9e23500a38b55a6d6c7dcbe9f55429a5163311a7bf9344190960cb1fea",
				"file_ext":".jpg",
				"width":500
			}
		},
		"ground_truth":{
			"classification":{
				"category_id":1
			}
		}
	}`

	r1 := `{
				"source":"voc/image.jpg",
				"timestamp":"125695373218523",
				"file_key":"c6991f9e23500a38b55a6d6c7dcbe9f55429a5163311a7bf9344190960cb1fea",
				"file_ext":".jpg",
				"width":500
			}`

	t2 := `{
		"source":"voc/image.jpg",
		"timestamp":"125695373218523",
		"file_key"}`

	t3 := `{
		"source":"voc/image.jpg",
		"timestamp":"125695373218523",
		"file_key"`

	t4 := `{
		"media":{
			"image":{
				"source":"voc/image.jpg",
				"timestamp":"125695373218523",
				"file_key":"c6991f9e23500a38b55a6d6c7dcbe9f55429a5163311a7bf9344190960cb1fea",
				"file_ext":".jpg",
				"width":500, 
				"embed": {
					"age": 10
				}
			}
		}
	}`

	r4 := `{
				"source":"voc/image.jpg",
				"timestamp":"125695373218523",
				"file_key":"c6991f9e23500a38b55a6d6c7dcbe9f55429a5163311a7bf9344190960cb1fea",
				"file_ext":".jpg",
				"width":500, 
				"embed": {
					"age": 10
				}
			}`

	t5 := `[
				{"file_key": "a"},
				{"file_key": "b"},
				{"file_key": "c"}
	]`

	t6 := `[
				{"file_key": "a"},
				"file_key": "xxx",
				{"file_key": "b"},
				{"file_key": "xxx",
				{"file_key": "c"}
	]`

	t.Run("reg", func(t *testing.T) {
		reg := regexp.MustCompile("file_key")
		loc := reg.FindAllStringIndex(t5, -1)
		assert.Equal(t, 3, len(loc))
	})

	t.Run("ExtractContent", func(t *testing.T) {
		res := ExtractContent(t1, "file_key", "{", "}")
		assert.Equal(t, r1, res[0])

		res = ExtractContent(t2, "file_key", "{", "}")
		assert.Equal(t, t2, res[0])

		res = ExtractContent(t3, "file_key", "{", "}")
		assert.Equal(t, 0, len(res))

		res = ExtractContent(t4, "file_key", "{", "}")
		assert.Equal(t, r4, res[0])

		res = ExtractContent(t5, "file_key", "{", "}")
		assert.Equal(t, 3, len(res))

		res = ExtractContent(t6, "file_key", "{", "}")
		assert.Equal(t, 3, len(res))
		assert.Equal(t, res, []string{`{"file_key": "a"}`, `{"file_key": "b"}`, `{"file_key": "c"}`})
	})
}
