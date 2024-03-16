FROM golang:alpine as golangbook
RUN addgroup -S golangbook \
  && adduser -S -u 10000 -g golangbook golangbook
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata
RUN echo ${PWD} && ls -lR

FROM scratch
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golangbook /go/bin/golangbook /golangbook
COPY --from=golangbook /etc/passwd /etc/passwd
# COPY .env /golangbook
USER golangbook
ENTRYPOINT ["/golangbook"]