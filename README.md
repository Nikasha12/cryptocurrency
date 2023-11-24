"# cryptocurrency" 
"# cryptocurrency" 

**Clone the repository:**
git clone https://github.com/Nikasha12/cryptocurrency.git
cd cryptocurrency

Run the application:
go run main.go 
•	The server will start on http://localhost:9090

Testing
•	To run tests, use:
go test

Testing Strategy
•	Success Test: Checks if the API returns the correct response for a successful request.
•	Missing Crypto Parameter Test: Verifies the handling of requests without the required crypto parameter.
•	API Request Failure Test: Ensures proper handling of failures when the external API call encounters an issue.
