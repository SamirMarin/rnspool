#!/usr/bin/env bash
cd ~/go/src/github.com/SamirMarin/rnspool/backend_webservice/
regex="Main|Login"
go test -run $regex