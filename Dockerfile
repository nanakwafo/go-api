FROM golang:1.22.2-alpine3.19

#sey working dirctory
WORKDIR /app

#copy the source code
COPY . .    
#download dependencies
RUN go mod tidy

#build the application
RUN go build -o api .

#expose the port        

EXPOSE 8000 


#run the executable
CMD [ "./api" ]
