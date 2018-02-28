package boot

import "strings"

var (
	APP_NAME = "App Name"

	APP_VERSION = "1.0.0"

	APP_API_URI = "https://reqres.in/"

	DEV_ENV = "dev"

	STAGING_ENV = "staging"

	PROD_ENV = "prod"

	REPLACER *strings.Replacer = strings.NewReplacer(".", "_")

	APP_CONFIG_PREFIX = `APP`

	APP_CONFIG_NAME = `application`

	CONFIG_PATH = `config`

	CONFIG_FILE_TYPE = `yaml`


)