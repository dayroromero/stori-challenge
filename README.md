## Hey Storier :)

You can find a video presentation about the challenge here --> https://www.loom.com/share/d2b1b9168dd24e4c8934152347840f40?t=175&sid=3997941d-9299-4743-bdb7-9f7f90973664 

# Project Description
This is a Golang desktop application that retrieves files uploaded to an AWS S3 bucket. The project runs on an AWS Lambda function triggered by new S3 file creations. It reads these files, loads them into a database, generates a summary, and sends the summary via email to the user.

## Installation
Dependencies can be installed using `go get`, and the AWS Lambda Go dependency should also be installed.

## Configuration
Configuration requires setting up the following environment variables:
- SENDER_EMAIL
- SENDER_PASSWORD
- SMTP_HOST
- SMTP_PORT
- DBSTRING_CONNECTION

## Usage
To compile the project, use the command:
go build -tags lambda.norpc -o bootstrap main.go

This creates a binary that needs to be compiled and deployed to an AWS Lambda function. Additionally, an S3 event trigger needs to be created.

## Additional Resources
- [AWS Lambda Go Documentation](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
- [AWS Lambda Go Image Guide](https://docs.aws.amazon.com/lambda/latest/dg/go-image.html)
- [AWS Lambda Go Package Guide](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)
- [AWS Lambda Foundation Architecture](https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html)
- [AWS Lambda S3 Integration Example](https://docs.aws.amazon.com/es_es/lambda/latest/dg/with-s3-example.html)

Feel free to reach out if you have any further questions or need assistance!