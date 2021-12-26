FROM golang:1.17-alpine
WORKDIR /app

#Copy go files to container
COPY go.mod ./
COPY go.mod ./
COPY main.go ./
COPY cmd ./cmd
COPY lib ./lib

#Copy application files to container
COPY templates ./templates
COPY recipes.yaml ./

RUN go mod tidy
RUN go build -o /meal-planner

EXPOSE 8080

CMD ["/meal-planner"]