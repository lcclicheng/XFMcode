package dhttp

import (
	"regexp"
	"strconv"

	"git.woa.com/pdata/datamore/tools/derror"
)

func ParamsRegexp(pattern string, params map[string]string) (bool, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	for _, v := range params {
		if !r.MatchString(v) {
			return false, nil
		}
	}
	return true, nil
}

// ParamsCheck params参数正则表达式校验
func ParamsCheck(params map[string]string) error {
	regOk, err := ParamsRegexp("^[0-9_a-z_A-Z\\,\\.\\&\\=#;_-]+$", params)
	if err != nil {
		return derror.BuildErrMsg(derror.LogicErrCode, "request params regexp err")
	}
	if !regOk {
		return derror.BuildErrMsg(derror.ParaErrCode, "request param regexp format is wrong")
	}

	return nil
}

// ParamsCheckPattern params参数自定义正则表达式校验
func ParamsCheckPattern(pattern string, params map[string]string) error {
	regOk, err := ParamsRegexp(pattern, params)
	if err != nil {
		return derror.BuildErrMsg(derror.LogicErrCode, "request params regexp err")
	}
	if !regOk {
		return derror.BuildErrMsg(derror.ParaErrCode, "request param regexp format is wrong")
	}

	return nil
}

func ParamsCheckEmpty(params ...string) error {
	for i, v := range params {
		if v == "" {
			return derror.BuildErrMsg(derror.ParaErrCode, "param not allowed to be empty "+strconv.Itoa(i))
		}
	}
	return nil
}
