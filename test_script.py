import requests

# URL of the local endpoint
# For testing it hits the helath check endpoint of the rate limiter itself
url = 'http://localhost:8080/limit/extra?qpm=param_value_1'

# Send requests
for i in range(77):
    try:
        response = requests.get(url)
        print(f"Request {i+1}: Status Code: {response.status_code}")
    except requests.exceptions.RequestException as e:
        print(f"Request {i+1} failed: {e}")