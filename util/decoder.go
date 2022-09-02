package util

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type MapStructureUnmarshaler interface {
	UnmarshalMapStructure(f reflect.Type, data interface{}) error
}

func UnmarshallerHookFunc(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	result := reflect.New(t).Interface()
	unmarshaller, ok := result.(MapStructureUnmarshaler)
	if !ok {
		return data, nil
	}
	if err := unmarshaller.UnmarshalMapStructure(f, data); err != nil {
		return nil, err
	}
	return result, nil
}

// DecodeOther uses mapstructure to decode into target structure. Unused keys cause errors.
func DecodeOther(other, cc interface{}) error {
	decoderConfig := &mapstructure.DecoderConfig{
		Result:           cc,
		ErrorUnused:      true,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.TextUnmarshallerHookFunc(),
			UnmarshallerHookFunc,
		),
	}

	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err == nil {
		err = decoder.Decode(other)
	}

	return err
}
