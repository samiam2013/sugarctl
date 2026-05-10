import requests

BASE = "https://api-us.libreview.io"
HEADERS = {
    "accept-encoding": "gzip",
    "cache-control": "no-cache",
    "connection": "Keep-Alive",
    "product": "llu.android",
    "version": "4.16.0",
    "Content-Type": "application/json",
}

def login(email, password):
    r = requests.post(f"{BASE}/llu/auth/login",
                      json={"email": email, "password": password},
                      headers=HEADERS)
    r.raise_for_status()
    data = r.json()
    token = data["data"]["authTicket"]["token"]
    account_id = data["data"]["user"]["id"]  # may need hashing — check response
    return token, account_id

def get_connections(token, account_id):
    headers = {**HEADERS,
               "authorization": f"Bearer {token}",
               "account-id": account_id}
    r = requests.get(f"{BASE}/llu/connections", headers=headers)
    r.raise_for_status()
    return r.json()
