package testing

import (
	"testing"
	"time"
	"github.com/ajaybodhe/encoding/models"
	"encoding/json"
	"bytes"
	"github.com/vmihailenco/msgpack"
	"github.com/davecgh/go-xdr/xdr2"
	"encoding/gob"
	goavro2 "github.com/linkedin/goavro"
	"github.com/ajaybodhe/encoding/models/colferModels"
)


func BenchmarkJson(b *testing.B) {
	e := &models.A{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		Loc:&models.B{Location:"pune"},
	}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		d, err := json.Marshal(e)
		if err!= nil {
			panic(err)
		}
		e1:=new(models.A)
		err= json.Unmarshal(d, e1)
		if err!= nil {
			panic(err)
		}
	}
}
func BenchmarkMsgPack(b *testing.B) {
	e := &models.A{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		Loc:&models.B{Location:"pune"},
	}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		d1, err :=msgpack.Marshal(e)
		if err!= nil {
			panic(err)
		}
		e2:=new(models.A)
		msgpack.Unmarshal(d1, e2)
		if err!= nil {
			panic(err)
		}
	}
}

func BenchmarkXDR(b *testing.B) {
	e := &models.A{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		Loc:&models.B{Location:"pune"},
	}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		var w bytes.Buffer
		i, err := xdr.Marshal(&w, e)
		if err!= nil || i <= 0 {
			panic(err)
		}
		e3:=new(models.A)
		i, err = xdr.Unmarshal(bytes.NewReader(w.Bytes()), e3)
		if err!= nil || i <= 0 {
			panic(err)
		}
	}
}

func BenchmarkAvro(b *testing.B) {
	e := &models.A{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		Loc:&models.B{Location:"pune"},
	}
	codec, err := goavro2.NewCodec(string(models.SchemaAvro))
	if err != nil {
		panic(err)
	}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		binary, err := codec.BinaryFromNative(nil,e.ToStringMap())
		if err != nil {
			panic(err)
		}
		e6Map, _, err := codec.NativeFromBinary(binary)
		if err != nil {
			panic(err)
		}
		e6 := new(models.A)
		e6.Loc = new(models.B)
		e6.Loc.Location = e6Map.(map[string]interface{})["B"].(map[string]interface{})["my.namespace.com.b"].(map[string]interface{})["Location"].(string)
		e6.Phone = e6Map.(map[string]interface{})["phone"].(string)
	}
}

func BenchmarkColfer(b *testing.B) {
	ca := &colferModels.ColferA{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		Loc:&colferModels.ColferB{Location:"pune"},
	}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		dColf, err := ca.MarshalBinary()
		if err!= nil  {
			panic(err)
		}
		ca1 := new(colferModels.ColferA)
		err = ca1.UnmarshalBinary(dColf)
		if err!= nil  {
			panic(err)
		}
	}
}

func BenchmarkGob(b *testing.B) {
	e := &models.A{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		Loc:&models.B{Location:"pune"},
	}
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		var w bytes.Buffer
		dec := gob.NewDecoder(&w)
		enc := gob.NewEncoder(&w)
		err := enc.Encode(e)
		if err!= nil  {
			panic(err)
		}
		e3:=new(models.A)
		err = dec.Decode(e3)
		if err!= nil {
			panic(err)
		}
	}
}
