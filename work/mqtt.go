/*
 * Copyright: Oleg Borodin <onborodin@gmail.com>
 */


package pmtrans

import (
    "errors"
    "time"
    "net/url"
    "fmt"

    "github.com/eclipse/paho.mqtt.golang"

    "pmapp/pmlog"
)


type RecvHandler func(subject string, payload []byte) 

type ITransport interface {
    Connect(URL string) error
    Send(subject string, payload []byte) error
    IsConnected() bool
    Recv(subject string, handler RecvHandler) error
    Disconnect() error
}


const (
    keepaliveTimeout    time.Duration   = 3 // sec
    waitTimeout         time.Duration   = 3 // sec
    pingTimeout         time.Duration   = 3 // sec
    reconnectTimeout    time.Duration   = 3 // sec
    disconnectTime      uint             = 1 // ms

    QosL1               byte = 1
    QosL2               byte = 2
    QosL3               byte = 4

    mTransportId        string  = "69e61d65-59ca-4b2c-8849-61ff9ecb39ea" 
)

type UUID = string
type MTransport struct {
    classId     UUID
    objectId    UUID
    password    string
    username    string
    url         string
    mc          mqtt.Client
}

func NewMTransport(objectId UUID) *MTransport {
    var trans MTransport
    trans.classId   = mTransportId
    trans.objectId  = objectId
    return &trans
}

func (this *MTransport) Connect(URL string) error {
    var err error
    err = this.parseURL(URL)
    if err != nil {
        return err
    }
    err = this.connect()
    if err != nil {
        return err
    }
    return err
}

func (this *MTransport) parseURL(URL string) error {
    var err error
    mURL, err := url.Parse(URL)
    if err != nil {
        return err
    }
    this.username    = mURL.User.Username()
    this.password, _ = mURL.User.Password()
    this.url = fmt.Sprintf("%s://%s", mURL.Scheme, mURL.Host)
    return err
}

func (this *MTransport) connect() error {
    var err error

    if this.mc != nil && this.mc.IsConnected() {
        this.mc.Disconnect(disconnectTime)
    }

    pmlog.LogDebug(this.username, this.password)

    opts := mqtt.NewClientOptions()
    opts.AddBroker(this.url)
    opts.SetUsername(this.username)
    opts.SetPassword(this.password)
    opts.SetClientID(this.classId)
    opts.SetKeepAlive(keepaliveTimeout)
    opts.SetPingTimeout(pingTimeout)
    opts.SetAutoReconnect(true)

    this.mc = mqtt.NewClient(opts)

    token := this.mc.Connect()
    for !token.WaitTimeout(waitTimeout * time.Second) {
    }

    err = token.Error()
    if err != nil {
        return err
    }
    return err
}

func (this *MTransport) Reconnect() error {
    return this.connect()
}

func (this *MTransport) Send(subject string, payload []byte) error {
    if this.mc == nil || !this.mc.IsConnected() {
        return errors.New("mqtt transport not connected")
    }
    token := this.mc.Publish(subject, QosL1, true, payload)
    return token.Error()
}

func (this *MTransport) Recv(subject string, handler RecvHandler) error {
    if this.mc == nil || !this.mc.IsConnected() {
        return errors.New("mqtt transport not connected")
    }
    messageHandler := func(client mqtt.Client, message mqtt.Message) {
        subject := message.Topic()
        payload := message.Payload()
        handler(subject, payload)  
    }
    token := this.mc.Subscribe(subject, QosL1, messageHandler)
    return token.Error()
}

func (this *MTransport) IsConnected() bool {
    if  this.mc != nil {
        return this.mc.IsConnected()
    }
    return false
}

func (this *MTransport) Disconnect() error {
    var err error
    if this.mc != nil && this.mc.IsConnected() {
        this.mc.Disconnect(disconnectTime)
    }
    return err
}
//EOF


