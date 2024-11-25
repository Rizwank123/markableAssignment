# Markable API

The Markable API is a web application designed to manage patient and user data. It provides a set of RESTful APIs for user authentication, patient management, and more.

## Table of Contents

- [Overview](#overview)
- [Access Control](#access-control)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [License](#license)

## Overview

The Markable API provides the following core features:
- User authentication (login and registration)
- Role-based access control (RBAC)
- CRUD operations for patient records
- Secure access to patient data using JWT
- Swagger documentation for API endpoints

## Access Control

The API implements role-based access control with the following permissions:

| Role          | View Patients | Create Patients | Update Patients | Delete Patients |
|---------------|---------------|-----------------|-----------------|-----------------|
| Doctor        | ✅            | ✅              | ✅              | ✅              |
| Receptionist  | ✅            | ✅              | ✅              | ✅              |
| Nurse         | ✅            | ❌              | ❌              | ❌              |

## Technologies Used

- **Backend Framework**: Go (Golang) with Echo framework
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Documentation**: Swagger
- **Configuration**: Viper

## Getting Started

### Prerequisites

- Go 1.19 or higher
- PostgreSQL
- mkcert (for SSL certificates)
- Caddy (for reverse proxy)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/markable.git
   cd markable
   ```

2. **Setup Project:**
   ```bash
   make Setup
   ```

3. **Run migrations:**
   ```bash
   make migrate
   ```

4. **Start the server:**
   ```bash
   go run github.com/markable/cmd
   ```

### Configuration

1. **Environment Setup**
   - Create and configure your `.env` file with:
     - Database username
     - Database password
     - Database name

2. **Local Domain Configuration**
   - Edit `/etc/hosts` and add:
     ```
     127.0.0.1 local.api.markable.co
     ```

3. **SSL Certificate Setup**
   - Install mkcert:
     ```bash
     # MacOS
     brew install mkcert
     
     # Windows
     choco install mkcert
     ```
   - Generate certificates:
     ```bash
     cd .cert
     mkcert local.api.markable.co
     ```

4. **Start Caddy Server**
   ```bash
   caddy run --config ~/Desktop/markableAssignment/config/caddyFile
   ```

## API Documentation

API documentation is available via Swagger UI when the server is running.
Access it at: `https://local.api.markable.co/swagger/index.html`

## License

This project is licensed under the MIT License.
 

