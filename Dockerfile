FROM golang:1.14.7

RUN apt-get update -qq && \ 
    apt-get install --yes -qq zip

