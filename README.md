# URL Shortener

A simple URL shortener application built with Go and Gin framework.

## Features

- Shorten long URLs 
- Resolve shortened URLs to their original form
- Rate limit on Shorten requests 


## Installation

TO DO Docker Compose File

1. Clone the repository:
   ```bash
   git clone https://github.com/Mu-Wahba/url-shortener.git
   cd url-shortener
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the root directory from .env.sample

## Usage

1. Start the application:
   ```bash
   go run main.go
   ```

2. Use the following endpoints:

   - **Shorten URL**: 
     - **POST** `/api/v2/url`
     - Request body: JSON with the long URL
     ```json
     {
       "url": "https://example.com" //REQUIRED
       "expiry" : 3, //In minutes , default to 1 year
       "custom_shorten_url": "97709d" //6 Chars , default to random chars
     }
     ```

   - **Resolve URL**: 
     - **GET** `/:url`
     - Replace `:url` with the shortened URL code. //97709d


## Rate Limiting

To prevent abuse of the URL shortening service, rate limiting is implemented. Below is an example of how to handle rate limits when making requests.

### Example of Rate Limit Response

If a user exceeds the allowed number of requests API_QUOTA within API_LIMIT_PERIOD, the server will respond with a rate limit error.

**Request:**
```bash
curl -X POST http://localhost:8080/api/v2/url \
-H "Content-Type: application/json" \
-d '{
  "url": "https://example.com"
}'
```

**Response (Rate Limit Exceeded):**
```json

  {
    "msg": "Rate limit excceded ",
    "remaining_minutes": 2
}

```
