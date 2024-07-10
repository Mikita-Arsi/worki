import requests


data = {
    "firstname": "loh",
    "username": "govno"
}

resp = requests.delete('http://localhost:1323/users/1')
print(requests.get('http://localhost:1323/users/').json())