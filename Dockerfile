FROM golang:1.13.6

MAINTAINER Ganang Wahyu Wicaksono gananggww@gmail.com

# installation basic lib
RUN go get github.com/astaxie/beego && go get github.com/beego/bee
RUN go get -u github.com/golang/dep/cmd/dep
# RUN bee migrate -driver="postgres" -conn="postgres://postgres:postgres@127.0.0.1:5432/resistan" -dir="database/migrations"

# no C compile 
RUN CGO_ENABLED=0 go install -a std

ENV REPO go/src/resistor
ENV IMG_APP /$REPO

RUN mkdir -p $IMG_APP

# set path project in images
WORKDIR $IMG_APP
# copy project in local (1st point) to WORKDIR in images (2ndpoint) 
COPY . .

# uncommend if you no init yet
# RUN dep init
RUN dep ensure

# build newest apps
RUN CGO_ENABLED=0 go build -ldflags '-d -w -s'

# port in images
EXPOSE 8080

# run the program
CMD ['bee migrate -driver="postgres" -conn="postgres://postgres:postgres@127.0.0.1:5432/resistan" -dir="database/migrations"']
CMD ["./resistor"]