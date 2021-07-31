package forismatic_api_test

import (
	"testing"

	forismatic "github.com/ifedor/forismatic-api"
	"github.com/stretchr/testify/assert"
)

func TestSetLang(t *testing.T) {
	assert.Equal(t, forismatic.LangEn, forismatic.Lang()) // check default value
	forismatic.SetLang(forismatic.LangRu)
	assert.Equal(t, forismatic.LangRu, forismatic.Lang())
}
