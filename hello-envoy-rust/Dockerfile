FROM rust:latest as builder

WORKDIR /app

#RUN apk add musl-dev

# create a new empty project
RUN cargo init

#COPY ./.cargo .cargo
#COPY ./vendor vendor
COPY Cargo.toml Cargo.lock ./
# build dependencies, when my source code changes, this build can be cached, we don't need to compile dependency again.
RUN cargo build
# remove the dummy build.
RUN cargo clean -p hello-envoy-rust

# build with x86_64-unknown-linux-musl to make it run with alpine.
RUN cargo install --path .

# second stage.
FROM gcr.io/distroless/cc-debian12
COPY --from=builder /usr/local/cargo/bin/* /usr/local/bin
CMD ["./hello-envoy-rust"]