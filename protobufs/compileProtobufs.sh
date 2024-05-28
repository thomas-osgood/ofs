#!/usr/bin/env bash

PROTOROOT=.
TARGETPROTOS=("common" "filehandler")

for i in ${TARGETPROTOS[@]};
do
    # if the target directory does not exist, create one.
    mkdir -p $PROTOROOT/$i
    
    # generate the proper files based on the protobuf definitions.
    protoc -I=. \
        --go_out=$i --go_opt=paths=source_relative --go-grpc_out=$i --go-grpc_opt=paths=source_relative \
        $i.proto

    echo [+] \"$i\" generation complete ...
done