# go-rest-api

## Overview

### BACKEND
This project delves into the creation of a RESTful API using Golang, focusing on structured design and best practices.

#### Application Structure
- The application is meticulously structured with distinct layers, including models and services.
- Utilization of *GORM*, a robust Go ORM library, facilitates seamless object-data mapping.

#### Routing
- *GORILLA/MUX* is employed for defining and managing the application's HTTP routes efficiently.

#### Configuration
- *GODOTENV* is incorporated to manage and reference essential application environment variables.

### INFRASTRUCTURE
The project incorporates Docker for containerization, streamlining the setup of a PostgreSQL container for the database.

## Getting Started

### Docker Setup
- The provided docker-compose file simplifies the process of setting up the PostgreSQL container.

#### Linux
```sh
docker-compose up -d
```

#### Windows
```sh
winpty docker-compose up -d
```

These steps ensure a quick and consistent environment setup, enabling seamless development and deployment of the Golang REST API.
