FROM ubuntu:22.04

ADD https://github.com/software-mansion/scarb/releases/download/v2.3.1/scarb-v2.3.1-x86_64-unknown-linux-musl.tar.gz /tmp

RUN apt-get update && \
    apt-get install -y wget git && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN tar -zxf /tmp/scarb-v2.3.1-x86_64-unknown-linux-musl.tar.gz -C /tmp/ && \
    mv /tmp/scarb-v2.3.1-x86_64-unknown-linux-musl/bin/*  /usr/local/bin/ && \
    rm -rf /tmp/scarb-v2.3.1-x86_64-unknown-linux-musl*

WORKDIR /app

CMD ["scarb", "cairo-run", "--available-gas", "999999999"]