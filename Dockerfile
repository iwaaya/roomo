FROM golang

RUN go get -u github.com/golang/dep/...

WORKDIR /go/src/github.com/iwaaya/roomo
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . /go/src/github.com/iwaaya/roomo

RUN make build && \
    cp ./build/roomo-api /bin/ && \
    cp ./build/parse /bin/ && \
    mkdir /etc/roomo && \
    cp docker/config.yaml.tpl /etc/roomo/ && \
    /bin/parse 

CMD ["/bin/roomo-api", "--config", "/etc/roomo/config.yaml"]
