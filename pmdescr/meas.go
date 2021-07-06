/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmdescr

import (
    "encoding/json"
)

type IMeasure interface {
}

type MMeasure struct {
    OwnerId     UUID            `json:"ownerId"`
    MeasureId   UUID            `json:"measureId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json""value"`
}

func NewMMeasure(ownerId UUID, measureId UUID, name string, dType DType, value DValue) *MMeasure {
    var config MMeasure
    config.OwnerId     = ownerId
    config.MeasureId   = measureId
    config.Name        = name
    config.Type        = dType
    config.Value       = value
    return &config
}

func (this *MMeasure) UnmarshalJSON(data JSON) error {
    var err error
    // NOP
    return err
}

func (this *MMeasure) MarshalJSON() (JSON, error) {
    return json.Marshal(*this)
}

//EOF
