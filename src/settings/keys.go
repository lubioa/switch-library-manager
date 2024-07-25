package settings

import (
	"errors"
	"github.com/magiconair/properties"
)

var (
	keysInstance *switchKeys
)

type switchKeys struct {
	keys map[string]string
}

func (k *switchKeys) GetKey(keyName string) string {
	return k.keys[keyName]
}

func SwitchKeys() (*switchKeys, error) {
	return keysInstance, nil
}

func InitSwitchKeys(baseFolder string) (*switchKeys, error) {
	settings := ReadSettings(baseFolder)
	p, err := properties.LoadFile(settings.Prodkeys, properties.UTF8)
	if err != nil {
		return nil, errors.New("Error trying to read prod.keys [reason:" + err.Error() + "]")
	}
	keysInstance = &switchKeys{keys: map[string]string{}}
	for _, key := range p.Keys() {
		value, _ := p.Get(key)
		keysInstance.keys[key] = value
	}

	return keysInstance, nil
}
