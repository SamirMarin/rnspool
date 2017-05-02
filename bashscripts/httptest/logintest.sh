#!/usr/bin/env bash
cd ~/go/src/github.com/SamirMarin/rnspool
regex="Main|Login"
go test -run $regex
