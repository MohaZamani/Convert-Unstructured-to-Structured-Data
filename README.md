# Convert Unstructured Data of Laptops to Structured Data


This project is a Golang-based solution designed to leverage OpenAI for converting unstructured data about laptops into structured data. Below are the instructions to build and run the project using Docker.

## Prerequisites
- **Docker**: Ensure Docker is installed on your machine.
- **Terminal**: A command line interface to execute Docker commands.

## Steps to Build and Run the Project in Docker

1. **Clone or Download the Repository**:  
   Download the project files and navigate to the root directory of the project.  
   ```bash
   git clone https://github.com/MohaZamani/Convert-Unstructured-to-Structured-Data
2. **Build the Docker Image**:  
   In the project root directory, run the following command to build the Docker image:
   ```bash
   docker build -t convert-u-t-s-data .
3. **Run the Docker Container**:
    ```bash
    docker run -p 8080:8080 convert-u-t-s-data
4. **Access the Application**:

Once the container is running, you can connect to the application at https://localhost:8080.

## API Documentation

For detailed API documentation, including request/response examples and additional information, please visit the following link:

[Postman API Documentation](https://documenter.getpostman.com/view/14995830/2sAXxS7BNn)