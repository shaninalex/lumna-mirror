from pprint import pprint
import requests
from oauth_config import BASE_URL, CLIENT_ID, REDIRECT_URI


def refresh_request(rt: str):
    data = {
        "grant_type": "refresh_token",
        "refresh_token": rt,
        "client_id": "angular-web",
    }

    resp = requests.post(f"{BASE_URL}/oauth/token", data=data)
    if resp.status_code != 200:
        raise Exception(f"unable to refresh token:\n{resp.text}")

    return resp.json()


def main():
    result = refresh_request("WYn4PbxWCJcFT5t9WZXR1pEysqVmL7gm2TAH_9aj_u6O4oyhUyapPUc2qGky6LIsmLnWNKu1Fa0fAoezM3LgeA")
    pprint(result, indent=2)

if __name__ == "__main__":
    main()