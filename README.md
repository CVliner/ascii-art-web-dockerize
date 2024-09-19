### Ascii-art-web-dockerize

Ascii-art-web is project similar to ascii-art, but working on your localhost.

### Running:

In terminal, just to start program without docker:

`go run  (. or main.go)` 
or 
`make run`

go to the server by accessing the link below in web browser

    http://localhost:4747

Or, in terminal, to start program with docker:
Firstly, build image from Dockerfile:

`docker build -t [image name] .`

Run container:
`docker run -p 4747:4747 [image name]`

And start in browser:

`http://localhost:4747`
