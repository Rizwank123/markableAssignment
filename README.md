
# Markable API

The Markable API is a web application designed to manage patient and user data. It provides a set of RESTful APIs for user authentication, patient management, and more.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- User authentication (login and registration)
- CRUD operations for patient records
- Secure access to patient data using JWT
- Swagger documentation for API endpoints

## Technologies Used

- Go (Golang) for backend development
- Echo framework for building the web server
- PostgreSQL for database management
- JWT for secure authentication
- Swagger for API documentation
- Viper for configuration management

## Installation

Follow these steps to set up the project locally:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/markable.git
  
2. **Setup Project:**

   ```bash
   cd markable
   make Setup
3. **Edit .env**

- set database username
- set database password and database name
- edit ``etc/hots`` add
``
127.0.0.1 local.api.markable.co
``
- install  **mkcert**
- ``brew install mkcert`` for macos and for windows install using ``choco install mkcert``
- open terminal run
- ``cd .cert``
- ``mkcert local.api.markable.co``
- run caddy ``caddy run --config ~/Deaktop/markableAssignment/config/caddyFile``

4. **Run migration**

   ```bash
    make migrate
5. **Run Server**

   ```bash
    go run github.com/markable/cmd
    
## License
This project is licensed under the MIT License
 
 

