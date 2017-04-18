#!/usr/bin/env bash
cd ~/go/src/github.com/SamirMarin/rnspool/backend_webservice/
regex="Main|Vehicle"
go test -run $regex
