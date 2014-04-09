#!/bin/bash

go test -c
./genpasswd.test -test.cpuprofile=genpasswd.prof -test.bench=.
go tool pprof genpasswd.test genpasswd.prof 
rm genpasswd.test genpasswd.prof
