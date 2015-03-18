package gofakkuapiwrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)


const (
	CONTENT_TYPE_MANGA  string = "manga"
	CONTENT_TYPE_DOUJIN string = "doujin"

	CATEGORY_TYPE_NEWEST        string = "newest"
	CATEGORY_TYPE_JAPANESE      string = "japanese"
	CATEGORY_TYPE_ENGLISH       string = "english"
	CATEGORY_TYPE_FAVORITES     string = "favorites"
	CATEGORY_TYPE_POPULAR       string = "popular"
	CATEGORY_TYPE_CONTROVERSIAL string = "controversial"
)

func FakkuPageURL(contentType string, categoryType string, n uint64) *url.URL {
	url, _ := url.Parse("https://api.fakku.net/" + contentType + "/" + categoryType + "/page/" + strconv.FormatUint(n, 10))
	return url
}

func GetFakkuPage(contentType string, categoryType string, n uint64) (FakkuPage, error) {
	var page FakkuPage

	url := FakkuPageURL(contentType, categoryType, n)
	jsonString, loadErr := loadJSON(url)

	if loadErr != nil {
		return page, loadErr
	}
	jsonDecoder := json.NewDecoder(strings.NewReader(jsonString))

	decodeErr := jsonDecoder.Decode(&page)
	if decodeErr != nil && decodeErr != io.EOF {
		return page, decodeErr
	}

	if len(page.Error) > 0 {
		return page, fmt.Errorf(page.Error)
	}

	return page, nil
}

func FakkuDescriptionURL(contentType string, urlContentName string) *url.URL {
	url, _ := url.Parse("https://api.fakku.net/" + contentType + "/" + urlContentName)
	return url
}

func GetFakkuDescription(contentType string, urlContentName string) (FakkuDescription, error) {
	var desc FakkuDescription

	url := FakkuDescriptionURL(contentType, urlContentName)
	jsonString, loadErr := loadJSON(url)

	if loadErr != nil {
		return desc, loadErr
	}
	jsonDecoder := json.NewDecoder(strings.NewReader(jsonString))

	decodeErr := jsonDecoder.Decode(&desc)
	if decodeErr != nil && decodeErr != io.EOF {
		return desc, decodeErr
	}

	if len(desc.Error) > 0 {
		return desc, fmt.Errorf(desc.Error)
	}

	return desc, nil
}

func FakkuReadContentURL(contentType string, urlContentName string) *url.URL {
	url, _ := url.Parse("https://api.fakku.net/" + contentType + "/" + urlContentName + "/read")
	return url
}

func GetFakkuReadContent(contentType string, urlContentName string) (FakkuReadContent, error) {
	var read FakkuReadContent

	url := FakkuReadContentURL(contentType, urlContentName)
	jsonString, loadErr := loadJSON(url)

	if loadErr != nil {
		return read, loadErr
	}
	jsonDecoder := json.NewDecoder(strings.NewReader(jsonString))

	decodeErr := jsonDecoder.Decode(&read)
	if decodeErr != nil && decodeErr != io.EOF {
		return read, decodeErr
	}

	if len(read.Error) > 0 {
		return read, fmt.Errorf(read.Error)
	}

	return read, nil
}
