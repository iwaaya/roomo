FROM golang

RUN go get -u github.com/golang/dep/...

WORKDIR /go/src/github.com/iwaaya/roomo
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . /go/src/github.com/iwaaya/roomo

RUN make build && \
    cp ./build/* /bin/

CMD ["/bin/roomo-manager", "--config", "/etc/roomo/config.yaml"]
