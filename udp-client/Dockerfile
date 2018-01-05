FROM golang:1.9-stretch

RUN go get github.com/DataDog/datadog-go/statsd
RUN mkdir -p src/github.com/scoiatael/datadog-client-issue-poc/client
ADD . src/github.com/scoiatael/datadog-client-issue-poc/client/
RUN go install github.com/scoiatael/datadog-client-issue-poc/client

ENTRYPOINT ["client", "server:8125"]
