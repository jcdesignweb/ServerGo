package config

type Configuration struct {
  uriApi string
}


func (c Configuration) GetBaseUrl() string {
  return "https://reqres.in/api/users"
}

func (c Configuration) GetPrefixApi() string {
  return "/api/v1"
}

func (c Configuration) GetHost() string {
  return "localhost"
}
func (c Configuration) GetPort() int {
  return 9000
}

