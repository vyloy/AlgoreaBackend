package types

import (
	"encoding/json"
	"testing"

	assert_lib "github.com/stretchr/testify/assert"
)

type SampleIntInput struct {
	ID       RequiredInt64
	ChildID  NullableInt64
	Order    OptionalInt64
	ParentID OptNullInt64
}

func (v *SampleIntInput) validate() error {
	return Validate(&v.ID, &v.ChildID, &v.Order, &v.ParentID)
}

func TestIntValid(t *testing.T) {
	assert := assert_lib.New(t)

	jsonInput := `{ "ID": 2147483645, "ChildID": 22, "Order": -1, "ParentID": 7 }`
	input := &SampleIntInput{}
	assert.NoError(json.Unmarshal([]byte(jsonInput), &input))
	assert.EqualValues(2147483645, input.ID.Value)
	assert.EqualValues(22, input.ChildID.Value)
	assert.EqualValues(-1, input.Order.Value)
	assert.EqualValues(7, input.ParentID.Value)
	assert.NoError(input.validate())
}

func TestIntWithNonInt(t *testing.T) {
	assert := assert_lib.New(t)

	jsonInput := `{ "ID": "not an int", "ChildID": 22, "Order": -1, "ParentID": 7 }`
	input := &SampleIntInput{}
	assert.Error(json.Unmarshal([]byte(jsonInput), &input))
}

func TestIntWithDefault(t *testing.T) {
	assert := assert_lib.New(t)

	jsonInput := `{ "ID": 0, "ChildID": 0, "Order": 0, "ParentID": 0 }`
	input := &SampleIntInput{}
	assert.NoError(json.Unmarshal([]byte(jsonInput), &input))
	assert.NoError(input.validate())
}

func TestIntWithNull(t *testing.T) {
	assert := assert_lib.New(t)

	jsonInput := `{ "ID": null, "ChildID": null, "Order": null, "ParentID": null }`
	input := &SampleIntInput{}
	assert.NoError(json.Unmarshal([]byte(jsonInput), &input))
	assert.Error(input.ID.Validate(), "was expecting a validation error")
	assert.NoError(input.ChildID.Validate())  // should be valid
	assert.Error(input.Order.Validate())      // should NOT be valid
	assert.NoError(input.ParentID.Validate()) // should be valid
	assert.Error(input.validate())
}

func TestIntWithNotSet(t *testing.T) {
	assert := assert_lib.New(t)

	jsonInput := `{}`
	input := &SampleIntInput{}
	assert.NoError(json.Unmarshal([]byte(jsonInput), &input))
	assert.Error(input.ID.Validate())         // should NOT be valid
	assert.Error(input.ChildID.Validate())    // should NOT be valid
	assert.NoError(input.Order.Validate())    // should be valid
	assert.NoError(input.ParentID.Validate()) // should be valid
	assert.Error(input.validate())
}
