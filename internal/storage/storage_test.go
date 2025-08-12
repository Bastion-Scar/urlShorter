package storage_test

import (
	"awesomeProject13/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapStorage_SaveAndLoad(t *testing.T) {
	st := storage.NewMapStorage()
	code := "testCode"
	url := "https://google.com"

	st.Save(code, url)

	got, ok := st.Load(code)
	assert.True(t, ok)
	assert.Equal(t, url, got)

	_, ok = st.Load("DimaLox")
	assert.False(t, ok)
}
