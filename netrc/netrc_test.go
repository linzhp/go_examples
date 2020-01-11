package netrc

import (
	bgentry "github.com/bgentry/go-netrc/netrc"
	jdxcode "github.com/jdxcode/netrc"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var netrcPath = "testdata/.netrc"

func TestBgentryReadWrite(t *testing.T) {
	oldContent, err := ioutil.ReadFile(netrcPath)
	assert.NoError(t, err)
	net, err := bgentry.ParseFile(netrcPath)
	assert.NoError(t, err)
	newContent, err := net.MarshalText()
	assert.Equal(t, string(oldContent), string(newContent))
}

func TestJdxcodeReadWrite(t *testing.T)  {
	oldContent, err := ioutil.ReadFile(netrcPath)
	assert.NoError(t, err)
	net, err := jdxcode.Parse(netrcPath)
	assert.NoError(t, err)
	assert.Equal(t, string(oldContent), net.Render())
}