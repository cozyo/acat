package test

import (
	"net/http"
	"testing"
	"time"
)

func TestGateway(t *testing.T)  {
	t.Run("测试ping", TestTime)
	t.Run("测试http请求", TestGetSite)
}

func TestTime(t *testing.T)  {
	now := time.Now().Unix()
	t.Log(now)
}

func TestGetSite(t *testing.T)  {
	r := baseUrl + adminLogin
	t.Logf("request url:%s", r)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}
