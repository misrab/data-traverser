package traverse

import (

)


type Node interface {
  Children() []Node

  SetName(string)
  GetName() string

  String() string
}
