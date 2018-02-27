package config

type Configuration struct {
  uriApi string
}


func (c Configuration) GetBaseUrl() string{
  return "https://reqres.in/api/users"
}
