FROM golang:1.15 as build
WORKDIR /build
COPY go.mod .
COPY loggen.go .
ENV CGO_ENABLED=0
RUN go build -a -ldflags \
  '-extldflags "-static"'

FROM scratch
COPY --from=build /build/loggen /
ENTRYPOINT [ "/loggen" ]
CMD [ "-log-prefix", "loggen-docker" ]
EXPOSE 80
