FROM golang:1.18-alpine AS build

# force dockerfile run as a basic user
RUN addgroup -g 1001 -S appuser && adduser -u 1001 -S appuser -G appuser 

COPY go.mod .

ENV GOPATH=""

RUN go mod tidy && go mod verify

COPY . .

# build the app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /bin/app

# copy over the app & user perm requirements
FROM scratch

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group
COPY --from=build /bin/app /bin/app

# this is a web app that runs on port 8050 so let's expose that port
EXPOSE 8050

# run as user
USER 1001

ENTRYPOINT ["/bin/app"]