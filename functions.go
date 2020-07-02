package zabbix

import "encoding/json"

func newRequest(method string, params interface{}) *Request {
	if params == nil {
		params = make(map[string]string)
	}

	return &Request{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Auth:    "",
		ID:      1,
	}
}

func newSession(username, password, uri string) (*Session, error) {
	s := new(Session)

	err := s.login(username, password, uri)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func jsonPretty(jUgly *[]byte) (err error) {
	var x interface{}

	r := json.Unmarshal(*jUgly, &x)
	if r != nil {
		err = r
		return
	}

	(*jUgly), r = json.MarshalIndent(x, "", "  ")
	if r != nil {
		err = r
		return
	}
	return
}
