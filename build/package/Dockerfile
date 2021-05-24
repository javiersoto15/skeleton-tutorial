FROM artifactory-rco.centene.com/pots_docker_production_local_rco/centene/golang:1.16.2-stretch AS builder

 # Make Tag
ARG tag
ENV TAG $tag

LABEL maintainer="CLAIMS RTR Support <CLAIMS_RTR_Support@CENTENE.COM>"

ENV GO111MODULE=on 

WORKDIR /go/src/gitlab.centene.com/rtr/claims/api/rtr-claims-hybrid-wrapper
 
COPY . .

ENV TZ America/Chicago

EXPOSE 8080

WORKDIR /go/src/gitlab.centene.com/rtr/claims/api/rtr-claims-hybrid-wrapper/cmd/claims-wrapper
 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


#Second stage for a smaller image
FROM artifactory-rco.centene.com/pots_docker_production_virtual_rco/centene/alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /root/
EXPOSE 8080
COPY --from=builder /go/src/gitlab.centene.com/rtr/claims/api/rtr-claims-hybrid-wrapper/cmd/claims-wrapper/app .
CMD ["./app"]