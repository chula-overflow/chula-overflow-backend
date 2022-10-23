FROM rust:slim-buster AS builder

WORKDIR /app

COPY apps/auth/ ./

COPY apps/proto ../proto/

RUN apt-get update && \
    apt-get install protobuf-compiler musl-tools -y && \
    rustup target add x86_64-unknown-linux-musl && \
    cargo build --target x86_64-unknown-linux-musl --release 

FROM scratch

WORKDIR /app

COPY --from=builder app/target/x86_64-unknown-linux-musl/release/service .

CMD ["./service"]
