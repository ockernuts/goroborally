# goroborally
Roborally game in go

# used tools
* swagger : ´go get -u github.com/go-swagger/go-swagger/cmd/swagger´
* dep: installed from (https://github.com/golang/dep/releases/tag/v0.3.2) . This is used to pull in dependencies into a "vendor" structure. 
* delve: the go debugger
* godef
* ide: Visual Studio Code


# Swagger model & API generation
* used: ´swagger generate server -f ./swagger.json -A roborally --flag-strategy pflag´ to generate the server code. 