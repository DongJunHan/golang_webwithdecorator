package main

import(
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)
}

