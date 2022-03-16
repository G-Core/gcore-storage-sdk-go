# G-Core Labs Storage API Golang SDK

The purpose of this project is to cover G-Core Storage API methods related with G-Core Terraform plugin (https://github.com/G-Core/terraform-provider-gcorelabs).

## Internal design of SDK

Since Storage service provides Swagger open docs so we generate an actual version of the client with https://goswagger.io/generate/client.html.
And we extend our wrapper of this client after. 
Read Makefile to get more technical details.

### Status
[![Build Status](https://travis-ci.com/G-Core/gcore-storage-sdk-go.svg?branch=main)](https://travis-ci.com/G-Core/gcore-storage-sdk-go)
