# Pace-Keeper  
<img align="middle" width="40px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" />
<img align="middle" width="70px" src="https://avatars.githubusercontent.com/u/1529926?s=200&v=4" />

This project implements a rate limiter in Go (Golang) using Gin that supports two different algorithms:
- Token Bucket
- Sliding Window

Features
- Modified Token Bucket Algorithm: Allows bursts of requests followed by a occasional refill of tokens.
- Sliding Window Algorithm: Provides a more granular approach to rate limiting by considering a sliding time window.
- Redis Integration: Uses Redis for state management to ensure consistency across distributed instances.
- Configuration: Flexible settings for rate limits, algorithms, and Redis connection.


Prerequisites
- Go 1.23 or higher
- Redis (version 6.0 or higher recommended)

Local development
- Clone repo
- Make sure redis is running
- Run limiter inside Limiter  TARGET="http://localhost:3131" PARAM="qpm" go run ./main.go
- Run test_server inside Tester go run main.go
- Modiy loop limit inside test_script and do python test_script.py

Prod deployment (Steps that were followed by me)
- Created helm chart
- Used redis.ClusterClient instead of redis.Client in code
- Provided TARGET and PARAM as env variables 
- My target was running inside same EKS cluster

Additional Details
- Currently the algo switch is at 200 after 200 qps precise sliding window is used
- If we have an api which provides us limits for the query param we want to put a limit on we can poll it at regular intervals (inside Limiter/main.go)
