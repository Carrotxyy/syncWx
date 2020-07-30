package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
	以post方式，向指定路径，发送指定类型的数据

	@path 请求路径
	@data 数据
	@out 装载响应数据
	@contentType 请求体类型

	返回参数 (错误)
*/

func HttpPost(path string, data interface{}, out interface{}, contentType string) error {
	// 将数据json化
	jsonStr, _ := json.Marshal(data)
	// 发送请求
	resp, err := http.Post(path, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("发送请求错误:", err)
		return err
	}
	// 关闭响应流
	defer resp.Body.Close()
	// 读取响应
	result, _ := ioutil.ReadAll(resp.Body)
	// 解析响应数据
	err = json.Unmarshal(result, out)
	if err != nil {
		fmt.Println("解析响应数据错误!", err)
		return err
	}
	return nil
}

/**
	请求数据

	@path 请求路径
	@out 装载响应数据

	返回参数 (错误)
*/
func HttpGet(path string, out interface{}) error {
	// 发起请求
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println("请求url错误:", err)
		return err
	}
	// 获取响应流
	result, _ := ioutil.ReadAll(resp.Body)
	// 解析数据
	err = json.Unmarshal(result, out)
	if err != nil {
		fmt.Println("解析响应数据错误！", err)
		return err
	}
	return nil
}

/**
	获取key

	@path 获取key的路由

	返回参数 (key)
 */
func GetKey(path string) interface{} {
	// 发起请求
	res, err := http.Get(path)
	if err != nil {
		fmt.Println("获取key:", err)
		return ""
	}
	// 关闭响应流
	defer res.Body.Close()
	// 获取响应数据
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
