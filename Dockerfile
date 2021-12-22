FROM golang:1.17.2

WORKDIR $GOPATH/src/MS_Local

ADD . .
RUN mkdir /home/temp
#RUN GOPROXY="https://goproxy.cn,direct" go build .
EXPOSE 8080

CMD ["/bin/bash", "-c", "./MS_Local"]
