package main

import (
	"fmt"
	matcher "matcher/src"
)

func main()  {
  fmt.Print(matcher.Routes["test"].NeedsAuth)
}
