/*
 * Copyright: Oleg Borodin <onborodin@gmail.com>
 */

package pmstores

import (
    "errors"
    "net/url"
    "path/filepath"
    "encoding/json"

    "github.com/jmoiron/sqlx"
    _  "github.com/mattn/go-sqlite3"
    
    "pmapp/pmdrivers"
    //"pmapp/pmlog"
)

type IStore interface {
    StoreDevice(device pmdrivers.IDevice) error
    LoadDevice(objectId UUID) (pmdrivers.IDevice, error)

    LoadDevices() ([]pmdrivers.IDevice, error)
    StoreDevices([]pmdrivers.IDevice) (error)
} 

const (
    DStoreClassId UUID = "64c92cb5-bba0-4477-97af-c85e20af95d4"
)

type UUID = string
type JSON = []byte

type DStore struct {
    path        string
    schema      string
    classId     UUID
    dbx         *sqlx.DB
}

func NewDStore() *DStore {
    var store DStore
    store.classId = DStoreClassId
    return &store
} 

func (this *DStore) Open(sURL string) error {
    var err error
    mURL, err := url.Parse(sURL)
    if err != nil {
        return err
    }
    this.path   = filepath.Join(mURL.Host, mURL.Path)
    this.schema = mURL.Scheme
    this.dbx, err = sqlx.Open(this.schema, this.path)
    if err != nil {
        return err
    }
    err = this.dbx.Ping()
    if err != nil {
        return err
    }
    return err
}

func (this DStore) Close() error {
    var err error
    if this.dbx != nil {
        return this.dbx.Close()
    }
    return err
}

const (
    schema string = `
        -- DROP TABLE IF EXISTS devices;
        CREATE TABLE IF NOT EXISTS devices (
            object_id   TEXT PRIMARY KEY,
            object_name TEXT NOT NULL UNIQUE,
            class_id    TEXT NOT NULL
        );

        -- DROP TABLE IF EXISTS configs;
        CREATE TABLE IF NOT EXISTS configs (
            device_id   TEXT PRIMARY KEY,
            class_id    TEXT NOT NULL,
            type        TEXT NOT NULL,
            value       BLOB NOT NULL
        );
    `
)


func (this *DStore) Migrate() error {
    var err error
    _, err = this.dbx.Exec(schema)
    if err != nil {
        return err
    }
    return err
}

type DeviceRecord struct {
    ObjectId        UUID                `db:"object_id"`
    ObjectName      string              `db:"object_name"`
    ClassId         UUID                `db:"class_id"`
}

type ConfigRecord struct {
    DeviceId        UUID                `db:"device_id"`
    ClassId         UUID                `db:"class_id"`
    Type            string              `db:"type"`
    Value           JSON                `db:"value"`
}

func (this *DStore) StoreDevice(device pmdrivers.IDevice) error {
    var err error
    err = this.storeDevice(device)
    if err != nil {
        return err
    }
    
    return err
}

func (this *DStore) StoreDevices(devices []pmdrivers.IDevice) error {
    var err error
    for i := range devices {
        err = this.storeDevice(devices[i])
        if err != nil {
            return err
        }
    }
    return err
}


func (this *DStore) storeDevice(device pmdrivers.IDevice) error {
    var err error
    request := `INSERT OR REPLACE INTO devices(object_id, object_name, class_id) VALUES ($1, $2, $3)`
    _, err = this.dbx.Exec(request, device.GetObjectId(), device.GetObjectName(), device.GetClassId())
    if err != nil {
        return err
    }
    configs := device.GetConfigs()
    for _, config := range configs {
        err = this.storeConfig(device.GetObjectId(), config)
        if err != nil {
            return err
        }
    }
    return err
}

func (this *DStore) storeConfig(objectId UUID, config pmdrivers.IConfig) error {
    var err error
    request := `INSERT OR REPLACE INTO configs(device_id, class_id, type, value) VALUES ($1, $2, $3, $4)`

    value, err := json.Marshal(config.GetValue())
    if err != nil {
        return err
    }
    _, err = this.dbx.Exec(request, objectId, config.GetClassId(), config.GetType(), value)
    if err != nil {
        return err
    }
    return err
}

func (this *DStore) LoadDevices() ([]pmdrivers.IDevice, error) {
    var err     error
    devices := make([]pmdrivers.IDevice, 0)

    devices, err = this.loadDevices()
    if err != nil {
        return devices, err
    }
    for i := range devices { 
        device, err := this.loadConfigs(devices[i])
        if err != nil {
            return devices, err
        }
        devices[i] = device
    }
    return devices, err
}

func (this *DStore) loadDevices() ([]pmdrivers.IDevice, error) {
    var err     error
    devices := make([]pmdrivers.IDevice, 0)

    deviceRecords := make([]DeviceRecord, 0)

    request := `SELECT object_id, object_name, class_id FROM devices`
    
    err = this.dbx.Select(&deviceRecords, request)
    if err != nil {
        return devices, err
    }

    if len(deviceRecords) == 0 {
        return devices, errors.New("devices not found")
    }
     
    for _, deviceRecord := range deviceRecords {
        switch deviceRecord.ClassId {
            case pmdrivers.GenericClassId:
                devices = append(devices, pmdrivers.NewGenericDevice(deviceRecord.ObjectId, deviceRecord.ObjectName))
            default:
                continue
                //return device, errors.New("device class not found")
        }
    }
    return devices, err
}






















func (this *DStore) LoadDevice(objectId UUID) (pmdrivers.IDevice, error) {
    var err     error
    var device  pmdrivers.IDevice

    device, err = this.loadDevice(objectId)
    if err != nil {
        return device, err
    }
    device, err = this.loadConfigs(device)
    if err != nil {
        return device, err
    }
    return device, err
}

func (this *DStore) loadDevice(objectId UUID) (pmdrivers.IDevice, error) {
    var err     error
    var device  pmdrivers.IDevice

    deviceRecords := make([]DeviceRecord, 0)

    request := `SELECT object_id, object_name, class_id FROM devices WHERE object_id = $1 LIMIT 1`
    
    err = this.dbx.Select(&deviceRecords, request, objectId)
    if err != nil {
        return device, err
    }

    if len(deviceRecords) == 0 {
        return device, errors.New("device not found")
    }
     
    deviceRecord := deviceRecords[0]
    switch deviceRecord.ClassId {
        case pmdrivers.GenericClassId:
            device = pmdrivers.NewGenericDevice(deviceRecord.ObjectId, deviceRecord.ObjectName)
        default:
            return device, errors.New("device class not found")
    }
    return device, err
}


func (this *DStore) loadConfigs(device pmdrivers.IDevice) (pmdrivers.IDevice, error) {
    var err     error

    configRecords := make([]ConfigRecord, 0)

    request := `SELECT class_id, type, value FROM configs WHERE device_id = $1`
    
    err = this.dbx.Select(&configRecords, request, device.GetObjectId())
    if err != nil {
        return device, err
    }

    for _, configRecord := range configRecords {
        var value interface{}
        switch pmdrivers.DType(configRecord.Type) {
            case pmdrivers.DTypeString:
                var tmpValue string
                err = json.Unmarshal(configRecord.Value, &tmpValue)
                value = tmpValue
            default:
                continue
        }
        err := device.SetConfigValue(configRecord.ClassId, value)
        if err != nil {
            return device, err
        }

        //config, err := device.GetConfig(configRecord.ClassId)
        //if err != nil {
            //return device, err
        //}
        //err = config.SetValue(value)
        //if err != nil {
            //return device, err
        //}
    }
    return device, err
}
//EOF
