FROM hyperpilot/slow_cooker:influxdb_support

ADD . /go/src/github.com/hyperpilotio/slow_cooker_wrapper

# RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/hyperpilotio/slow_cooker_wrapper

# RUN glide install

RUN go build -o /go/bin/slow_cooker_wrapper main.go

ENTRYPOINT ["/go/bin/slow_cooker_wrapper"]
