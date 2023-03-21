# from runtime image
FROM amirhossein21/runtime:latest

# change to app directory
WORKDIR app

# copy go.mod and go.sum files
COPY go.mod go.sum ./

# copy python files
COPY *.py .

# copy go files
COPY *.go .

# build
RUN go build -o ./main

# execute
CMD ./main