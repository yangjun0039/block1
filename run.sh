#!/bin/bash

rm block1.exe
rm *.db
go build -o block1 *.go
./block1