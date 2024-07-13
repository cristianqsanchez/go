package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
  Title       string
  Description string
  Tags        []string
  Body        string
}

const (
  titleKey       = "Title: "
  descriptionKey = "Description: "
  tagsKey        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
  scanner := bufio.NewScanner(postFile)

  getMetaValue := func(key string) string {
    scanner.Scan()
    return strings.TrimPrefix(scanner.Text(), key)
  }

  return Post{
    Title:       getMetaValue(titleKey),
    Description: getMetaValue(descriptionKey),
    Tags:        strings.Split(getMetaValue(tagsKey), ", "),
    Body:        readBody(scanner),
  }, nil
}

func readBody(scanner *bufio.Scanner) string {
  scanner.Scan()
  buf := bytes.Buffer{}
  for scanner.Scan() {
    fmt.Fprintln(&buf, scanner.Text())
  }
  return strings.TrimPrefix(buf.String(), "\n")
}
