# golang image where workspace (GOPATH) configured at /go. 
FROM golang 
# Copy the local package files to the container’s workspace. 
ADD . /go/src/github.com/goa-fhir/server 
# Setting up working directory 
WORKDIR /go/src/github.com/goa-fhir/server 
# Get godeps for managing and restoring dependencies 
RUN go get github.com/kardianos/govendor
# Restore godep dependencies 
RUN govendor install +local
# RUN godep restore 
# Build the taskmanager command inside the container. 
RUN go install github.com/goa-fhir/server 
# Run the taskmanager command when the container starts. 
ENTRYPOINT /go/bin/goa-fhir 
# Service listens on port 8080. 
EXPOSE 8080