package weak

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type mergeSuite struct {
	suite.Suite
	source map[interface{}]interface{}
	target map[interface{}]interface{}
}

func (s *mergeSuite) SetupTest() {
	s.source = map[interface{}]interface{}{
		"positiveMerge":          1,
		"nagativeMerge":          2,
		"positiveIgnoreEmpty":    0,
		"nagativeIgnoreEmpty":    2,
		"positiveIgnoreIncrease": 1,
		// "nagativeIgnoreIncrease": nil,
		"positiveOverwriteDiscored": 1,
		"nagativeOverwriteDiscored": 2,
	}
	s.target = map[interface{}]interface{}{
		"positiveMerge":       3,
		"nagativeMerge":       2,
		"positiveIgnoreEmpty": 3,
		"nagativeIgnoreEmpty": 0,
		// "positiveIgnoreIncrease":    nil,
		"nagativeIgnoreIncrease":    3,
		"positiveOverwriteDiscored": "3",
		"nagativeOverwriteDiscored": 3,
	}
}
func (s *mergeSuite) TestMerge() {
	merger := Merger{}
	assert.Nil(s.T(), merger.Merge(s.source, &s.target))
	assert.Equal(s.T(), s.source["positiveMerge"], 1)
	assert.Equal(s.T(), s.source["nagativeMerge"], 2)
	assert.Equal(s.T(), s.source["positiveIgnoreEmpty"], 0)
	assert.Equal(s.T(), s.source["nagativeIgnoreEmpty"], 2)
	assert.Equal(s.T(), s.source["positiveIgnoreIncrease"], 1)
	assert.Equal(s.T(), s.source["nagativeIgnoreIncrease"], nil)
	assert.Equal(s.T(), s.source["positiveOverwriteDiscored"], 1)
	assert.Equal(s.T(), s.source["nagativeOverwriteDiscored"], 2)
	assert.Equal(s.T(), s.target["positiveMerge"], 1)
	assert.Equal(s.T(), s.target["nagativeMerge"], 2)
	assert.Equal(s.T(), s.target["positiveIgnoreEmpty"], 0)
	assert.Equal(s.T(), s.target["nagativeIgnoreEmpty"], 2)
	assert.Equal(s.T(), s.target["positiveIgnoreIncrease"], 1)
	assert.Equal(s.T(), s.target["nagativeIgnoreIncrease"], 3)
	assert.Equal(s.T(), s.target["positiveOverwriteDiscored"], "3")
	assert.Equal(s.T(), s.target["nagativeOverwriteDiscored"], 2)
}
func (s *mergeSuite) TestMergeIgnoreEmpty() {
	merger := Merger{
		IgnoreEmpty: true,
	}
	assert.Nil(s.T(), merger.Merge(s.source, &s.target))
	assert.Equal(s.T(), s.source["positiveMerge"], 1)
	assert.Equal(s.T(), s.source["nagativeMerge"], 2)
	assert.Equal(s.T(), s.source["positiveIgnoreEmpty"], 0)
	assert.Equal(s.T(), s.source["nagativeIgnoreEmpty"], 2)
	assert.Equal(s.T(), s.source["positiveIgnoreIncrease"], 1)
	assert.Equal(s.T(), s.source["nagativeIgnoreIncrease"], nil)
	assert.Equal(s.T(), s.source["positiveOverwriteDiscored"], 1)
	assert.Equal(s.T(), s.source["nagativeOverwriteDiscored"], 2)
	assert.Equal(s.T(), s.target["positiveMerge"], 1)
	assert.Equal(s.T(), s.target["nagativeMerge"], 2)
	assert.Equal(s.T(), s.target["positiveIgnoreEmpty"], 3)
	assert.Equal(s.T(), s.target["nagativeIgnoreEmpty"], 2)
	assert.Equal(s.T(), s.target["positiveIgnoreIncrease"], 1)
	assert.Equal(s.T(), s.target["nagativeIgnoreIncrease"], 3)
	assert.Equal(s.T(), s.target["positiveOverwriteDiscored"], "3")
	assert.Equal(s.T(), s.target["nagativeOverwriteDiscored"], 2)
}
func (s *mergeSuite) TestMergeIgnoreIncrease() {
	merger := Merger{
		IgnoreIncrease: true,
	}
	assert.Nil(s.T(), merger.Merge(s.source, &s.target))
	assert.Equal(s.T(), s.source["positiveMerge"], 1)
	assert.Equal(s.T(), s.source["nagativeMerge"], 2)
	assert.Equal(s.T(), s.source["positiveIgnoreEmpty"], 0)
	assert.Equal(s.T(), s.source["nagativeIgnoreEmpty"], 2)
	assert.Equal(s.T(), s.source["positiveIgnoreIncrease"], 1)
	assert.Equal(s.T(), s.source["nagativeIgnoreIncrease"], nil)
	assert.Equal(s.T(), s.source["positiveOverwriteDiscored"], 1)
	assert.Equal(s.T(), s.source["nagativeOverwriteDiscored"], 2)
	assert.Equal(s.T(), s.target["positiveMerge"], 1)
	assert.Equal(s.T(), s.target["nagativeMerge"], 2)
	assert.Equal(s.T(), s.target["positiveIgnoreEmpty"], 0)
	assert.Equal(s.T(), s.target["nagativeIgnoreEmpty"], 2)
	assert.Equal(s.T(), s.target["positiveIgnoreIncrease"], nil)
	assert.Equal(s.T(), s.target["nagativeIgnoreIncrease"], 3)
	assert.Equal(s.T(), s.target["positiveOverwriteDiscored"], "3")
	assert.Equal(s.T(), s.target["nagativeOverwriteDiscored"], 2)
}
func (s *mergeSuite) TestMergeOverwriteDiscord() {
	merger := Merger{
		OverwriteDiscord: true,
	}
	assert.Nil(s.T(), merger.Merge(s.source, &s.target))
	assert.Equal(s.T(), s.source["positiveMerge"], 1)
	assert.Equal(s.T(), s.source["nagativeMerge"], 2)
	assert.Equal(s.T(), s.source["positiveIgnoreEmpty"], 0)
	assert.Equal(s.T(), s.source["nagativeIgnoreEmpty"], 2)
	assert.Equal(s.T(), s.source["positiveIgnoreIncrease"], 1)
	assert.Equal(s.T(), s.source["nagativeIgnoreIncrease"], nil)
	assert.Equal(s.T(), s.source["positiveOverwriteDiscored"], 1)
	assert.Equal(s.T(), s.source["nagativeOverwriteDiscored"], 2)
	assert.Equal(s.T(), s.target["positiveMerge"], 1)
	assert.Equal(s.T(), s.target["nagativeMerge"], 2)
	assert.Equal(s.T(), s.target["positiveIgnoreEmpty"], 0)
	assert.Equal(s.T(), s.target["nagativeIgnoreEmpty"], 2)
	assert.Equal(s.T(), s.target["positiveIgnoreIncrease"], 1)
	assert.Equal(s.T(), s.target["nagativeIgnoreIncrease"], 3)
	assert.Equal(s.T(), s.target["positiveOverwriteDiscored"], 1)
	assert.Equal(s.T(), s.target["nagativeOverwriteDiscored"], 2)
}

func TestMergeSuite(t *testing.T) {
	suite.Run(t, new(mergeSuite))
}
