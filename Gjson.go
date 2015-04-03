package GJSON

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type GjsonDecoder struct{}

func (g GjsonDecoder) Decode(encoded []byte) (map[string]interface{}, error) {
	var j map[string]interface{}
	err := json.Unmarshal([]byte(encoded), &j)
	if err != nil {
		return j, err
	}
	return j, nil
}

func (g GjsonDecoder) printValues(data map[string]interface{}) {
	for i, v := range data {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			g.printValues(v.(map[string]interface{}))
		} else {
			fmt.Printf(" %-20s |  %-20v| %-100v  \n", i, reflect.TypeOf(v).Kind(), v)
		}
	}
}
func (g GjsonDecoder) PrintJSON(encoded []byte) error {
	m, err := g.Decode(encoded)
	if err != nil {
		return err
	}
	g.printValues(m)
	return nil
}
