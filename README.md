# Example/Usage

```go
package main

import (
	wrapper "github.com/dajii/gofakkuapiwrapper"
	"fmt"
)


func main() {
	p, err := wrapper.GetFakkuPage(wrapper.CONTENT_TYPE_MANGA, wrapper.CATEGORY_TYPE_NEWEST, 12)
	fmt.Printf("Err: %v\n", err)
	fmt.Printf("P: %v\n", p)

	d,err2 := wrapper.GetFakkuDescription(wrapper.CONTENT_TYPE_MANGA, p.Content[0].Content_url_name())
	fmt.Printf("Err2: %v\n", err2)
	fmt.Printf("D: %v\n", d)

	v,err3 := wrapper.GetFakkuReadContent(wrapper.CONTENT_TYPE_MANGA, p.Content[0].Content_url_name())
	fmt.Printf("Err3: %v\n", err3)
	fmt.Printf("V: %v\n", v)
}
```
