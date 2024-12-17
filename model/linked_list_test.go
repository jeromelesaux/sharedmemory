package model_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jeromelesaux/sharedmemory/model"
	"github.com/stretchr/testify/assert"
)

func TestJsonLinkedList(t *testing.T) {
	t.Run("ok_one_element", func(t *testing.T) {
		l := model.NewLinkedList()
		l.Add("hello world")

		content, err := json.Marshal(l)
		assert.NoError(t, err)
		fmt.Println(string(content))
	})

	t.Run("ok_2_elements", func(t *testing.T) {
		l := model.NewLinkedList()
		l.Add("hello world")
		l.Add("how are you?")

		content, err := json.Marshal(l)
		assert.NoError(t, err)
		fmt.Println(string(content))
	})

	t.Run("ok_umarshal_one", func(t *testing.T) {
		l := model.NewLinkedList()
		l.Add("hello world")

		content, err := json.Marshal(l)
		assert.NoError(t, err)

		l2 := model.NewLinkedList()
		err = json.Unmarshal(content, &l2)
		assert.NoError(t, err)
		assert.Equal(t, l, l2)
	})

	t.Run("ok_umarshal_two", func(t *testing.T) {
		l := model.NewLinkedList()
		l.Add("hello world")
		l.Add("how are you?")

		content, err := json.Marshal(l)
		assert.NoError(t, err)

		l2 := model.NewLinkedList()
		err = json.Unmarshal(content, &l2)
		assert.NoError(t, err)
		assert.Equal(t, l, l2)
	})
}
