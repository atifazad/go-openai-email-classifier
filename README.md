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

## Usage

- The API can be accessed at `http://localhost:8080/api/classify`.
- Send a POST request with the email content in JSON format to classify the email.

## License

This project is licensed under the MIT License.
