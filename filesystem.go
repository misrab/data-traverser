package traverse

import (
  "io/ioutil"
)


/*
  Implements Adapter interface for the file system
*/


type Filesystem struct {

}


// in this case the input is a filepath
func GetNodeFilesystem(s string) Node {
  result := new(fsNode)
  result.SetName(s)

  return result
}


type fsNode struct {
  Name string
}

func (n *fsNode) String() string {
  return n.Name
}
func (n *fsNode) SetName(s string) {
  n.Name = s
}
func (n *fsNode) GetName() string {
  return n.Name
}

func (n *fsNode) Children() []Node {
  dir := n.Name

  // list the contents and return list of nodes
  files, _ := ioutil.ReadDir(dir)
  N := len(files)
  result := make([]Node, N)
  for i, f := range files {
    result[i] = GetNodeFilesystem(f.Name())
  }

  return result
}
