# API Documentation for KubeIT

This page gives a broad overview about all available HTTP Routes in KubeIT and their function.

### Available Routes


| Route            | HTTP-METHOD | DESCRIPTION                                    | LINK |
|------------------|-------------|------------------------------------------------|------|
| /s3/init         | POST        | Initiates a S3 upload                          |[s3init](#s3init) |
| /s3/upload       | GET         | Get a upload URL for file or path              |[s3upload](#s3upload) |
| /s3/finish       | GET         | Finish a previous S3 upload                    |[s3finish](#s3finish) |
| /s3/download     | GET         | Get download URL of an uploaded file           |[s3download](#s3download) |
| /v1/apply        | POST        | Apply a new workflow                           |[applyworkflow](#applyworkflow) |
| /v1/status       | GET         | Get status of one or multiple workflows        |[getstatus](#getstatus) |
| /v1/scheme       | GET         | Get one or all schemes and their parameters    |[getscheme](#getscheme) |
| /v1/createscheme | POST        | Create a new scheme                            |[createscheme](#createscheme) |
| /v1/result       | GET         | Get results of a finished workflow             |[getresults](#getresults) |
| /v1/delete       | GET         | Delete a workflow or a collection of workflows |[deleteworkflow](#deleteworkflow) |