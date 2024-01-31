
# Go REST API with Gorilla Mux

This is a simple Go-based RESTful API project using the Gorilla Mux router. The API allows users to manage events with basic CRUD (Create, Read, Update, Delete) operations. It includes endpoints for creating new events, retrieving individual or all events, updating event details, and deleting events.

## Features
- **Create Event**: Add a new event by providing a title and description through a POST request.
- **Get One Event**: Retrieve details of a specific event using its ID through a GET request.
- **Get All Events**: Fetch a list of all events through a GET request.
- **Update Event**: Modify the details of a specific event using a PATCH request.
- **Delete Event**: Remove a specific event using a DELETE request.

## Usage
1. **Homepage**: The root endpoint ("/") welcomes users to the API.
2. **Create Event**: Use a POST request to "/event" with the event details in the request body.
3. **Get One Event**: Retrieve a specific event by making a GET request to "/events/{id}".
4. **Get All Events**: Fetch all events by making a GET request to "/events".
5. **Update Event**: Update event details by sending a PATCH request to "/events/{id}".
6. **Delete Event**: Remove a specific event with a DELETE request to "/events/{id}".

## Getting Started
1. Clone the repository.
2. Install dependencies, including Gorilla Mux.
3. Run the Go application.
4. Access the API using the provided endpoints.

Feel free to contribute to this project or use it as a starting point for your own Go-based REST API development.

Happy coding!
