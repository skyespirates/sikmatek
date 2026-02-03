FROM node:20 AS frontend-builder

WORKDIR /client
COPY client/package*.json ./
RUN npm install
COPY client/ .
RUN npm run build   # produces client/dist

FROM golang:1.25.4 AS backend-builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd

FROM golang:1.25.4

WORKDIR /app
COPY --from=backend-builder /app/server .
COPY --from=frontend-builder /client/dist ./client/dist

ENV DSN="root:secret@tcp(db:3306)/sikmatek" \ 
    PORT=3000

CMD ["./server"]