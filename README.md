# SimpleChatApp

This is a simple real-time chat system built with **Golang** for the backend and **React** for the frontend, using **WebSockets** for real-time communication.

## Project Structure

The project is organized into two main folders:

- **websocket-chat/**: Contains the Go server code, manages WebSocket connections and chat logic.
- **application/**: Contains the React app that allows users to connect and participate in the chat.

## Features

- WebSocket connection
- Send and receive messages in real-time
- Broadcast messages to all connected users

## Requirements

- **Golang** (version 1.23.1 or higher)
- **Node.js** (version 20 or higher)
- **npm** or **yarn**
- **gcc** (10.3.0 or higher)

## Installation and Execution

### Websocket server (Golang)

1. Navigate to the `backend` folder:

   ```bash
   cd backend
   ```

2. Install dependencies and run the Go server:
   ```bash
   go mod tidy
   go run main.go
   ```

The server will be running on `http://localhost:8080`.

### Frontend (React)

1. Navigate to the `frontend` folder:

   ```bash
   cd frontend
   ```

2. Install dependencies:

   ```bash
   yarn install
   ```

3. Start the React app:
   ```bash
   npm start
   ```

The React app will be accessible via `http://localhost:3000`.

## Technologies Used

- **Golang**: Backend, WebSocket management
- **React**: Frontend for user interface
- **WebSocket**: Real-time communication protocol

## Authors

- Ibrahima Sylla [isyll711@gmail.com](mailto:isyll711@gmail.com)

---

Thank you!
