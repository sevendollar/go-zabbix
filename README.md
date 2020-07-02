I written this zabbix API package in Go.
just so everyone out there uses Zabbix can have fun playing with it.

examples
func main() {
	user := "Admin"
	password := "password"
	uri := "http://zabbix/api_jsonrpc.php"

  // start a session
	s, err := newSession(user, password, uri)
	if err != nil {
		log.Fatal(err)
		return
	}
	// make sure to logout properly, just so to have server resources.
	defer s.logout()

	// show the token after successfully logined
	s.showToken()

	// get user infomation
	rlt, r := s.do(newRequest(
		"user.get",
		map[string]interface{}{
			"output": "extend",
		},
	))
	if r != nil {
		log.Fatal(r)
		return
	}

	// make json output pretty.
	jsonPretty(&rlt)

	// print the output.
	fmt.Printf("%s\n", rlt)
}
