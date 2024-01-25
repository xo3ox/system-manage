package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    {{- range .Fields}}
        {{- if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
    Start{{.FieldName}}  *{{.FieldType}}  `json:"start{{.FieldName}}" form:"start{{.FieldName}}"`
    End{{.FieldName}}  *{{.FieldType}}  `json:"end{{.FieldName}}" form:"end{{.FieldName}}"`
        {{- end }}
       {{- end }}
    request.PageInfo
    {{- if .NeedSort}}
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
    {{- end}}
}
