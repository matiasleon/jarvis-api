package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJarvisContextBuilder_Append_should_return_appended_string(t *testing.T) {

	// init
	javisContextBuilder := NewJarvisContextBuilder("context")

	// execute
	javisContextBuilder.Append("hola")
	contextString := javisContextBuilder.Build()

	// assert
	assert.EqualValues(t, "context. hola", contextString)
}

func TestJarvisContextBuilder_Append_should_return_appended_string_two_times(t *testing.T) {

	// init
	javisContextBuilder := NewJarvisContextBuilder("context")

	// execute
	javisContextBuilder.Append("hola")
	javisContextBuilder.Append("hola 2")
	contextString := javisContextBuilder.Build()

	// assert
	assert.EqualValues(t, "context. hola. hola 2", contextString)
}
