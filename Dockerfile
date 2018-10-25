FROM golang:1.11 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/event_registration
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
# exposed port must be the same as APP_PORT
ENV APP_PORT=3158
ENV APP_NAME=events
ENV APP_DB_DRIVER=postgres
ENV APP_DB_SOURCE=postgres://xgmethyc:h2KYmnYJ15ZezhXWOB5NzwFBCNK55P7D@stampy.db.elephantsql.com:5432/xgmethyc
ENV APP_KEY=secret
# exposed port must be the same as APP_PORT
EXPOSE 3158
ENTRYPOINT ["./app"]

