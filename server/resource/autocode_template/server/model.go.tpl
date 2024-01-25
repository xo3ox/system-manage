// 自动生成模板{{.StructName}}
package {{.Package}}

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	{{ if .HasTimer }}"time"{{ end }}
	{{ if .NeedJSON }}"gorm.io/datatypes"{{ end }}
)

// {{.StructName}} 结构体
type {{.StructName}} struct {
      global.GVA_MODEL_V2 {{- range .Fields}}
            {{- if eq .FieldType "enum" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};type:enum({{.DataTypeLong}});"` // {{.Comment}}
            {{- else if eq .FieldType "picture" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};"` // {{.Comment}}
            {{- else if eq .FieldType "file" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};"` // {{.Comment}}
            {{- else if eq .FieldType "pictures" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};"` // {{.Comment}}
            {{- else if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};"` // {{.Comment}}
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};"` // {{.Comment}}
            {{- end }} {{- end }}
}

{{ if .TableName }}
// TableName {{.StructName}} 表名
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}
