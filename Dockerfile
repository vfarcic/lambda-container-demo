FROM public.ecr.aws/bitnami/golang:1.15 AS build
COPY main.go .
RUN go get -d -v github.com/aws/aws-lambda-go/lambda
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM public.ecr.aws/lambda/go:1
COPY --from=build /go/app ${LAMBDA_TASK_ROOT}
CMD [ "app" ]
