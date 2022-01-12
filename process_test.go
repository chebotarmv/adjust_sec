package main

import (
	"adjust/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MD5_Hash(t *testing.T) {
	t.Run("check MD5 hashing", func(t *testing.T) {
		service := service.New()
		_, err := service.GetHash("www.google.com")
		assert.NoError(t, err)
	})
}
