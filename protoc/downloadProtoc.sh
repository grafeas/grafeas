#!/usr/bin/env bash

# This downloads and installs the protobuf compiler depending on the platform

if [ "$(uname)" == "Darwin" ]; then
    # Under Mac OS X platform
    echo 'Downloading MacOs protobuf compiler'
    if [ "$(arch)" == "arm64" ]; then
        echo 'No aarch64 version of protobuf compiler, downloading the x86_64 version'
    fi
    curl https://github.com/google/protobuf/releases/download/v3.13.0/protoc-3.13.0-osx-x86_64.zip -o protoc.zip -L
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    # Under GNU/Linux platform
    echo 'Downloading Linux protobuf compiler'
    if [ "$(uname -m)" == "aarch64" ]; then
        curl https://github.com/google/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-aarch_64.zip -o protoc.zip -L
    else
        curl https://github.com/google/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip -o protoc.zip -L
    fi
elif [ "$(expr substr $(uname -s) 1 5)" == "MINGW" ]; then
    # Under Windows platform
    echo 'Downloading Windows protobuf compiler'
    curl https://github.com/google/protobuf/releases/download/v3.13.0/protoc-3.13.0-win64.zip -o protoc.zip -L
fi
