# ⚠️ DEPRECATED ⚠️

This repository is now **DEPRECATED** and is no longer maintained.

All storage functionality has been moved into the new, unified G-Core Go SDK:

👉 **New SDK:** [**`gcore-go`**](https://github.com/G-Core/gcore-go)

Please migrate your projects to the new SDK for ongoing support and updates.

**Examples of how to use the storage features in the new SDK can be found [here](https://github.com/G-Core/gcore-go/tree/main/examples/storage).**

---

# G-Core Labs Storage API Golang SDK

The purpose of this project is to cover G-Core Storage API methods related with G-Core Terraform plugin (https://github.com/G-Core/terraform-provider-gcore).

## Internal design of SDK

Since Storage service provides Swagger open docs so we generate an actual version of the client with https://goswagger.io/generate/client.html.
And we extend our wrapper of this client after. 
Read Makefile to get more technical details.

### Status
[![Build Status](https://travis-ci.com/G-Core/gcore-storage-sdk-go.svg?branch=main)](https://travis-ci.com/G-Core/gcore-storage-sdk-go)
