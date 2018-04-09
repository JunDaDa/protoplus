package gogopb

import (
	"fmt"
	"github.com/davyxu/protoplus/codegen"
)

// 报错行号+7
const protoCodeTemplate = `// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
syntax = "proto3";

package {{.PackageName}};

{{range $a, $enumobj := .Enums}}
enum {{.Name}} {	{{range .Fields}}
	{{.Name}} = {{PbTagNumber $enumobj .}}; {{end}}
}{{end}}

{{range $a, $obj := .Structs}}
{{ObjectLeadingComment .}}
message {{.Name}} {	{{range .Fields}}
	{{PbTypeName .}} {{GoFieldName .}} = {{PbTagNumber $obj .}};{{FieldTrailingComment .}} {{end}}
}
{{end}}
`

func Run(ctx *Context) error {

	gen := codegen.NewCodeGen("proto").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(UsefulFunc).
		ParseTemplate(protoCodeTemplate, ctx)

	if gen.Error() != nil {
		fmt.Println(string(gen.Data()))
		return gen.Error()
	}

	return gen.WriteOutputFile(ctx.OutputFileName).Error()
}