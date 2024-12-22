import requests

# URL of the local endpoint
url = 'http://localhost:3131/verify'

# Send 60 requests
for i in range(61):
    try:
        response = requests.get(url)
        print(f"Request {i+1}: Status Code: {response.status_code}")
    except requests.exceptions.RequestException as e:
        print(f"Request {i+1} failed: {e}")