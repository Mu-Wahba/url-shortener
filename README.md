# URL Shortener

A simple URL shortener application built with Go and Gin framework.

## Features

- Shorten long URLs 
- Resolve shortened URLs to their original form
- Rate limit on Shorten requests 


## Installation


1. Clone the repository:
   ```bash
   git clone https://github.com/Mu-Wahba/url-shortener.git
   cd url-shortener
   ```

2. Copy the .env file:
   - Copy the `.env.sample` file to `.env`:
     ```bash
     cp .env.sample .env
     ```
   - Adjust the `.env` file to suit your needs.

3. Start The Application using Docker Compose:
   - Ensure you have a docker installed.
   - Start the application using Docker Compose:
     ```bash
     docker-compose up -d
     ```


## Usage
Use the following endpoints:

   - **Shorten URL**: 
     - **POST** `/api/v2/url`
     - Request body: JSON with the long URL
     ```json
     {
       "url": "https://example.com" ,
       "expiry" : 3 , 
       "custom_shorten_url": "97709d" 
     }
     ```
      - url: REQUIRED – The long URL that you want to shorten.
      - expiry: Optional – Time in minutes after which the URL will expire (default is 1 year).
      - custom_shorten_url: Optional – A custom 6-character string for the shortened URL. If omitted, a random string will be used.

   - **Resolve URL**: 
     - **GET** `/:url`
     - Replace `:url` with the shortened URL code.
     ```bash
     curl http://localhost:8080/97709d
     ```


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
