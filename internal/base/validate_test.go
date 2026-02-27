package base_test

import (
	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
	"testing"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("When ValidateRequest has empty fields - true", func(t *testing.T) {
		t.Parallel()

		in := new(base.ValidateRequest)
		rsl := []string{
			"ValidateRequest userID is empty!",
			"ValidateRequest title is empty!",
			"ValidateRequest description is empty!",
		}
		result := base.Validate(in)

		assert.Equal(t, rsl, result)
	})

	t.Run("When ValidateRequest has is nil - true", func(t *testing.T) {
		t.Parallel()

		rsl := []string{
			"ValidateRequest is nil!",
		}
		result := base.Validate(nil)

		assert.Equal(t, rsl, result)
	})

	t.Run("When ValidateRequest has empty field UserId - true", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{UserID: "", Title: "some Title", Description: "Some description"}
		rsl := []string{
			"ValidateRequest userID is empty!",
		}
		result := base.Validate(in)

		assert.Equal(t, rsl, result)
	})

	t.Run("When ValidateRequest has empty field Description - true", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{UserID: "6534-234-2345-hj45", Title: "some Title", Description: ""}
		rsl := []string{
			"ValidateRequest description is empty!",
		}
		result := base.Validate(in)

		assert.Equal(t, rsl, result)
	})
}
