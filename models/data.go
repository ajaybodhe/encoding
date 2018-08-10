package models

import (
	"github.com/linkedin/goavro"
	//"fmt"
)

//go:generate msgp
type B struct {
	Location string `msg:"loc"`
}

func (b *B)ToStringMap() map[string]interface{} {
	m := map[string]interface{} {
		"Location":b.Location,
	}
	return m
}

//go:generate msgp
type A struct {
	Name     string `msg:"name"`
	BirthDay int64 `msg:"bd"`
	Phone    string `msg:"p"`
	Siblings int32 `json:"-" msg:"s"`
	Spouse   bool `msg:"sp"`
	Money    float64 `msg:"m"`
	Loc *B `msg:"b"`
}

func (e *A) ToStringMap() map[string]interface{} {
	m := map[string]interface{}{
		"name":     e.Name,
		"birthday": e.BirthDay,
		"phone":    e.Phone,
		"siblings": e.Siblings,
		"spouse":   e.Spouse,
		"money":    e.Money,
		
	}
	if e.Loc != nil {
		m["B"] = goavro.Union("my.namespace.com.b", e.Loc.ToStringMap())
	} else {
		m["B"] = goavro.Union("null", nil)
	}
	//fmt.Println("avro map;", m)
	return m
}

var (
	SchemaAvro = `{
		"namespace": "my.namespace.com",
		"type":	"record",
		"name": "indentity",
		"fields": [
			{ "name": "name", "type": "string"},
			{ "name": "birthday", "type": "long"},
			{ "name": "phone", "type": "string"},
			{ "name": "siblings", "type": "int"},
			{ "name": "spouse", "type": "boolean"},
		    { "name": "money", "type": "double"},
			{ "name": "B", "type": ["null",{
				"namespace": "my.namespace.com",
				"type":	"record",
				"name": "b",
				"fields": [
					{ "name": "Location", "type": "string" }
				]
			}],"default":null}
		]
}`
)
