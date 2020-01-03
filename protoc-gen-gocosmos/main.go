package main

import (
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
	_ "github.com/regen-network/cosmos-proto/interfacetype"
)

func main() {
	req := command.Read()
	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)

	vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
	vanity.ForEachFile(files, vanity.TurnOnSizerAll)
	vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)
	//vanity.ForEachFile(files, vanity.TurnOffGoStringerAll)
	//vanity.ForEachFile(files, vanity.TurnOffGoGettersAll)

	resp := command.Generate(req)
	command.Write(resp)
}
