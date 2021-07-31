package forismatic_api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	forismatic "github.com/ifedor/forismatic-api"
	"github.com/stretchr/testify/assert"
)

func TestQuoteBadResponse(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer s.Close()

	forismatic.SetBaseURL(s.URL)
	r, err := forismatic.Quote(forismatic.XmlFormat)
	assert.Error(t, err)
	assert.Nil(t, r)
}

func TestQuote(t *testing.T) {
	quote := "<forismatic><quote><quoteText>People are like stained glass windows: they sparkle and shine when the sun is out, but when the darkness sets in their true beauty is revealed only if there is a light within. </quoteText><quoteAuthor>Elizabeth Kubler-Ross</quoteAuthor><senderName></senderName><senderLink></senderLink><quoteLink>http://forismatic.com/en/1261fad063/</quoteLink></quote></forismatic>"
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(quote))
	}))
	defer s.Close()

	forismatic.SetBaseURL(s.URL)
	r, err := forismatic.Quote(forismatic.XmlFormat)
	assert.NoError(t, err)
	assert.Equal(t, quote, string(r))
}

func TestQuoteJson(t *testing.T) {
	quote := `{"quoteText":"Everything that irritates us about others can lead us to an understanding of ourselves. ", "quoteAuthor":"Carl Jung", "senderName":"", "senderLink":"", "quoteLink":"http://forismatic.com/en/ff8766e47b/"}`

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(quote))
	}))
	defer s.Close()

	forismatic.SetBaseURL(s.URL)
	q, err := forismatic.QuoteJson()
	assert.NoError(t, err)
	if assert.NotNil(t, q) {
		assert.Equal(t, "Everything that irritates us about others can lead us to an understanding of ourselves. ", q.Text)
		assert.Equal(t, "Carl Jung", q.Author)
		assert.Equal(t, "", q.SenderName)
		assert.Equal(t, "", q.SenderLink)
	}
}
