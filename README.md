# Receipt Processor

## Overview
The application processes retail receipts and calculates reward points based on specific rules. It exposes a RESTful API with two main endpoints:

- `POST /receipts/process`: Accepts a JSON receipt, processes it, and returns a unique receipt ID.
- `GET /receipts/{id}/points`: Returns the number of points awarded for the receipt with the given ID.

The application is built using Go and can be run locally or within a Docker container for ease of deployment.

## Prerequisites
- Go 
- Docker

## Getting Started

### Running Locally

1. Clone the Repository
```bash
git clone 
cd receipt-processor
```

2. Install Dependencies
```bash
go mod tidy
```

3. Run the Application
```bash
go run main.go processor.go models.go
```

The server will start and listen on port 9000.

### Running with Docker

1. Build the Docker Image
```bash
docker build -t receipt-processor .
```

2. Run the Docker Container
```bash
docker run -p 9000 receipt-processor
```

The application will be accessible at http://localhost:9000.

## API Documentation

### Process Receipt
- Endpoint: `/receipts/process`
- Method: `POST`
- Content-Type: `application/json`
- Description: Accepts a JSON receipt and returns a unique receipt ID.

#### Request Body
```json
{
  "retailer": "Retailer Name",
  "purchaseDate": "YYYY-MM-DD",
  "purchaseTime": "HH:MM",
  "total": "Total Amount",
  "items": [
    {
      "shortDescription": "Item Description",
      "price": "Item Price"
    },
    {
      "shortDescription": "Item Description",
      "price": "Item Price"
    }
  ]
}
```

#### Response
Status Code: 200 OK
```json
{
  "id": "generated-receipt-id"
}
```

### Get Points
- Endpoint: `/receipts/{id}/points`
- Method: `GET`
- Description: Returns the number of points awarded for the receipt with the given ID.

#### URL Parameters
- `{id}`: The receipt ID obtained from processing the receipt

#### Response
Status Code: 200 OK
```json
{
  "points": totalPoints
}
```

## Testing the Application

#### Example: Process a Receipt
Request:
```bash
curl -X POST -H "Content-Type: application/json" -d @examples/simple-receipt.json http://localhost:9000/receipts/process
```

Response:
```json
{
  "id": "your-generated-receipt-id"
}
```

#### Example: Get Points for a Receipt
Request:
```bash
curl http://localhost:9000/receipts/your-generated-receipt-id/points
```

Response:
```json
{
  "points": calculatedPoints
}
```

### Edge Case Testing
A comprehensive set of test cases has been considered to ensure robustness:

- Receipts with missing or empty fields
- Invalid date or time formats
- Negative or zero values for prices and totals
- Receipts with special characters in retailer names
- High-precision prices with more than two decimal places

## Project Structure

- `main.go`: Contains the entry point and HTTP server setup
- `models.go`: Defines data models used in the application
- `processor.go`: Implements the business logic for points calculation
- `Dockerfile`: Instructions to build the Docker image
- `go.mod` and `go.sum`: Go module files for dependency management