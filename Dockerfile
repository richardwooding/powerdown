FROM golang:alpine as build
# Redundant, current golang images already include ca-certificates
RUN apk --no-cache add ca-certificates

FROM scratch
# copy the ca-certificate.crt from the build stage
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/powerdown"]
COPY powerdown /