package forismatic_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type responseFormat string

const (
	XmlFormat  responseFormat = "xml"
	JsonFormat responseFormat = "json"
	HtmlFormat responseFormat = "html"
	TextFormat responseFormat = "text"
)

type (
	getQuoteResponse struct {
		Text       string `json:"quoteText"`
		Author     string `json:"quoteAuthor"`
		SenderName string `json:"senderName"`
		SenderLink string `json:"senderLink"`
		QuoteLink  string `json:"quoteLink"`
	}
)

var baseApiUri = "http://api.forismatic.com/api/1.0"

func SetBaseURL(s string) {
	baseApiUri = s
}
func QuoteJson() (*getQuoteResponse, error) {
	bytes, err := Quote(JsonFormat)
	if err != nil {
		return nil, err
	}
	var q getQuoteResponse
	if err := json.Unmarshal(bytes, &q); err != nil {
		return nil, err
	}
	return &q, nil
}

func Quote(f responseFormat) ([]byte, error) {
	url := fmt.Sprintf("%s/?format=%s&method=%s&lang=%s", baseApiUri, f, "getQuote", currentLanguageOption)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("forismatic bad api response: status - %d; body - %s",
			resp.StatusCode, string(body)))
	}

	return body, nil
}
