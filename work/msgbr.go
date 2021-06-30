/*
 * Copyright: Oleg Borodin <onborodin@gmail.com>
 */


package main

import (
    //"fmt"
)

type UUID  string
type SubFunct func(subject []byte, payload []byte)

type IBroker interface {
    Subscribe(subject []byte, handler SubFunct) (UUID, error)
    Unsubsribe(subId UUID) error
    Publish(subject []byte, payload []byte) error
}

type Subscriber struct {
    Id          UUID
    Subject     []byte
    Handler     SubFunct
}
func NewSubscriber(subject []byte, handler SubFunct) *Subscriber {
    var subscr Subscriber
    subscr.Subject  = subject
    subscr.Handler  = handler
    subscr.Id       = "xxxx"
    return &subscr
}

func (this *Subscriber) GetId() UUID {
    return this.Id
}

type MBroker struct {
    subscribers map[UUID]*Subscriber
}

func NewMBroker() *MBroker {
    var broker MBroker
    broker.subscribers = make(map[UUID]*Subscriber)
    return &broker
}

func (this *MBroker) Subscribe(subject []byte, handler SubFunct) (UUID, error) {
    var subscrId UUID
    var err error
    subscr := NewSubscriber(subject, handler)
    subscrId = subscr.GetId()
    this.subscribers[subscrId] = subscr
    return subscrId, err 
}

func (this *MBroker) Unsubsribe(subId UUID) error {
    var err error
    return err
}

func (this *MBroker) Publish(subject []byte, payload []byte) error {
    var err error
    return err
}



func main() {
    //var broker IBroker
    //broker = NewMBroker()
    //fmt.Println(broker.Hello())
}





