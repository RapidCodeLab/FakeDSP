module github.com/RapidCodeLab/fakedsp

go 1.19

replace github.com/RapidCodeLab/fakedsp/internal/server => ./internal/server

replace github.com/RapidCodeLab/fakedsp/pkg/ads_db => ./pkg/ads_db

replace github.com/RapidCodeLab/fakedsp/pkg/config => ./pkg/config

require (
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/haxqer/vast v0.0.0-20220726014210-128aa4704401
	github.com/ilyakaznacheev/cleanenv v1.3.0
	github.com/mxmCherry/openrtb/v16 v16.0.0
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
