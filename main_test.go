package main

import(
	"testing"
	"log"
	"bufio"
	"bytes"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)
	data , _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!",string(data))
}

func TestDecoHandler(t *testing.T){
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	
	buffer := &bytes.Buffer{}
	log.SetOutput(buffer)

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK,res.StatusCode)

	read := bufio.NewReader(buffer)
	line,_ , err1 := read.ReadLine()
	assert.NoError(err1)
	assert.Contains(string(line),"[LOGGER1] Started")
	

	line,_ , err1 = read.ReadLine()
	assert.NoError(err1)
	assert.Contains(string(line),"[LOGGER1] Completed")
}
