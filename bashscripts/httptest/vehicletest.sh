#!/usr/bin/env bash
cd ~/go/src/github.com/SamirMarin/rnspool/
regex="Main|Vehicle"
go test -run $regex
