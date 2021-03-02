package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	baseUrl := "http://localhost:8000"

	// 1. 请求
	var (
		resp *http.Response
		err error
	)
	resp, err = http.Get(baseUrl + "/")

	// 2. 检测
	assert.NoError(t, err, "有错误发生, err 不为空")
	assert.Equal(t, 200, resp.StatusCode, "应返回状态码 200")
}
