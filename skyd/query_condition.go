package skyd

import (
	"errors"
	"fmt"
)

//------------------------------------------------------------------------------
//
// Constants
//
//------------------------------------------------------------------------------

const (
	QueryConditionUnitSteps    = "steps"
	QueryConditionUnitSessions = "sessions"
	QueryConditionUnitSeconds  = "seconds"
)

//------------------------------------------------------------------------------
//
// Typedefs
//
//------------------------------------------------------------------------------

// A condition step made within a query.
type QueryCondition struct {
	query       *Query
	Within      int
	WithinUnits string
	Steps       QueryStepList
}

//------------------------------------------------------------------------------
//
// Constructors
//
//------------------------------------------------------------------------------

// Creates a new condition.
func NewQueryCondition(query *Query) *QueryCondition {
	return &QueryCondition{
		query:       query,
		Within:      0,
		WithinUnits: QueryConditionUnitSteps,
	}
}

//------------------------------------------------------------------------------
//
// Accessors
//
//------------------------------------------------------------------------------

// Retrieves the query this condition is associated with.
func (c *QueryCondition) Query() *Query {
	return c.query
}

//------------------------------------------------------------------------------
//
// Methods
//
//------------------------------------------------------------------------------

//--------------------------------------
// Serialization
//--------------------------------------

// Encodes a query condition into an untyped map.
func (c *QueryCondition) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"type":        QueryStepTypeCondition,
		"within":      c.Within,
		"withinUnits": c.WithinUnits,
		"steps":       c.Steps.Serialize(),
	}
}

// Decodes a query condition from an untyped map.
func (c *QueryCondition) Deserialize(obj map[string]interface{}) error {
	if obj == nil {
		return errors.New("skyd.QueryCondition: Unable to deserialize nil.")
	}
	if obj["type"] != QueryStepTypeCondition {
		return fmt.Errorf("skyd.QueryCondition: Invalid step type: %v", obj["type"])
	}
	
	// Deserialize "within".
	if within, ok := obj["within"].(float64); ok {
		c.Within = int(within)
	} else {
		return fmt.Errorf("Invalid 'within': %v", obj["within"])
	}

	// Deserialize "within units".
	if withinUnits, ok := obj["withinUnits"].(string); ok {
		switch withinUnits {
		case QueryConditionUnitSteps, QueryConditionUnitSessions, QueryConditionUnitSeconds:
			c.WithinUnits = withinUnits
		default:
			return fmt.Errorf("Invalid 'within units': %v", withinUnits)
		}
	} else {
		return fmt.Errorf("Invalid 'within units': %v", obj["within"])
	}

	// Deserialize steps.
	var err error
	c.Steps, err = DeserializeQueryStepList(obj["steps"], c.query)
	if err != nil {
		return err
	}

	return nil
}

//--------------------------------------
// Code Generation
//--------------------------------------

// Generates Lua code for the query.
func (c *QueryCondition) Codegen() (string, error) {
	return "", nil
}