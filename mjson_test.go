package mjson_benchmarks

import (
	"github.com/haihuxia/mjson"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var json = `{
  "text_matches": [
    {
      "object_url": "https://api.github.com/",
      "object_type": "Issue",
      "property": "body",
      "fragment": "comprehensive windows font I know of).\n\nIf we can find a commonly distributed windows .\n",
      "matches": [
        {
          "text": "windows",
          "indices": [
            14,
            21
          ]
        },
        {
          "text": "windows",
          "indices": [
            78,
            85
          ]
        }
      ]
    },
    {
      "object_url": "https://api.github.com",
      "object_type": "IssueComment",
      "property": "body",
      "fragment": " right after that are a bit broken IMHO :). \n",
      "matches": [
        {
          "text": "Windows",
          "indices": [
            163,
            170
          ]
        }
      ]
    }
  ]
}`

var mappingMap = map[string]string{"text_matches.object_url": "o_url", "text_matches.object_type": "o_type",
	"text_matches.property": "o_property"}

func readYAML(filePath string) (spec *mjson.Spec, ok bool) {
	buff, err := ioutil.ReadFile(filePath)
	if err != nil {
		return spec, false
	}
	err = yaml.Unmarshal(buff, &spec)
	if err != nil {
		return spec, false
	}
	return spec, true
}

func TestMappingSpec(t *testing.T) {
	spec, ok := readYAML("mapping.yaml")
	if !ok {
		return
	}
	t.Log(mjson.MappingSpec(json, spec))
}

func TestMappingString(t *testing.T) {
	for key, val := range mappingMap {
		t.Log(mjson.MappingString(json, key, val))
	}
}

func BenchmarkMappingString(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mjson.MappingString(json, "text_matches.object_type", "o_type")
	}
}

func BenchmarkMappingSpec(t *testing.B) {
	spec, ok := readYAML("mapping.yaml")
	if !ok {
		return
	}
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mjson.MappingSpec(json, spec)
	}
}

func BenchmarkMappingMap(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for key, val := range mappingMap {
			mjson.MappingString(json, key, val)
		}
	}
}
