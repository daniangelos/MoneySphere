package constants

import "fmt"

var ErrNotEnoughBalance error = fmt.Errorf("not enough balance")
var ErrClientNotFound error = fmt.Errorf("client id not found")
