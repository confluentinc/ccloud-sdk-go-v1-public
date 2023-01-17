package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
)

const (
	version = "1.1"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println(version)
		os.Exit(0)
	}

	req := command.Read()
	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)

	for _, opt := range []func(*descriptor.FileDescriptorProto){
		vanity.TurnOnMarshalerAll,
		vanity.TurnOnUnmarshalerAll,
		vanity.TurnOnSizerAll,
	} {
		vanity.ForEachFile(files, opt)
	}

	vanity.ForEachFieldInFiles(files, func(field *descriptor.FieldDescriptorProto) {
		setCustomTagFieldOption("db", field)
		setCustomTagFieldOption("url", field)
	})

	resp := command.Generate(req)
	command.Write(resp)
}

func setCustomTagFieldOption(fieldKey string, field *descriptor.FieldDescriptorProto) {
	moreTagsExtension := gogoproto.GetMoreTags(field)

	var moreTags reflect.StructTag
	if moreTagsExtension != nil {
		moreTags = reflect.StructTag(*moreTagsExtension)
	}

	val := string(moreTags)
	if _, ok := moreTags.Lookup(fieldKey); !ok {
		fieldName := *field.Name
		fieldTag := fieldName + ",omitempty"
		repeatedNativeType := (!field.IsMessage() && !gogoproto.IsCustomType(field) && field.IsRepeated())
		if !gogoproto.IsNullable(field) && !repeatedNativeType {
			fieldTag = fieldName
		}
		if val != "" {
			val += " "
		}
		val += fmt.Sprintf(`%s:"%s"`, fieldKey, fieldTag)
	}

	if field.Options == nil {
		field.Options = &descriptor.FieldOptions{}
	}

	if err := proto.SetExtension(field.Options, gogoproto.E_Moretags, &val); err != nil {
		panic(err)
	}
}
