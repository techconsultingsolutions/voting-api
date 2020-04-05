# voting-api 
# https://www.youtube.com/watch?v=UtNeLfXaBJM

#Good example of go
#https://www.youtube.com/watch?v=ZYgHrcKWH18

#dev machine setup for building go
docker run -it -v ${PWD}:/src --name "golang-dev"  -w "/src" golang /bin/bash

How to run main package ?

go build .
go run .

OR

go run github.com/techconsultingsolutions/voting-api