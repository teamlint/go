package redis

import "encoding/json"

var DefaultSerializer = defaultSerializer{}

type defaultSerializer struct{}

func (d defaultSerializer) Marshal(value interface{}) ([]byte, error) {
	if jsonM, ok := value.(json.Marshaler); ok {
		return jsonM.MarshalJSON()
	}

	return json.Marshal(value)
}

func (d defaultSerializer) Unmarshal(b []byte, outPtr interface{}) error {
	if jsonUM, ok := outPtr.(json.Unmarshaler); ok {
		return jsonUM.UnmarshalJSON(b)
	}

	return json.Unmarshal(b, outPtr)
}
