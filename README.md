# Emrello

A simple todo app that helps users organize their tasks with ease.  
This project was developed as a case study for **Privia Security**.

🔗 [Live Site](https://www.emrello.com)

---

## 📚 Table of Contents

- [Predefined Users](#predefined-users)
- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Fullstack Architecture](#fullstack-architecture)
- [Frontend](#frontend)
- [Backend](#backend)
- [Project Structure](#project-structure)
- [Deployment](#deployment)
- [License](#license)

---

### Predefined Users

Below is the list of predefined users in the system:

| **Username** | **Password**   | **Role** |
|--------------|----------------|----------|
| user1        | 1234           | user     |
| user2        | 1234           | user     |
| user3        | 1234           | user     |
| user4        | 1234           | user     |
| user5        | 1234           | user     |
| user6        | 1234           | user     |
| Emre         | admin1234      | admin    |
| Berat        | admin1234      | admin    |
| Simge        | admin1234      | admin    |

---

## 🧠 Overview

**Emrello** is a fullstack web application built with Go using the Fiber framework on both frontend and backend.  
It offers a seamless task management experience where users can create and manage multiple todo lists and tasks. The system also includes admin-level functionality for viewing other users' todo lists.

---

## ✅ Features

- Create and delete todo lists
- Add and delete tasks within a list
- Edit task content
- Toggle task completion status
- Maintain multiple separate todo lists
- Admin user can view all users’ todo lists (read-only)

---

## 🛠 Tech Stack

- **Language**: Go
- **Framework**: Fiber
- **Frontend**: Server-Side Rendered (SSR) using Go's `html/template`
- **Backend**: RESTful API using Fiber
- **Deployment**: [Render.com](https://render.com)

---

## 🏗 Fullstack Architecture

The application follows a **modular architecture**, keeping the frontend and backend logically separated while sharing the same codebase and Fiber server.

- **Frontend**: Uses HTML templates rendered on the server with Go.
- **Backend**: Handles business logic, data operations, and serves frontend pages.
- **Routing**: Handled through Fiber, with clearly separated routes for API endpoints.
- **Data Storage**: In-memory storage as a mock service, but can easily be replaced with an actual database server. (PostgreSQL, MongoDB etc.)
- **Mono-repo**: The frontend and backend code is stored in a single repository, allowing for easy sharing of common files.

---

## 🎨 Frontend

The frontend is entirely written using **Go’s html/template package**, avoiding JavaScript altogether. The user experience is achieved through classic HTML form submissions and conditional rendering.

### Highlights:
- **MVC Pattern**: The frontend is structured using the **Model-View-Controller (MVC)** pattern:
  - **Model**: Shared domain models represent todo lists and tasks.
  - **View**: HTML templates render the user interface based on data and state.
  - **Controller**: Handlers receive form inputs, process logic, and return the appropriate view with updated data.
- **No JavaScript**: All interactions are handled server-side.
- **State Flags**: State such as "editing mode" is controlled via values injected into templates (e.g., `IsEditing`, `EditTaskID`).
- **Form Handling**: All actions (add, edit, delete) are managed via HTML form submissions to backend routes.
- **Middleware**: 
  - **Logging** middleware is used to log all HTTP requests.
  - **Authentication** middleware ensures users are logged in before accessing protected pages. Unauthenticated users are redirected to `/login`.
  - A **custom 404 handler** provides a clean fallback page for undefined routes.
- **Static File Hosting**: CSS and SVG files are hosted at the root (`/`)
- **Automatic Service Switching**: The frontend automatically selects between `MockServices` and `ApiServices` based on the `ENVIRONMENT` environment variable. In the "dev" environment, `MockServices` are used for local testing, requiring no real backend. In the "prod" environment, `ApiServices` are used to connect to the live backend. This setup allows for seamless local development and easy transition to the real backend in production.
- **Docker**: The frontend is containerized separately, ensuring easy setup and consistent environments across different machines. It runs independently of the backend, allowing for seamless local testing with mock services and real backend integration in production.
- **Responsive Design**: Layout is responsive and user-friendly across devices.

### Controller Endpoints

- **`GET /`**  
  - **Purpose**: Returns the home page, which contains the todo lists.  
    - **For Admins**: Displays the admin panel showing all users’ todo lists.  
    - **For Regular Users**: Displays the user's own todo lists.  
  - **Query Parameter**: 
    - If the `create` query parameter is provided, a panel for creating a new todo list will be displayed.

- **`GET /health`**  
  - **Purpose**: Returns a simple "healthy" response to check the application's health status. This is useful for monitoring and ensuring the app is running correctly.

- **`GET /login`**  
  - **Purpose**: Returns the login page where users can authenticate to access the app.

- **`POST /login`**  
  - **Form values**:  
    - `username` (required)  
    - `password` (required)  
  - **Purpose**: Sends the user credentials to the backend for authentication.  
    - If successful, the user will receive a JWT token stored in HttpOnly cookies to remember their login status.  
    - If authentication fails, an error message will be shown on the login page.

- **`POST /logout`**  
  - **Purpose**: Deletes the JWT token from the HttpOnly cookies, effectively logging the user out, and redirects back to the `/login` page.

- **`GET /todo-list`**  
  - **Query Parameters**:  
    - `id` (required): The ID of the todo list.  
    - `edit_name` (optional): Enables the view to edit the todo list name.  
    - `edit_todo_task_id` (optional): Enables the view to edit the content of a specific todo task.  
  - **Purpose**: Returns the todo list page for the specified list.  
  - **Notes**:  
    - Regular users can only view their own todo list.  
    - Admins can view other users' todo lists but cannot edit them.

- **`POST /todo-list`**  
  - **Form Values**:  
    - `user_id` (required): The ID of the user who owns the new todo list.  
    - `name` (required): The name of the new todo list.  
  - **Purpose**: Creates a new todo list with the specified name for the user with the specified ID.  
  - **Notes**: 
    - After the todo list is created, the user will be redirected to the newly created todo list page.

- **`POST /todo-list/patch`**  
  - **Form Values**:  
    - `id` (required): The ID of the todo list to be renamed.  
    - `name` (required): The new name for the todo list.  
  - **Purpose**: Renames the specified todo list.  
  - **Notes**: 
    - Only the authorized user (the owner) can delete the todo list. Unauthorized attempts will result in a `403 Forbidden` response.
    - `POST /todo-list/patch` is used instead of `PATCH /todo-list` because HTML form submission does not support the PATCH method.

- **`POST /todo-list/delete`**  
  - **Form Values**:  
    - `id` (required): The ID of the todo list to be deleted.  
  - **Purpose**: Deletes the specified todo list.  
  - **Notes**: 
    - Only the authorized user (the owner) can delete the todo list. Unauthorized attempts will result in a `403 Forbidden` response.  
    - `POST /todo-list/delete` is used instead of `DELETE /todo-list` because HTML form submission does not support the DELETE method.

- **`POST /todo-task`**  
  - **Form Values**:  
    - `list_id` (required): The ID of the list that this task belongs to.  
    - `content` (required): The content of the task.  
  - **Purpose**: Adds a new task with the given content to the specified todo list.  
  - **Notes**:  
    - Only authorized users can add tasks to their own todo lists.  
    - Admins are not allowed to add tasks to other users’ lists.

- **`POST /todo-task/patch`**  
  - **Form Values**:  
    - `action` (required): Specifies the action to perform on the task.  
      - `"toggle"`: Toggles the completion status of the task.  
      - `"edit"`: Enables the edit panel for editing the task content.  
  - **Purpose**: Handles updates to a task, either by toggling its completion state or enabling edit mode.  
  - **Notes**:  
    - This route uses `POST` instead of `PATCH` because HTML forms do not support the `PATCH` method.  
    - Only authorized users can modify their own tasks.  
    - Admin users are not permitted to modify other users’ tasks.

- **`POST /todo-task/delete`**  
  - **Form Values**:  
    - `id` (required): The ID of the task to be deleted.  
  - **Purpose**: Deletes the specified task from the todo list.  
  - **Notes**:  
    - This route uses `POST` instead of `DELETE` because HTML form submissions do not support the `DELETE` method.  
    - Only the owner of the task (authorized user) is allowed to delete it.  
    - Admin users can view other users’ tasks but are not permitted to delete them.

---

## ⚙️ Backend

The backend is a RESTful API built with **Go and Fiber**, following **Clean Architecture** principles.

### Layers:

### **Handlers**

The **Handlers** layer is responsible for processing and validating incoming HTTP requests. In this layer:

- **Request Parsing**: We extract and parse request parameters, form values, and body data.
- **Request Validation**: Input data is validated to ensure it meets the required criteria, such as checking required fields, types, or authorization.
- **Usecase Interaction**: After parsing and validating the request, the handler invokes the corresponding usecase (business logic) by calling the usecase interface, which has been injected into the handler. This ensures the separation of concerns, keeping the handler lightweight and focused on request processing.

### **Usecases**

The **Usecases** layer represents the heart of the business logic in the application. It is where the core rules of the application are implemented. 

- **Business Logic**: This layer is responsible for executing the operations that fulfill the application's requirements, like task creation, list modification, or any other core functionality.
- **Separation from Handlers**: Usecases are agnostic of the handlers from which they are called, focusing solely on business rules and application logic.
- **Repository Interaction**: The usecases interact with the data layer (repositories) to fetch or persist data. The repositories are injected into the usecase as an interface, promoting loose coupling and making it easier to swap implementations (e.g., switching from in-memory storage to a database).
- **Pure Logic**: The usecases do not handle HTTP requests, responses, or validation. Their only job is to execute the necessary business logic by using the data provided by the handler.

This separation ensures that the core logic of the application remains testable, modular, and easy to maintain.

### **Repositories**

The **Repositories** layer is part of the domain layer and serves as the interface between the business logic (usecases) and the data storage.

- **Data Access**: This layer contains the necessary interfaces for accessing and manipulating data, whether it's reading from or writing to a database, in-memory storage, or any other form of persistence.
- **Domain Model**: Each repository is designed to handle a single domain model, ensuring that the data access logic is scoped to a specific entity (e.g., Todo List, Todo Task).
- **Separation from Usecases**: The repositories do not have knowledge of the usecases that invoke them. They are strictly focused on data operations and are abstracted away from business logic. This allows for clear separation of concerns.
- **Interface Driven**: Repositories are injected into the usecases as interfaces, enabling easy swapping of implementations (e.g., swapping in-memory storage for a relational database) without modifying the usecase layer.
- **Domain-Centric**: The repositories are designed around the domain models, meaning they deal directly with the entities relevant to the business, ensuring data integrity and consistency.
- **In-Memory Implementation**: For demonstration purposes, in-memory repositories have been implemented. However, thanks to the Clean Architecture principles, it is easy to replace the in-memory repositories with real databases such as PostgreSQL, MongoDB, or any other database solution. This flexibility allows for easy migration to production-ready systems without disrupting the overall architecture.

This layer acts as a bridge, providing the necessary functionality for the usecases to interact with data while keeping the persistence logic decoupled from business logic.

### **Domain Layer**

The **Domain Layer** is at the heart of the application, representing the core business logic and entities. It encapsulates the essential domain models and the logic directly related to them.

- **Model Objects**: This layer contains the domain models, which represent the core entities of the system, such as Todo Lists and Todo Tasks. These models define the structure and behavior of the data and are used throughout the application.
- **Domain Logic**: While the focus of the domain layer is on the entities themselves, it may also include domain-related logic, such as methods that operate directly on the models to enforce rules or perform calculations.
- **No External Dependencies**: The domain layer is kept free of external dependencies. It does not rely on any frameworks, databases, or infrastructure concerns. This ensures that the business logic is isolated and can be independently tested or changed without impacting other parts of the system.
- **Pure Business Focus**: The domain layer is purely business-driven. It’s not concerned with how data is stored or retrieved (handled by repositories) or how requests are processed (handled by usecases and handlers). It simply defines the core entities and their behaviors.

This design ensures that the domain logic is highly cohesive, maintainable, and adaptable to different use cases, as it remains decoupled from the external systems that may change over time.

### Endpoints Overview:

### Backend REST API Endpoints

#### **`GET /health`**  
- **Purpose**: Returns the health status of the application.  
- **Response**:  
  - If the application is healthy, the response will be `"healthy"`.  
- **Example cURL**:
    ```bash
    curl --location 'http://127.0.0.1:8080/health'
    ```  
  - **Response**:
    ```json
    "healthy"
    ```

#### **`POST /users/login`**  
- **Purpose**: Authenticates the user and returns a JWT token for subsequent requests.  
- **Request**:
    - **Headers**:  
      - `Authorization`: Basic Auth with the format `username:password` encoded in Base64.
    - **Body**: N/A (credentials sent via Basic Auth header).
- **Response**:  
  - On successful authentication, a JWT token is returned in the `value` field.  
  - Example successful response:
    ```json
    {
      "message": "200 - OK",
      "value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkVtcmUiLCJyb2xlIjoidXNlciIsInN1YiI6IjAifQ.Qg4BKss8SENaq4AgWA3xu7YnDKiX7oEqbmyQTA0V-PI"
    }
    ```
  - If authentication fails, an appropriate error message will be returned.

- **Example cURL**:
    ```bash
    curl --location 'http://127.0.0.1:8080/users/login' \
    --header 'Authorization: Basic RW1yZToxMjM0'
    ```

#### **`GET /todo-lists`**  
- **Purpose**: Fetches a specific todo list based on the provided `id`.  
- **Query Parameters**:  
    - `id` (required): The unique ID of the todo list to be fetched.
- **Request**:
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location 'http://127.0.0.1:8080/todo-lists?id=1' \
    --header 'Authorization: Bearer <JWT_TOKEN>'
    ```
- **Response**:  
    - **200 OK**: Returns the requested todo list data.  
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": {
            "id": "1",
            "user_id": "0",
            "name": "Personal Goals",
            "created_at": "2025-05-07T23:16:41.368987+03:00",
            "modified_at": "2025-05-07T23:16:41.368987+03:00",
            "deleted_at": null,
            "completion_percent": 100,
            "completed_tasks": 2,
            "total_tasks": 2
        }
      }
      ```
    - **404 Not Found**: The todo list with the given `id` does not exist.  
      - Example response:
      ```json
      {
        "message": "Todo list not found"
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`GET /todo-lists`**  
- **Purpose**: Fetches a list of all todo lists associated with the authenticated user.  
- **Request**:
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location 'http://127.0.0.1:8080/todo-lists' \
    --header 'Authorization: Bearer <JWT_TOKEN>'
    ```
- **Response**:  
    - **200 OK**: Returns a list of todo lists associated with the authenticated user.  
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": [
          {
            "id": "0",
            "user_id": "0",
            "name": "Work Tasks",
            "created_at": "2025-05-07T23:16:41.368986+03:00",
            "modified_at": "2025-05-07T23:16:41.368987+03:00",
            "deleted_at": null,
            "completion_percent": 67,
            "completed_tasks": 2,
            "total_tasks": 3
          },
          {
            "id": "1",
            "user_id": "0",
            "name": "Personal Goals",
            "created_at": "2025-05-07T23:16:41.368987+03:00",
            "modified_at": "2025-05-07T23:16:41.368987+03:00",
            "deleted_at": null,
            "completion_percent": 100,
            "completed_tasks": 2,
            "total_tasks": 2
          },
          {
            "id": "2",
            "user_id": "1",
            "name": "Shopping List",
            "created_at": "2025-05-07T23:16:41.368987+03:00",
            "modified_at": "2025-05-07T23:16:41.368987+03:00",
            "deleted_at": null,
            "completion_percent": 25,
            "completed_tasks": 1,
            "total_tasks": 4
          }
        ]
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`GET /todo-lists?user_id={user_id}`**  
- **Purpose**: Fetches a list of todo lists associated with a specific user.  
- **Request**:
    - **Query Parameters**:  
      - `user_id` (required): The ID of the user whose todo lists are to be fetched.
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location 'http://127.0.0.1:8080/todo-lists?user_id=1' \
    --header 'Authorization: Bearer <JWT_TOKEN>'
    ```
- **Response**:  
    - **200 OK**: Returns a list of todo lists associated with the specified user.  
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": [
          {
            "id": "2",
            "user_id": "1",
            "name": "Shopping List",
            "created_at": "2025-05-07T23:16:41.368987+03:00",
            "modified_at": "2025-05-07T23:16:41.368987+03:00",
            "deleted_at": null,
            "completion_percent": 25,
            "completed_tasks": 1,
            "total_tasks": 4
          }
        ]
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`POST /todo-lists`**  
- **Purpose**: Creates a new todo list for the specified user.  
- **Request**:
    - **Form Data**:
      - `user_id` (required): The ID of the user who owns the todo list.
      - `name` (required): The name of the new todo list.
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location 'http://127.0.0.1:8080/todo-lists' \
    --header 'Authorization: Bearer <JWT_TOKEN>' \
    --form 'user_id="1"' \
    --form 'name="naber"'
    ```
- **Response**:  
    - **200 OK**: Returns the newly created todo list details.  
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": {
          "id": "3",
          "user_id": "1",
          "name": "naber",
          "created_at": "2025-05-07T23:23:57.889141+03:00",
          "modified_at": "2025-05-07T23:23:57.889141+03:00",
          "deleted_at": null,
          "completion_percent": 0,
          "completed_tasks": 0,
          "total_tasks": 0
        }
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`PATCH /todo-lists`**  
- **Purpose**: Updates the details of an existing todo list.  
- **Request**:
    - **Form Data**:
      - `id` (required): The ID of the todo list to be updated.
      - `name` (optional): The new name of the todo list.
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location --request PATCH 'http://127.0.0.1:8080/todo-lists' \
    --header 'Authorization: Bearer <JWT_TOKEN>' \
    --form 'id="2"' \
    --form 'name="new tdoooo"'
    ```
- **Response**:  
    - **200 OK**: Returns the updated todo list details.  
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": {
          "id": "2",
          "user_id": "1",
          "name": "new tdoooo",
          "created_at": "2025-05-07T23:16:41.368987+03:00",
          "modified_at": "2025-05-07T23:24:56.466992+03:00",
          "deleted_at": null,
          "completion_percent": 25,
          "completed_tasks": 1,
          "total_tasks": 4
        }
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`DELETE /todo-lists`**  
- **Purpose**: Deletes a specific todo list.  
- **Request**:
    - **Form Data**:
      - `id` (required): The ID of the todo list to be deleted.
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location --request DELETE 'http://127.0.0.1:8080/todo-lists' \
    --header 'Authorization: Bearer <JWT_TOKEN>' \
    --form 'id="2"'
    ```
- **Response**:  
    - **200 OK**: Returns the deleted todo list details with `deleted_at` field populated.
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": {
          "id": "2",
          "user_id": "1",
          "name": "new tdoooo",
          "created_at": "2025-05-07T23:16:41.368987+03:00",
          "modified_at": "2025-05-07T23:25:31.825894+03:00",
          "deleted_at": "2025-05-07T23:25:31.825894+03:00",
          "completion_percent": 25,
          "completed_tasks": 1,
          "total_tasks": 4
        }
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`GET /todo-tasks`**  
- **Purpose**: Retrieves all tasks within a specific todo list.  
- **Request**:
    - **Query Parameters**:
      - `list_id` (required): The ID of the todo list for which tasks will be retrieved.
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location 'http://127.0.0.1:8080/todo-tasks?list_id=0' \
    --header 'Authorization: Bearer <JWT_TOKEN>'
    ```
- **Response**:  
    - **200 OK**: Returns a list of tasks in the specified todo list.
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": [
          {
            "id": "0",
            "todo_list_id": "0",
            "created_at": "2025-05-07T23:16:41.368993+03:00",
            "modified_at": "2025-05-07T23:16:41.368994+03:00",
            "deleted_at": null,
            "content": "Finish the UI layout",
            "is_completed": true
          },
          {
            "id": "1",
            "todo_list_id": "0",
            "created_at": "2025-05-07T23:16:41.368994+03:00",
            "modified_at": "2025-05-07T23:16:41.368994+03:00",
            "deleted_at": null,
            "content": "Write the backend logic",
            "is_completed": false
          },
          {
            "id": "2",
            "todo_list_id": "0",
            "created_at": "2025-05-07T23:16:41.368994+03:00",
            "modified_at": "2025-05-07T23:16:41.368994+03:00",
            "deleted_at": null,
            "content": "Connect frontend to backend",
            "is_completed": true
          }
        ]
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`POST /todo-tasks`**  
- **Purpose**: Creates a new task within a specific todo list.  
- **Request**:
    - **Form Data**:
      - `list_id` (required): The ID of the todo list to which the task will be added.
      - `content` (required): The content or description of the task.
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location 'http://127.0.0.1:8080/todo-tasks' \
    --header 'Authorization: Bearer <JWT_TOKEN>' \
    --form 'list_id="0"' \
    --form 'content="buy a new garbage bin"'
    ```
- **Response**:  
    - **200 OK**: Returns the newly created task with details.
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": {
          "id": "9",
          "todo_list_id": "0",
          "created_at": "2025-05-07T23:26:43.710925+03:00",
          "modified_at": "2025-05-07T23:26:43.710925+03:00",
          "deleted_at": null,
          "content": "buy a new garbage bin",
          "is_completed": false
        }
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`PATCH /todo-tasks`**  
- **Purpose**: Updates a specific task, such as toggling its completion status or modifying task details.  
- **Request**:
    - **Form Data**:
      - `id` (required): The ID of the task to update.
      - `action` (required): The action to perform on the task (e.g., `toggle` to toggle completion status).
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location --request PATCH 'http://127.0.0.1:8080/todo-tasks' \
    --header 'Authorization: Bearer <JWT_TOKEN>' \
    --form 'id="7"' \
    --form 'action="toggle"'
    ```
- **Response**:  
    - **200 OK**: Returns the updated task with details.
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": {
          "id": "7",
          "todo_list_id": "2",
          "created_at": "2025-05-07T23:28:07.778645+03:00",
          "modified_at": "2025-05-07T23:28:09.535368+03:00",
          "deleted_at": null,
          "content": "Eggs",
          "is_completed": false
        }
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`PATCH /todo-tasks`**  
- **Purpose**: Updates a specific task, such as toggling its completion status, modifying task details, or editing task content.  
- **Request**:
    - **Form Data**:
      - `id` (required): The ID of the task to update.
      - `action` (required): The action to perform on the task (e.g., `toggle` to toggle completion status, `edit` to modify the task content).
      - `content` (optional): The new content for the task, required when the action is `edit`.
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location --request PATCH 'http://127.0.0.1:8080/todo-tasks' \
    --header 'Authorization: Bearer <JWT_TOKEN>' \
    --form 'id="6"' \
    --form 'action="edit"' \
    --form 'content="patched task content"'
    ```
- **Response**:  
    - **200 OK**: Returns the updated task with details.
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": {
          "id": "6",
          "todo_list_id": "2",
          "created_at": "2025-05-07T23:28:07.778644+03:00",
          "modified_at": "2025-05-07T23:29:12.971604+03:00",
          "deleted_at": null,
          "content": "patched task content",
          "is_completed": false
        }
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

#### **`DELETE /todo-tasks`**  
- **Purpose**: Deletes a specific task from the todo list.  
- **Request**:
    - **Form Data**:
      - `id` (required): The ID of the task to delete.
    - **Headers**:
      - `Authorization`: Bearer token (JWT) used for authentication.
    - **Example cURL**:
    ```bash
    curl --location --request DELETE 'http://127.0.0.1:8080/todo-tasks' \
    --header 'Authorization: Bearer <JWT_TOKEN>' \
    --form 'id="7"'
    ```
- **Response**:  
    - **200 OK**: Returns the deleted task with a `deleted_at` timestamp.
      - Example successful response:
      ```json
      {
        "message": "200 - OK",
        "value": {
          "id": "7",
          "todo_list_id": "2",
          "created_at": "2025-05-07T23:28:07.778645+03:00",
          "modified_at": "2025-05-07T23:29:42.776017+03:00",
          "deleted_at": "2025-05-07T23:29:42.776016+03:00",
          "content": "Eggs",
          "is_completed": false
        }
      }
      ```
    - **401 Unauthorized**: Invalid or missing JWT token.  
      - Example response:
      ```json
      {
        "message": "Invalid JWT token"
      }
      ```

---

## 🚀 Deployment

The application is deployed on **[Render.com](https://render.com)**.

- Environment variables (e.g., secrets, configs) are managed using a `.env` file.
- Render's environment settings are used to set these variables securely during deployment.
- No external database is used; all data is stored in memory for demonstration purposes.

---

## 📄 License

This project is for demonstration purposes as part of a case study and is not currently under an open-source license.