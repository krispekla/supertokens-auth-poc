# supertokens-auth-poc

This project is a Proof of Concept (POC) aimed at testing SuperTokens as an authentication solution. It is structured as a monorepo and consists of three main components:

1. **API**: A backend service written in Go.
2. **Frontend**: A React-based frontend application.
3. **Dashboard**: A Node.js-based dashboard for user management.

## API

The API is built using Go and provides backend services for the application. It uses Docker Compose to define containers for PostgreSQL and SuperTokens Core service.

### Setup

1. Navigate to the `api` directory:

    ```sh
    cd api
    ```

2. Copy the `.env.template` to `.env` and fill in the required environment variables:

    ```sh
    cp .env.template .env
    ```

3. Install dependencies:

    ```sh
    go mod download
    ```

4. Run the API:

    ```sh
    go run cmd/api.go
    ```

5. Start the Docker containers:

    ```sh
    docker-compose up
    ```

## Frontend

The frontend is a React application built with Next.js.

### Setup

1. Navigate to the `front` directory:

    ```sh
    cd front
    ```

2. Copy the `.env.template` to `.env` and fill in the required environment variables:

    ```sh
    cp .env.template .env
    ```

3. Install dependencies:

    ```sh
    npm install
    ```

4. Run the development server:

    ```sh
    npm run dev
    ```

## Dashboard

The dashboard is a Node.js application for managing authenticated users. It uses SuperTokens Dashboard recipe to manage users and their roles. It connects directly to the SuperTokens Core service and is intended to be used as internal admin tool.

### Setup

1. Navigate to the `dash` directory:

    ```sh
    cd dash
    ```

2. Copy the `.env.template` to `.env` and fill in the required environment variables:

    ```sh
    cp .env.template .env
    ```

3. Install dependencies:

    ```sh
    npm install
    ```

4. Run the dashboard:

    ```sh
    npm start
    ```

## SuperTokens Integration

This project uses SuperTokens for authentication. The configuration for SuperTokens can be found in the following files:

- API: [api/cmd/api.go](api/cmd/api.go)
- Frontend: [front/src/app/components/supertokensProvider.tsx](front/src/app/components/supertokensProvider.tsx)
- Dashboard: [dash/app.js](dash/app.js)
