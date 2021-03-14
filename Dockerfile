FROM golang:1.10 AS builder

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/pradipsilwal/kubernetes-yaml-generator/

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

#Download the dependencies
RUN go get -u github.com/gorilla/mux && go get go.mongodb.org/mongo-driver/mongo 

RUN CGO_ENABLED=0 GOOS=linux go build -o /appbuild api/main.go


# #final image ---------------------
FROM alpine:3.12


# #copy the go build from previous image 
COPY --from=builder /appbuild /
# 

# #set the work directoru
WORKDIR /

#Expose the port for the application
EXPOSE 8080

RUN pwd

#Run the app
CMD ["./appbuild"]