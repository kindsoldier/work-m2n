/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmdescr

import (
    "encoding/json"
)

type IMeasureDescr interface {
}

type MeasureDescr struct {
    OwnerId     UUID            `json:"ownerId"`
    MeasureId   UUID            `json:"measureId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json""value"`
}

func NewMeasureDescr(ownerId UUID, measureId UUID, name string, dType DType, value DValue) *MeasureDescr {
    var config MeasureDescr
    config.OwnerId     = ownerId
    config.MeasureId   = measureId
    config.Name        = name
    config.Type        = dType
    config.Value       = value
    return &config
}

func (this *MeasureDescr) UnmarshalJSON(data JSON) error {
    var err error
    // NOP
    return err
}

func (this *MeasureDescr) MarshalJSON() (JSON, error) {
    return json.Marshal(*this)
}

//EOF
