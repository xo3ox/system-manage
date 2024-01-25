package function

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

// SpiltSlice 拆分切片
// slice 要拆分的切片 , size 一行有几个
func SpiltSlice[T any](slice []T, size int) [][]T {
	lens := len(slice)
	splitSlice := make([][]T, 0)
	mod := math.Ceil(float64(lens) / float64(size))
	for i := 0; i < int(mod); i++ {
		if i == int(mod)-1 {
			splitSlice = append(splitSlice, slice[i*size:])
		} else {
			splitSlice = append(splitSlice, slice[i*size:i*size+size])
		}

	}
	return splitSlice
}

// ExportExcel 导出Excel
// excelPath 导出路径，excelName 导出文件名，header 头部标题，data 二维切片
func ExportExcel(excelPath, excelName string, header []string, data [][]any) error {

	var (
		col string
		sh  = 'A'
		pre = ""
	)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet(excelName)
	if err != nil {
		return err
	}
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	for _, v := range header {
		col = pre + string(sh) + "1"
		err = f.SetCellValue(excelName, col, v)
		if err != nil {
			return err
		}
		sh++
		if sh > 'Z' {
			sh = 'A'
			pre = string(sh)
		}
	}

	for k, line := range data {
		sh = 'A'
		pre = ""
		numS := strconv.Itoa(k + 2)
		for _, w := range line {
			col = pre + string(sh) + numS
			err = f.SetCellValue(excelName, col, w)
			if err != nil {
				return err
			}
			sh++
			if sh > 'Z' {
				sh = 'A'
				pre = string(sh)
			}
		}
	}

	if err := f.SaveAs(excelPath + excelName + ".xlsx"); err != nil {
		return err
	}
	return nil
}

// Random 随机获取切片中指定个的数据
// data 切片， queryCount 随机个数
func Random[T any](data []T, queryCount int) (result []T) {
	copyData := make([]T, len(data))
	copy(copyData, data)
	// rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < queryCount; i++ {
		index := rand.Intn(len(data))
		result = append(result, data[index])
		data = append(data[:index], data[index+1:]...)
		if len(data) == 0 {
			tempData := make([]T, len(copyData))
			copy(tempData, copyData)
			data = tempData
		}
	}
	return
}

// ParamsInterface 调用第三方接口统一封装
// restFul 请求方式 POST、GET		url 请求地址		reqBody 请求体		headers 请求头
// params GET请求query入参		respData 接收响应的模型
func ParamsInterface(restFul, uri string, reqBody io.Reader, headers, params map[string]string, respData interface{}) error {

	client := &http.Client{}

	req, err := http.NewRequest(restFul, uri, reqBody)
	if err != nil {
		return fmt.Errorf("创建http.NewRequest失败:%v", err)
	}

	if len(headers) > 0 {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	if len(params) > 0 {
		param := make(url.Values)

		for k, v := range params {
			param.Set(k, v)
		}

		req.URL.RawQuery = param.Encode()
	}

	global.GVA_LOG.Debug("请求url是:", zap.Any("url", req.URL.String()))

	// 发送http请求，获得相应
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("发送client.Do失败,err:", zap.Error(err))
		return fmt.Errorf("发送client.Do失败,err:%v", err)
	}

	if resp.StatusCode != http.StatusOK {
		global.GVA_LOG.Error("HTTP响应失败", zap.Any("resp.StatusCode", resp.StatusCode))
		return fmt.Errorf("HTTP响应失败，状态码：%v", resp.StatusCode)
	}

	// 读响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("读取响应体内容失败", zap.Error(err), zap.String("body", string(body)))
		return fmt.Errorf("读取响应体内容失败,err:%v", err)
	}
	defer resp.Body.Close()

	global.GVA_LOG.Info("响应体中的内容是：", zap.String("body", string(body)))

	// 解析body中数据到结构体
	if err = json.Unmarshal(body, &respData); err != nil {
		global.GVA_LOG.Error("反序列化响应体内容失败", zap.Error(err), zap.String("body", string(body)))
		return fmt.Errorf("反序列化响应体内容失败,err:%v", err)
	}

	// // 解析body中数据到结构体
	// if err = json.NewDecoder(resp.Body).Decode(&respData); err != nil {
	// 	global.GVA_LOG.Error("反序列化响应体内容失败", zap.Error(err))
	// 	return fmt.Errorf("反序列化响应体内容失败,err:%v", err)
	// }

	global.GVA_LOG.Debug("解析出的json数据是:", zap.Any("respData", respData))

	return nil
}
