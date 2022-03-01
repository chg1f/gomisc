package weak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type flatSuite struct {
	suite.Suite
	input  interface{}
	output map[string]interface{}
}

func (s *flatSuite) SetupTest() {
	type Temp struct {
		JsonTag          interface{} `json:"json_tag"`
		JsonOmitemptyTag interface{} `json:"json_omitempty_tag,omitempty"`
		XMLTag           interface{} `xml:"xml_tag"`
		NotExistTag      interface{}
	}
	f := &Flater{
		StructTag:    "json",
		UintptrValue: true,
		FlatMap:      true,
		FlatStruct:   true,
		FlatArray:    true,
		FlatSlice:    true,
	}
	s.input = map[interface{}]interface{}{
		"a": "a",
		1:   1,
		"list": []interface{}{
			"a",
			1,
		},
		"map": map[interface{}]interface{}{
			"a": "a",
			1:   1,
		},
		"struct": Temp{
			JsonTag:          "x",
			JsonOmitemptyTag: "x",
			XMLTag:           "x",
			NotExistTag:      "x",
		},
		"listmap": []interface{}{
			"a",
			1,
			map[interface{}]interface{}{
				"a": "a",
				1:   1,
			},
		},
		"maplist": map[interface{}]interface{}{
			"a": "a",
			1:   1,
			"list": []interface{}{
				"a",
				1,
			},
		},
	}
	s.output = f.Flat(s.input)
}
func (s *flatSuite) TestFlat() {
	assert.Equal(s.T(), "a", s.output["a"])
	assert.Equal(s.T(), 1, s.output["1"])
	assert.Equal(s.T(), "a", s.output["list.[0]"])
	assert.Equal(s.T(), 1, s.output["list.[1]"])
	assert.Equal(s.T(), "a", s.output["map.a"])
	assert.Equal(s.T(), 1, s.output["map.1"])
	assert.Equal(s.T(), "x", s.output["struct.json_tag"])
	assert.Equal(s.T(), "x", s.output["struct.json_omitempty_tag"])
	assert.Equal(s.T(), "x", s.output["struct.XMLTag"])
	assert.Equal(s.T(), "x", s.output["struct.NotExistTag"])
	assert.Equal(s.T(), "a", s.output["listmap.[0]"])
	assert.Equal(s.T(), 1, s.output["listmap.[1]"])
	assert.Equal(s.T(), "a", s.output["listmap.[2]a"])
	assert.Equal(s.T(), 1, s.output["listmap.[2]1"])
	assert.Equal(s.T(), "a", s.output["maplist.a"])
	assert.Equal(s.T(), 1, s.output["maplist.1"])
	assert.Equal(s.T(), "a", s.output["maplist.list.[0]"])
	assert.Equal(s.T(), 1, s.output["maplist.list.[1]"])
	fmt.Print(s.output)
}

func TestFlatSuite(t *testing.T) {
	suite.Run(t, new(flatSuite))
}
