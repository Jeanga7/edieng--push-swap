#!/bin/bash

cd Checker
go build 
mv Checker ../script/build
cd ..
cd push-swap
go build
mv push-swap ../script/build
