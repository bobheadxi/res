package res

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/bobheadxi/res/base"
)

// KV is used for defining specific values to be unmarshalled from BaseResponse
// data
type KV struct {
	Key   string
	Value interface{}
}

// Unmarshal reads the response and unmarshalls the BaseResponse as well any
// requested key-value pairs.
// For example:
//
// 	  var prop = map[string]string
//    api.Unmarshal(resp.Body, api.KV{Key: "prop", Value: &prop})
//
// Values provided in KV.Value MUST be explicit pointers, even if the value is
// a pointer type, ie maps and slices.
func Unmarshal(r io.Reader, kvs ...KV) (*base.Response, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("could not read bytes from reader: %s", err.Error())
	}

	// Unmarshal data into a BaseResponse, replacing BaseResponse.Data with a
	// map to preserve raw JSON data in the keys
	var (
		data = make(map[string]json.RawMessage)
		resp = base.Response{Data: &data}
	)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, fmt.Errorf("could not unmarshal data from reader: %s", err.Error())
	}

	// Unmarshal all requested kv-pairs, silently ignoring errors
	for _, kv := range kvs {
		json.Unmarshal(data[kv.Key], kv.Value)
	}

	return &resp, nil
}
