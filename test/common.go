package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type HttpCommonRsp struct {
	Code int			`json:"code"`
	Data interface{}	`json:"data"`
	Msg string			`json:"msg"`
}

const (
	SuccessBusinessCode = 200
)

func commonTest(url string, req * http.Request, t *testing.T)  {
	t.Logf("request token=%v", qToken)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("req url: %v status : %v", url, rsp.Status)
	if rsp.StatusCode != http.StatusOK {
		t.Error("StatusCode != 200")
		return
	}
	body, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("req url: %v body : \n%s", url, body)
	var obj HttpCommonRsp
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		t.Error(err)
		return
	}
	if obj.Code != SuccessBusinessCode {
		t.Errorf("business code != %v", SuccessBusinessCode)
		t.Errorf("obj == %+v, obj", obj)
		return
	}
}