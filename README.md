# sls-phone-validation

> AWS Lambda / API Gateway caching example using Serverless and Go

## Components

**API Layer**  
`GET` endpoint using API Gateway to validate mobile phone numbers.

`/validatePhone?phoneNumber=<mobile phone number>`    

This endpoint will invoke a lambda function written in Go that will send a request to the Numverify API. The response will be cached for 3600 seconds (1 hour) using API Gateway parameter caching integration.

**CI/CD**  
The proyect uses Github Actions to build and test the code.

**Front-End**:  
[![Edit ui-phone-validation](https://codesandbox.io/static/img/play-codesandbox.svg)](https://codesandbox.io/s/3r673qv08q?fontsize=14)
