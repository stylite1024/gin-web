package test

import (
	"gin-web/internal/router"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试文件以_test.go为后缀名
// 测试函数前缀为Test

// 测试
// go test

// 打印详细过程
// go test -v

// 只执行某一个测试方法
// go test -v -run TestAdd

// 根据testing.Short()标志，使用-short参数跳过测试
// go test -v -short

func TestIris(t *testing.T) {
	// parallel test 并行测试
	t.Parallel()
	if testing.Short() {
		t.Skip("skip this test case")
	}

	// case list
	type args struct {
		username string
	}
	type want struct {
		w int
	}
	testCases := []struct {
		name string // case name
		args args   // case args
		want want   // case want
	}{
		{"testCase1", args{"abc"}, want{200}},
		{"testCase2", args{"def"}, want{300}},
	}

	r := router.InitRouter()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// mock一个HTTP请求
			req := httptest.NewRequest(
				"GET",                              // 请求方法
				"/api/v1/test",                            // 请求URL
				strings.NewReader(tc.args.username), // 请求参数
			)

			// mock一个响应记录器
			resp := httptest.NewRecorder()

			// 让server端处理mock请求并记录返回的响应内容
			r.ServeHTTP(resp, req)

			// 校验状态码是否符合预期
			assert.Equal(t, tc.want.w, resp.Code)
		})
	}
}
