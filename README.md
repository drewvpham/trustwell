# My REST Service

## Overview
This project provides a REST service for managing event data, including create, read, delete, and list operations. The service uses bearer token authentication and can be deployed using Docker.

## Endpoints

### Create Event
- **Method:** POST
- **Endpoint:** `/events`
- **Description:** Creates a new event.
- **Request Body:**
  ```json
  {
    "type": "shipping",
    "contents": [
      {
        "gtin": "1234",
        "lot": "adffda",
        "bestByDate": "2021-01-13",
        "expirationDate": "2021-01-17"
      }
    ]
  }
