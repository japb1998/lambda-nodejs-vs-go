build:
	GOARCH=arm64 GOOS=linux go build  -tags lambda.norpc -o bin/bootstrap ./go-lambda
	zip -j bin/go-lambda.zip bin/bootstrap

deploy:
	serverless deploy --aws-profile personal

load-js:
	artillery run --output js-lambda-report.json artillery-js.yml
load-go:
	artillery run --output go-lambda-report.json artillery-go.yml
	artillery report go-lambda-report.json