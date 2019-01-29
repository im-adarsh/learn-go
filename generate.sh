#!/usr/bin/env bash

protoc calculator/proto/calculator.proto --go_out=plugins=grpc:.