package main

import (
	"encoding/json"
	"time"
	"fmt"
	"github.com/ajaybodhe/encoding/models"
	"github.com/ajaybodhe/encoding/models/colferModels"
	"github.com/ajaybodhe/encoding/models/gogoModels"
	"github.com/ajaybodhe/encoding/models/gencodeModels"
	"github.com/davecgh/go-xdr/xdr2"
	"bytes"
	"github.com/vmihailenco/msgpack"
	"encoding/gob"
	//"github.com/ikkerens/ikeapack"
	goavro2 "github.com/linkedin/goavro"
	
)

var (
	codec *goavro2.Codec
)
// xdr2, gencode
func main() {
	e := &models.A{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		Loc:&models.B{Location:"pune"},
	}
	/*
	JSON
	 */
	d, err := json.Marshal(e)
	fmt.Println("json eData:", string(d), err, len(d))
	e1:=new(models.A)
	err= json.Unmarshal(d, e1)
	fmt.Println("json Data:", e1, err)
	
	/*
	msg pack
	 */
	d1, err :=msgpack.Marshal(e)
	fmt.Println("msg eData:", string(d1), err)
	e2:=new(models.A)
	err = msgpack.Unmarshal(d1, e2)
	fmt.Println("msg Data:", e2.Loc, err)
	
	/*
	xdr
	 */
	var w bytes.Buffer
	i, err := xdr.Marshal(&w, e)
	fmt.Println("xdr eData:", w.String(), i, err)
	e3:=new(models.A)
	i, err = xdr.Unmarshal(bytes.NewReader(w.Bytes()), e3)
	fmt.Println("xdr Data:", e3, i, err)
	
	/*
	gob
	 */
	var wgob bytes.Buffer
	dec := gob.NewDecoder(&wgob)
	enc := gob.NewEncoder(&wgob)
	err = enc.Encode(e)
	fmt.Println("gob eData:", wgob.String(), err)
	e4:=new(models.A)
	err = dec.Decode(e4)
	fmt.Println("gob Data:", e4, err)
	
	/*
	IKEA
	 */
	//b := new(bytes.Buffer)
	//err = ikea.Pack(b, e)
	//fmt.Println("ikea eData:", b.String(), err)
	//e5:=new(models.A)
	//err = ikea.Unpack(b, e5)
	//fmt.Println("ikea Data:", e5, err)
	
	/*
	avro
	 */
	codec, err = goavro2.NewCodec(string(models.SchemaAvro))
	if err != nil {
		panic(err)
	}
	binary, err := codec.BinaryFromNative(nil,e.ToStringMap())
	fmt.Println("avro eData:", string(binary), err)
	e6Map, _, err := codec.NativeFromBinary(binary)
	e6 := new(models.A)
	e6.Loc = new(models.B)
	e6.Loc.Location = e6Map.(map[string]interface{})["B"].(map[string]interface{})["my.namespace.com.b"].(map[string]interface{})["Location"].(string)
	e6.Phone = e6Map.(map[string]interface{})["phone"].(string)
	fmt.Println("avro Data:", e6, err)
	
	/*
	colfer
	 */
	ca := &colferModels.ColferA{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		Loc:&colferModels.ColferB{Location:"pune"},
	}
	dColf, err := ca.MarshalBinary()
	fmt.Println("colf eData:", string(dColf), err, len(dColf))
	ca1 := new(colferModels.ColferA)
	err = ca1.UnmarshalBinary(dColf)
	fmt.Println("colf Data:", ca1.Loc.Location, err)
	
	/*
	gogoprotobuf
	 */
	gogo := &gogoModels.GogoProtoBufA{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		GogoProtoBufB:&gogoModels.GogoProtoBufB{Location:"pune"},
	}
	dGogo, err := gogo.Marshal()
	fmt.Println("gogo eData:", string(dGogo), err, len(dGogo))
	gogo1 := new(gogoModels.GogoProtoBufA)
	err = gogo1.Unmarshal(dGogo)
	fmt.Println("gogo Data:", gogo1.GogoProtoBufB.Location, err)
	
	/*
	gencode
	 */
	genc := &gencodeModels.GencodeUnsafeA{
		Name:"Ajay",
		BirthDay:time.Now().Unix(),
		Phone:"7758863774",
		Siblings:1,
		Spouse:true,
		Money:3.44,
		Loc:&gencodeModels.GencodeUnsafeB{Location:"pune"},
	}
	dGenc, err := genc.Marshal(nil)
	fmt.Println("gencode eData:", string(dGenc), err, len(dGenc))
	genc1 := new(gencodeModels.GencodeUnsafeA)
	_, err = genc1.Unmarshal(dGenc)
	fmt.Println("gencode Data:", genc1.Loc.Location, err)
}
