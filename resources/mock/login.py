from pprint import pprint
import requests
import urllib

from oauth_challenge_codes import codes
from oauth_config import BASE_URL, CLIENT_ID, REDIRECT_URI

def authenticate() -> str:
    payload = {
        "email": "test@test.com",
        "password": "111"
    }
    resp = requests.post(f"{BASE_URL}/api/v1/auth/login", json=payload)
    result = resp.json()
    if resp.status_code != 200:
        raise Exception(resp.text)
    return result["data"]["login_token"]


def authorize(challenge, login_token) -> str:
    params = {
        "response_type": "code",
        "client_id": CLIENT_ID,
        "redirect_uri": REDIRECT_URI,
        "scope": "openid profile",
        "state": "xyz",
        "code_challenge_method": "S256",
        "code_challenge": challenge,
    }
    headers = {
        "Authorization": f"Bearer {login_token}"
    }

    url = f"{BASE_URL}/oauth/authorize?" + urllib.parse.urlencode(params)
    resp = requests.get(url, headers=headers, allow_redirects=False)

    if resp.status_code not in (302, 303):
        raise Exception(f"Unexpected status: {resp.status_code}, body: {resp.text}")

    redirect_url = resp.headers.get("Location")
    result = urllib.parse.urlparse(redirect_url)
    q = urllib.parse.parse_qs(result.query)
    if "code" not in q:
        raise Exception("\"code\" not found in redirect url")
    
    return q["code"][0]


def token(code: str, code_verifier: str) -> dict:
    data = {
        "code": code,
        "grant_type": "authorization_code",
        "client_id": CLIENT_ID,
        "redirect_uri": REDIRECT_URI,
        "code_verifier": code_verifier,
    }
    url = f"{BASE_URL}/oauth/token"
    resp = requests.post(url, data=data)

    if resp.status_code != 200:
        raise Exception(resp.text)

    return resp.json()


def main():
    # 1) authenticate to get login_token
    # POST: /api/auth/login with test@test.com:111
    # get login token in response
    login_token = authenticate()

    # 2) build PKCE codes
    verify, challenge = codes()

    # 3) /oauth/authorize request to get "code"
    code = authorize(challenge, login_token)

    # 4) send post to /oauth/token to get JWT access_token
    result = token(code, verify)

    pprint(result, indent=2)

if __name__ == "__main__":
    main()
