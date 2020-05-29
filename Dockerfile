FROM golang:1.14.2-alpine AS builder

RUN apk add --no-cache ca-certificates git
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

WORKDIR /src
ADD . /src
RUN cd /src && ./scripts/build.sh

FROM scratch AS final

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/build/aker /app

EXPOSE 3300
USER nobody:nobody

ENTRYPOINT [ "/app/unbans" ]
