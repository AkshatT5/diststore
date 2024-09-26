# DistStore API Documentation

DistStore provides a RESTful API for interacting with the distributed file storage system. This document outlines the available endpoints and their usage.

## Base URL

All API endpoints are relative to the base URL of your DistStore instance. For local development, this is typically `http://localhost:8080`.

## Endpoints

### List Files

Retrieve a list of all files in the storage system.

- **URL:** `/api/list`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200
  - **Content:** Array of file names
    ```json
    ["file1.txt", "file2.jpg", "folder/file3.pdf"]
    ```

### Store File

Upload a file to the storage system.

- **URL:** `/api/store`
- **Method:** `POST`
- **Content-Type:** `multipart/form-data`
- **Form Data:**
  - `key`: Unique identifier for the file
  - `file`: The file to be uploaded
- **Success Response:**
  - **Code:** 200
  - **Content:** `File stored successfully`
- **Error Response:**
  - **Code:** 400 BAD REQUEST
  - **Content:** Error message

### Get File

Retrieve a file from the storage system.

- **URL:** `/api/get`
- **Method:** `GET`
- **URL Params:** 
  - `key`: Unique identifier of the file
- **Success Response:**
  - **Code:** 200
  - **Content:** The requested file
- **Error Response:**
  - **Code:** 404 NOT FOUND
  - **Content:** Error message

### Delete File

Delete a file from the storage system.

- **URL:** `/api/delete`
- **Method:** `DELETE`
- **URL Params:**
  - `key`: Unique identifier of the file to delete
- **Success Response:**
  - **Code:** 200
  - **Content:** `{"message": "File deleted successfully"}`
- **Error Response:**
  - **Code:** 404 NOT FOUND
  - **Content:** Error message

## Error Handling

All endpoints may return the following error responses:

- **Code:** 400 BAD REQUEST
  - Returned when the request is malformed or missing required parameters
- **Code:** 500 INTERNAL SERVER ERROR
  - Returned when an unexpected error occurs on the server

Error responses will include a message describing the error.

## Examples

### Uploading a File

```bash
curl -X POST -F "key=example.txt" -F "file=@/path/to/local/file.txt" http://localhost:8080/api/store
```

### Retrieving a File

```bash
curl -O http://localhost:8080/api/get?key=example.txt
```

### Deleting a File

```bash
curl -X DELETE http://localhost:8080/api/delete?key=example.txt
```

### Listing All Files

```bash
curl http://localhost:8080/api/list
```

## Rate Limiting

Currently, there are no rate limits imposed on API requests. However, please be considerate and avoid making too many requests in a short period.

## Versioning

This API is currently in version 1. Future versions will be announced and documented separately.