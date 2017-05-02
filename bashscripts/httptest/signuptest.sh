#!/usr/bin/env bash
cd ~/go/src/github.com/SamirMarin/rnspool/
regex="Main|SignUp"
go test -run $regex
