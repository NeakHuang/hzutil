// Date: 2023/7/8
// Author:
// Description：

package hzconv

import "github.com/mitchellh/mapstructure"

var MapStruct = map2Struct{}

type map2Struct struct {
}

func (map2Struct) Decode(input interface{}, output interface{}, tagName ...string) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   output,
		TagName:  "json",
	}
	if len(tagName) > 0 {
		config.TagName = tagName[0]
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}
