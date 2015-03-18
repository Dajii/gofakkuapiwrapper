package gofakkuapiwrapper

import (
	"net/url"
	"strings"
	"strconv"
	"fmt"
)


type FakkuTag struct {
	Attribute      string
	Attribute_link string
	Attribute_id   string
}

type FakkuTranslator struct {
	Attribute      string
	Attribute_link string
	Attribute_id   string
}

type FakkuSerie struct {
	Attribute      string
	Attribute_link string
	Attribute_id   string
}

type FakkuArtist struct {
	Attribute      string
	Attribute_link string
	Attribute_id   string
}

type FakkuContentImages struct {
	Cover  string
	Sample string
}

type FakkuContent struct {
	Content_name        string
	Content_url         string
	Content_description string
	Content_language    string
	Content_category    string
	Content_date        uint64
	Content_filesize    uint64
	Content_favorites   uint64
	Content_comments    uint64
	Content_pages       uint64
	Content_poster      string
	Content_poster_url  string
	Content_tags        []FakkuTag
	Content_translators []FakkuTranslator
	Content_series      []FakkuSerie
	Content_artists     []FakkuArtist
}

func (f *FakkuContent) Content_url_path() string {
	url, _ := url.Parse(f.Content_url)
	return url.Path
}

func (f *FakkuContent) Content_url_name() string {
	sp := strings.Split(f.Content_url, "/")
	return sp[len(sp)-1]
}

type FakkuPage struct {
	Content []FakkuContent
	Total   uint64
	Page    string
	Pages   uint64
	Error   string
}

type FakkuDescription struct {
	Content FakkuContent
	Error   string
}

type FakkuReadContentImage struct {
	Thumb string
	Image string
}

type FakkuReadContent struct {
	Content FakkuContent
	Pages   map[string]FakkuReadContentImage
	Error   string
}

func (f * FakkuReadContent) Image(n uint64) (FakkuReadContentImage,error) {
	image, found := f.Pages[strconv.FormatUint(n, 10)]
	if !found {
		return image, fmt.Errorf("No images found for n = %d", n)
	}
	return image, nil
}
