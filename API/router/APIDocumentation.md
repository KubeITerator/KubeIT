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


# s3init

Initiates a S3 upload. (HTTP-METHOD: POST) 
Request JSON format:

```json
{
  "filename": "NAMETOFILE",
  "multi": true
  
}
```

Response JSON format:

```json
{
  "passkey": "PASSKEY"
}
```

The S3 init request accepts a desired file name (`NAMETOFILE`) and the boolean variable `multi`. Multi must specify either a single, or multi-parted S3 upload.
The Request returns JSON data containing a `passkey` that is used for following requests.

# s3upload

Upload request that is followed by the (s3init)[#s3init] request. (HTTP-METHOD: GET)

Query parameter: `key`, references a passkey returned by (s3init)[#s3init]

Example:

```
example.com/s3/upload?key=PASSKEY
```

Returns a JSON formatted object:

```json
{
  "url": "S3-UPLOAD-URL"
}
```

`url` contains the upload URL that can be populated with data using a HTTP PUT Request.

# s3finish

Finishes a previous initiated S3-upload. (HTTP-METHOD: GET)

Query parameter: `key`, references a passkey returned by (s3init)[#s3init]

Example:

```
example.com/s3/finish?key=PASSKEY
```

Returns a HTTP response code `200` on success.

# s3download

Requests download URL for a previously uploaded file.

Query parameter: `key`, references a passkey returned by (s3init)[#s3init]

Example:

```
example.com/s3/download?key=PASSKEY
```

Returns a JSON formatted object:

```json
{
  "url": "S3-DOWNLOAD-URL"
}
```

`url` contains the download URL that can be used to access the uploaded data.


# applyworkflow

Creates a new workflow based on a available scheme. (HTTP-METHOD: POST)

Request JSON format:

```json
[
  {"parametername.parametercategory": "VALUE"},
  {"parametername2.parametercategory2": "VALUE"}
]
```

The request consists of a list of key value pairs that refer to kubeit specific parameters and their desired values.
In a scheme kubeit specific parameters contain the leading kubeit tag. (example: `{{kubeit.category.name}}`) 

Response JSON format:

In case of missing parameters, a list of missing parameters is returned:

```json
{
"status":  "Missing parameters",
"missing": []
}

```

In case of success:
```json
{
"status": "Successful",
"wfname": "WORFLOWNAME"
}
```

# getstatus




