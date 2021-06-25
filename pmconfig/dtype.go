

/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmconfig

type DType  string
type DValue interface{}

const (
    DTypeString     DType = "string"
    DTypeInteger    DType = "integer"
    DTypeBool       DType = "bool"
    DTypeNumeric    DType = "numeric"
)
//EOF
