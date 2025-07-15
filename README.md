# Balancer Demo

This project demonstrates a simple load balancer and demo HTTP servers in Go.

## Structure
- `main.go`: Starts the load balancer on port 3000.
- `demo_servers/demo_servers.go`: Starts three demo servers on ports 3001, 3002, and 3003.
- `core/`, `pool/`, `health_checker/`, `types/`: Core logic for balancing and health checking.

## Prerequisites
- Go 1.18+

## How to Run

1. **Clone the repository** (if not already):
   ```bash
   git clone <repo-url>
   cd balancer
   ```

2. **Start demo servers** (in one terminal):
   ```bash
   go run demo_servers/demo_servers.go
   ```

3. **Start the load balancer** (in another terminal):
   ```bash
   go run main.go
   ```

4. **Test the load balancer**:
   - Open your browser or use curl:
     ```bash
     curl http://localhost:3000/api/v1
     ```
   - You should see responses from different servers (ports 3001, 3002, 3003) in round-robin fashion.

## Health Check
- Each demo server exposes `/health` endpoint for health checking.

## Stopping
- Press `Ctrl+C` in each terminal to stop the servers.

## Notes
- Make sure ports 3000, 3001, 3002, and 3003 are free before running.
# balancer
