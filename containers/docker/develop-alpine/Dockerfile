FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 https://github.com/ethereum/elementalcore && \
  (cd elementalcore && make geth) && \
  cp elementalcore/build/bin/geth /geth && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /elementalcore && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/geth"]
