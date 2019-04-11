#!/bin/bash

export GOOS='linux'
export GOARCH='amd64'

cd ../
project_path=$(pwd)
project_name="${project_path##*/}"

rm -rf ./bin/*

mkdir -p ./bin/${project_name}/configs
cp -rf ./configs/ ./bin/${project_name}/configs/
rm ./bin/${project_name}/configs/config.go
cd cmd/${project_name}
go build -o ../../bin/$project_name/${project_name}
cd ../../
tar -zcvf ./${project_name}.tar.gz -C ./bin/ $project_name
rm -rf ./bin/*
mv ./${project_name}.tar.gz ./bin/
