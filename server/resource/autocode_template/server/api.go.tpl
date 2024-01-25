package {{.Package}}

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    {{.Package}}Req "github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    {{- if .NeedValid }}
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
    {{- else if .AutoCreateResource}}
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
    {{- end }}
)

type {{.StructName}}Api struct {
}

var {{.Abbreviation}}Service = service.ServiceGroupApp.{{.PackageT}}ServiceGroup.{{.StructName}}Service


// Create{{.StructName}} 创建
// @Tags {{.StructName}}
// @Summary 创建
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body {{.Package}}.{{.StructName}} true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Create{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	{{- if .AutoCreateResource }}
    {{.Abbreviation}}.CreatedBy = utils.GetUserID(c)
	{{- end }}
    {{- if .NeedValid }}
    verify := utils.Rules{
    {{- range $index,$element := .Fields }}
       {{- if $element.Require }}
        "{{$element.FieldName}}":{utils.NotEmpty()},
        {{- end }}
    {{- end }}
    }
	if err := utils.Verify({{.Abbreviation}}, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
    {{- end }}
	if err := {{.Abbreviation}}Service.Create{{.StructName}}(&{{.Abbreviation}}); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(),"创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete{{.StructName}} 删除
// @Tags {{.StructName}}
// @Summary 删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "传递id"
// @Success 200 {object} response.Response "返回结果"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [delete]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := {{.Abbreviation}}Service.Delete{{.StructName}}({{.Abbreviation}}); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(),"删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Delete{{.StructName}}ByIds 批量删除
// @Tags {{.StructName}}
// @Summary 批量删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body request.IdsReq true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /{{.Abbreviation}}/delete{{.StructName}}ByIds [delete]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete{{.StructName}}ByIds(c *gin.Context) {
	var ids request.IdsReq
    err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := {{.Abbreviation}}Service.Delete{{.StructName}}ByIds(ids{{- if .AutoCreateResource }}{{- end }}); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(),"批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update{{.StructName}} 更新
// @Tags {{.StructName}}
// @Summary 更新
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - body {{.Package}}.{{.StructName}} true "传递object"
// @Success 200 {object} response.Response "返回结果"
// @Router /{{.Abbreviation}}/update{{.StructName}} [put]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Update{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindJSON(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	    {{- if .AutoCreateResource }}
    {{.Abbreviation}}.UpdatedBy = utils.GetUserID(c)
        {{- end }}
	{{- if .NeedValid }}
      verify := utils.Rules{
      {{- range $index,$element := .Fields }}
         {{- if $element.Require }}
          "{{$element.FieldName}}":{utils.NotEmpty()},
          {{- end }}
      {{- end }}
      }
    if err := utils.Verify({{.Abbreviation}}, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
    {{- end }}
	if err := {{.Abbreviation}}Service.Update{{.StructName}}({{.Abbreviation}}); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(),"更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Find{{.StructName}} 用ID查询
// @Tags {{.StructName}}
// @Summary 用id查询
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param id query int true "id"
// @Success 200 {object} response.Response{data={{.Package}}.{{.StructName}}} "返回结果"
// @Router /{{.Abbreviation}}/find{{.StructName}} [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Find{{.StructName}}(c *gin.Context) {
	var {{.Abbreviation}} {{.Package}}.{{.StructName}}
	err := c.ShouldBindQuery(&{{.Abbreviation}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if re{{.Abbreviation}}, err := {{.Abbreviation}}Service.Get{{.StructName}}({{.Abbreviation}}.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithDetailed(err.Error(),"查询失败", c)
	} else {
		response.OkWithData(gin.H{"re{{.Abbreviation}}": re{{.Abbreviation}}}, c)
	}
}

// Get{{.StructName}}List 分页获取列表
// @Tags {{.StructName}}
// @Summary 分页获取列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param - query {{.Package}}Req.{{.StructName}}Search true "传递object"
// @Success 200 {object} response.Response{data=[]{{.Package}}.{{.StructName}}} "返回结果"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Get{{.StructName}}List(c *gin.Context) {
	var info {{.Package}}Req.{{.StructName}}Search
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := {{.Abbreviation}}Service.Get{{.StructName}}InfoList(info); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithDetailed(err.Error(),"获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     info.Page,
            PageSize: info.PageSize,
        }, "获取成功", c)
    }
}
