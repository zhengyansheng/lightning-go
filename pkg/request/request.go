package request

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Post(body []byte, header map[string]string, url string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d", resp.StatusCode)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}

func Delete(url string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d", resp.StatusCode)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}
func Get(url string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d", resp.StatusCode)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}

func GetWithHeader(url string, header map[string]string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote server occer error and http code is %d", resp.StatusCode)
	}
	response, _ := ioutil.ReadAll(resp.Body)
	return response, nil
}
