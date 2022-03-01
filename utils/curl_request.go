package utils

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CurlRequestData struct {
	Method        string
	FullURL       string
	RequestBody   interface{}
	RequestQuery  map[string]string
	CustomHeaders map[string]string
	CustomCookies []*http.Cookie
	Context       context.Context
	TimeoutSecond int
}

var tr = &http.Transport{
	TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	DisableKeepAlives: true,
	Dial: (&net.Dialer{
		KeepAlive: time.Second * 60,
	}).Dial,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
}

func (reqData *CurlRequestData) DoRequest() (statusCode int, respBody []byte, respCookie []*http.Cookie, err error) {
	timeout := reqData.TimeoutSecond
	if reqData.TimeoutSecond == 0 {
		timeout = 20
	}

	client := &http.Client{
		Timeout:   time.Second * time.Duration(timeout),
		Transport: tr,
	}
	urlEncoded := false
	var reqBody *bytes.Buffer
	var urlEncodedPayload *strings.Reader
	switch body := reqData.RequestBody.(type) {
	case nil:
		reqBody = bytes.NewBuffer(nil)
	case []byte:
		reqBody = bytes.NewBuffer(body)
	case string:
		reqBody = bytes.NewBuffer([]byte(body))
	case url.Values:
		urlEncoded = true
		urlEncodedPayload = strings.NewReader(body.Encode())
	default:
		var err error
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return 0, nil, nil, err
		}
		reqBody = bytes.NewBuffer(bodyJSON)
	}
	var req *http.Request
	if urlEncoded == true {
		req, err = http.NewRequest(reqData.Method, reqData.FullURL, urlEncodedPayload)
		if err != nil {
			log.Println(err.Error())
			return 0, nil, nil, err
		}
	} else {
		req, err = http.NewRequest(reqData.Method, reqData.FullURL, reqBody)
		if err != nil {
			log.Println(err.Error())
			return 0, nil, nil, err
		}
	}

	reqQuery := req.URL.Query()
	if reqData.RequestQuery != nil {
		for k, v := range reqData.RequestQuery {
			reqQuery.Add(k, v)
		}
		req.URL.RawQuery = reqQuery.Encode()
	}

	if reqData.Context != nil {
		req = req.WithContext(reqData.Context)
	}

	for k, v := range reqData.CustomHeaders {
		if k == "Host" {
			req.Host = v
		} else {
			req.Header.Add(k, v)
		}
	}

	if len(reqData.CustomCookies) != 0 {
		for _, v := range reqData.CustomCookies {
			req.AddCookie(v)
		}
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return 0, nil, nil, err
	}
	defer res.Body.Close()

	respBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return 0, nil, nil, err
	}

	return res.StatusCode, respBody, res.Cookies(), nil
}

func (reqData *CurlRequestData) StopRedirectRequest() (statusCode int, respbody []byte, redirectUrl string, err error) {
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}
	redirect := false
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirect = true
		redirectUrl = req.URL.String()
		return errors.New("stop redirect")
	}
	urlEncoded := false
	var reqBody *bytes.Buffer
	var urlEncodedPayload *strings.Reader
	switch body := reqData.RequestBody.(type) {
	case nil:
		reqBody = bytes.NewBuffer(nil)
	case []byte:
		reqBody = bytes.NewBuffer(body)
	case string:
		reqBody = bytes.NewBuffer([]byte(body))
	case url.Values:
		urlEncoded = true
		urlEncodedPayload = strings.NewReader(body.Encode())
	default:
		var err error
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return 0, nil, "", err
		}
		reqBody = bytes.NewBuffer(bodyJSON)
	}
	var req *http.Request
	if urlEncoded == true {
		req, err = http.NewRequest(reqData.Method, reqData.FullURL, urlEncodedPayload)
		if err != nil {
			log.Println(err.Error())
			return 0, nil, "", err
		}
	} else {
		req, err = http.NewRequest(reqData.Method, reqData.FullURL, reqBody)
		if err != nil {
			log.Println(err.Error())
			return 0, nil, "", err
		}
	}

	reqQuery := req.URL.Query()
	if reqData.RequestQuery != nil {
		for k, v := range reqData.RequestQuery {
			reqQuery.Add(k, v)
		}
		req.URL.RawQuery = reqQuery.Encode()
	}

	if reqData.Context != nil {
		req = req.WithContext(reqData.Context)
	}

	for k, v := range reqData.CustomHeaders {
		if k == "Host" {
			req.Host = v
		} else {
			req.Header.Add(k, v)
		}
	}

	if len(reqData.CustomCookies) != 0 {
		for _, v := range reqData.CustomCookies {
			req.AddCookie(v)
		}
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		if redirect == true {
			err = errors.New("stop redirect")
		}
		return 0, nil, redirectUrl, err
	}
	defer res.Body.Close()

	respbody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return 0, nil, "", err
	}

	return res.StatusCode, respbody, req.URL.String(), nil
}
