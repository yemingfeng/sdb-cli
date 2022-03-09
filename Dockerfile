FROM ubuntu:18.04

COPY cli cli

ENTRYPOINT "./cli"