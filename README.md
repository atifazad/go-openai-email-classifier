# Email Classifier

This project is an Email Classifier built using Go, leveraging the go-gpt3 classification API for classifying emails. It stores classification results in an SQLite database and provides a JSON API for frontend interaction.

## Setup Instructions

1. **Clone the repository:**

   ```
   git clone <repository-url>
   cd email-classifier
   ```

2. **Install dependencies:**

   ```
   go mod tidy
   ```

3. **Set up your OpenAI API key:**
   Make sure to set your `OpenAI_API_Key` in your environment variables.

4. **Run the application:**
   ```
   go run cmd/main.go
   ```

## API Endpoints

### Classify an Email

- **URL**: `/classify`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "subject": "Test Email",
    "content": "This is a test email content."
  }
  ```
- **Response**:
  ```json
  {
    "classification": "not spam"
  }
  ```

### Retrieve All Classifications

- **URL**: `/classifications`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "id": 1,
      "subject": "Test Email",
      "content": "This is a test email content.",
      "classification": "not spam"
    }
  ]
  ```

### Retrieve a Classification by ID

- **URL**: `/classification?id=1`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "id": 1,
    "subject": "Test Email",
    "content": "This is a test email content.",
    "classification": "not spam"
  }
  ```

### Delete a Classification by ID

- **URL**: `/delete-classification?id=1`
- **Method**: `DELETE`
- **Response**: `204 No Content`
