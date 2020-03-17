# Foobar

Golang Simple Project


## Installation

install golang in [go bro](https://golang.org/doc/install?download=go1.13.6.linux-amd64.tar.gz)

```bash
docker run --name docker-postgres -e POSTGRES_PASSWORD=password -d resistan

docker run -d --name elasticsearch --net somenetwork -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:tag
```


insert in your ~/.bashrc or ~/.bash_profile
```vim
export GOROOT=/usr/local/go #this path of your go lang install
export GOPATH=$HOME/golang-project #this is path for your package and project init
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin

```


run you run the project you must install [depedency](https://github.com/golang/dep)

or install  

```
go get github.com/astaxie/beego
go get -u github.com/beego/bee

```

```bash
bee migrate -driver="postgres" -conn="postgres://postgres:postgres@127.0.0.1:5432/resistan" -dir="database/migrations"

dep ensure
```

config/app.conf
## Usage

```conf
appname = resistor
httpport = 8081
runmode = dev
autorender = false
copyrequestbody = true
EnableDocs = true
sqlconn = 


psqluser=postgres
psqlport=5432
psqldb=resistan 
psqlurls=localhost
psqlpass=postgres

elasticsearch_url=http://localhost:9200/elasticsearch # your elastic url
```


run the project
```
bee run
```
