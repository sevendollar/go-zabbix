package zabbix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Session struct {
	URI    string `json:uri`
	Token  string `json:"token"`
	client *http.Client
}

// use this method to call the Zabbix API
// reference this URL https://www.zabbix.com/documentation/2.0/manual/appendix/api/api
// to get more infomation.
func (s *Session) Do(r *Request) (result []byte, err error) {
	r.Auth = s.Token
	payload, err := json.Marshal(r)

	req, err := http.NewRequest("POST", s.URI, bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json-rpc")

	s.client = new(http.Client)

	resp, err := s.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// z, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%s\n", z)

	zMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&zMap)
	if err != nil {
		return
	}
	if v := zMap["error"]; v != nil {
		// fmt.Printf("%T", err)
		err = errors.New(v.(map[string]interface{})["data"].(string))
		return
	}

	result, err = json.Marshal(zMap["result"])
	// fmt.Println(zMap["result"].(string))

	return
}

func (s *Session) Login(username, password, uri string) (err error) {
	s.URI = uri

	r := NewRequest("user.login", map[string]string{
		"user":     username,
		"password": password,
	})

	result, err := s.Do(r)
	if err != nil {
		err = fmt.Errorf("Zabbix Error: %v", err)
		return
	}

	var v interface{}
	err = json.Unmarshal(result, &v)
	if err != nil {
		return
	}
	s.Token = v.(string)

	return
}

func (s *Session) Logout() (err error) {
	result, r := s.Do(NewRequest("user.logout", nil))
	if r != nil {
		err = r
		return
	}
	if v := string(result); v != "true" {
		err = fmt.Errorf(v)
		return
	}
	fmt.Println("Logouted!")
	return
}

func (s *Session) ShowToken() {
	fmt.Println("Token:", s.Token)
}
