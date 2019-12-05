# Start from go base image
FROM golang:1.13-alpine

# Install git for fetching dependencies
RUN apk update && apk add --no-cache git && apk add wget && apk add --update netcat-openbsd

# Set ENV variables
ENV REPO_URL=github.com/sauravgsh16/graphql-go

ENV GOPATH=/app

ENV APP_PATH=${GOPATH}/${REPO_URL}

ENV WORK_PATH=${APP_PATH}

# Label
LABEL maintainer="Saurav Ghosh <sauravgsh16@gmail.com>"

# Set working directory
WORKDIR ${WORK_PATH}

# Copy all the application files
COPY . .

# Download all dependecies
RUN go mod download

# Build go project
RUN go build -o graphql-go .

# EXPOSE PORT
EXPOSE 8080

# Get wait-for
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for

# Change permission
RUN chmod +x wait-for

CMD ["sh", "wait-for", "db:5432", "--", "./graphql-go"]