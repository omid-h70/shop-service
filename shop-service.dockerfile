# Base go image
# Build in Vendor mod
#
#
#FROM golang:1.18 as builder
FROM golang:1.18-alpine as builder

ARG BINARY_PATH=/app

RUN mkdir /app

COPY . /app

############## Attention
#if you want to change dir, use WORKDIR
#if you use "RUN cd" you must do your other instructions in one "RUN" cycle
#every other instrction that are in another "RUN", are forgotten !!!!!!!!!!!!!!!!
#
#
#

WORKDIR $BINARY_PATH

RUN CGO_ENABLED=0 go build -mod=vendor -o shopService

RUN chmod +x shopService

CMD [ "sh", "-c", "./shopService" ]
#CMD [ "./shopService" ]

###### Production Image - tiny One !
# It was actually very common to have one Dockerfile to use for development (which contained everything needed to build your application),
# and a slimmed-down one to use for production, which only contained your application and exactly what was needed to run it.
# This has been referred to as the “builder pattern”. Maintaining two Dockerfiles is not ideal.
# then moving the built package to a tiny docker image
#
#
#FROM alpine:latest

#RUN mkdir /shopService

#COPY . /shopService

#WORKDIR $BINARY_PATH

#RUN CGO_ENABLED=0 go build -mod=vendor -o shopService

#RUN chmod +x $BINARY_PATH/shopService

#CMD [ "$BINARY_PATH" ]