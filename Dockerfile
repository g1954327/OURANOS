FROM golang:1.19-bullseye as builder

ADD . /go/ouranos
WORKDIR /go/ouranos
RUN make clean && make && adduser --disabled-login --disabled-password nonroot

FROM scratch

COPY --from=builder /go/ouranos/ouranos /usr/bin/ouranos
COPY --from=builder /etc/passwd /etc/passwd
USER nonroot

ENTRYPOINT [ "/usr/bin/ouranos" ]
