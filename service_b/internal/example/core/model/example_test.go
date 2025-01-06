package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewExample(t *testing.T) {
	t.Run("it should be able to create a new Example", func(t *testing.T) {
		in := NewExampleInput{
			Name: "Example",
		}
		m := NewExample(in)

		assert.NotEmpty(t, m.ID)
		assert.Equal(t, in.Name, m.Name)
	})
}

func Test_ExampleMarshal(t *testing.T) {
	t.Run("it should be able to marshal Example", func(t *testing.T) {
		in := NewExampleInput{
			Name: "Example",
		}
		m := NewExample(in)

		data, err := m.MarshalJSON()

		assert.Nil(t, err)
		assert.NotNil(t, data)
	})
}
