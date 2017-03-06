package traverse

import (
  "testing"

  "fmt"
)


func TestFilesystem(t *testing.T) {
  n := GetNodeFilesystem("/")

  fmt.Printf("%v\n", n.GetName())
  children := n.Children()
  fmt.Printf("%v\n", children)
}
