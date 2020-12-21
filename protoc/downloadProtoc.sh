#!/usr/bin/env bash

# This downloads and installs the protobuf compiler depending on the platform

if [ "$(uname)" == "Darwin" ]
then
    # Under Mac OS X platform 
    echo 'Downloading MacOs protobuf compiler'
    curl https://github.com/google/protobuf/releases/download/v3.13.0/protoc-3.13.0-osx-x86_64.zip -o protoc.zip -L
else [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]
    # Under GNU/Linux platform
    echo 'Downloading Linux protobuf compiler'
    curl https://github.com/google/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip -o protoc.zip -L
fi
