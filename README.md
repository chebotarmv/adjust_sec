# Repositories events 

this program requests urls and return MD5 hash of response

## Usage 

To use program  build it with command 
```
go build 
```

then execute program ```./adjust``` with command line arguments

-parallel - Parallel set max number of parallel execution. optional - example [-parallel=5] max - 100, default - 10

and the list of urls to process

## Tests 

To run test use command ```go test``` in directory