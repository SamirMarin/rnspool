#!/usr/bin/env bash
cd ~/go/src/github.com/SamirMarin/rnspool/
regex="Main|SetUpRide"
go test -run $regex
