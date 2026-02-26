from typing import Any, Dict, List, Optional

import requests

from .base import BaseClient
from .exceptions import APIError, AuthError, NetworkError
from .models import *


class MEFrpClient(BaseClient):
    def _request(
        self, method: str, path: str, json_data: Any = None, params: Optional[Dict[str, Any]] = None
    ) -> Any:
        url = f"{self.base_url}{path}"
        try:
            with requests.Session() as session:
                session.trust_env = not self.bypass_system_proxy
                resp = session.request(
                    method,
                    url,
                    json=json_data,
                    params=params,
                    headers=self._get_headers(),
                    timeout=self.timeout,
                )
        except requests.exceptions.RequestException as e:
            raise NetworkError(f"Network request failed: {str(e)}") from e

        if resp.status_code == 401:
            raise AuthError("Unauthorized: invalid token", 401)

        try:
            data = resp.json()
        except ValueError as e:
            raise APIError(f"Failed to decode response: {resp.text}", resp.status_code) from e

        if data.get("code") != 200:
            raise APIError(data.get("message", "Unknown error"), data.get("code"))

        return data.get("data")

    # --- Public API ---
    def get_register_email_code(self, email: str, captcha_token: str):
        return self._request(
            "POST", "/public/register/emailCode", {"email": email, "captchaToken": captcha_token}
        )

    def register(self, username, email, email_code, password):
        return self._request(
            "POST",
            "/public/register",
            {"username": username, "email": email, "emailCode": email_code, "password": password},
        )

    def login(self, username, password, captcha_token):
        data = self._request(
            "POST",
            "/public/login",
            {"username": username, "password": password, "captchaToken": captcha_token},
        )
        if data and "token" in data:
            self.set_token(data["token"])
        return data

    def get_statistics(self) -> Statistics:
        data = self._request("GET", "/public/statistics")
        return Statistics(**data)

    # --- User API ---
    def get_user_info(self) -> UserInfo:
        data = self._request("GET", "/auth/user/info")
        return UserInfo(**data)

    def sign(self, captcha_token: str):
        return self._request("POST", "/auth/user/sign", {"captchaToken": captcha_token})

    def get_user_groups(self) -> List[UserGroup]:
        data = self._request("GET", "/auth/user/groups")
        return [UserGroup(**g) for g in data]

    # --- Proxy API ---
    def get_proxy_list(self) -> ProxyListResponse:
        data = self._request("GET", "/auth/proxy/list")
        proxies = [Proxy(**p) for p in data["proxies"]]
        nodes = [NodeConnection(**n) for n in data["nodes"]]
        return ProxyListResponse(proxies=proxies, nodes=nodes)

    def create_proxy(self, req_data: Dict):
        return self._request("POST", "/auth/proxy/create", req_data)

    def delete_proxy(self, proxy_id: int):
        return self._request("POST", "/auth/proxy/delete", {"proxyId": proxy_id})

    def update_proxy(self, req_data: Dict):
        return self._request("POST", "/auth/proxy/update", req_data)

    def kick_proxy(self, proxy_id: int):
        return self._request("POST", "/auth/proxy/kick", {"proxyId": proxy_id})

    def toggle_proxy(self, proxy_id: int):
        return self._request("POST", "/auth/proxy/toggle", {"proxyId": proxy_id})

    # --- Node API ---
    def get_node_list(self) -> List[Node]:
        data = self._request("GET", "/auth/node/list")
        return [Node(**n) for n in data]

    def get_node_status(self, node_id: int) -> NodeStatus:
        data = self._request("GET", "/auth/node/status", params={"nodeId": node_id})
        return NodeStatus(**data)

    # --- Ads API ---
    def get_user_ads(self) -> List[Ads]:
        data = self._request("GET", "/auth/ads/manage")
        return [Ads(**a) for a in data]

    def get_user_ad_credits(self) -> List[AdCredit]:
        data = self._request("GET", "/auth/ads/credits")
        return [AdCredit(**c) for c in data]

    # --- Order API ---
    def get_orders(self, page: int = 1, page_size: int = 10) -> List[Order]:
        data = self._request("GET", "/auth/orders", params={"page": page, "pageSize": page_size})
        return [Order(**o) for o in data]

    # --- Node Donate API ---
    def get_user_node_donates(self) -> List[NodeDonate]:
        data = self._request("GET", "/auth/node/donate/list")
        return [NodeDonate(**d) for d in data]

    def apply_node_donate(self, req_data: Dict):
        return self._request("POST", "/auth/node/donate", req_data)

    # --- Store API ---
    def get_store_products(self) -> List[StoreItem]:
        data = self._request("GET", "/public/store/products")
        return [StoreItem(**p) for p in data]

    # --- Easy Startup API ---
    def get_easy_startup_config(self, proxy_id: int) -> EasyStartProxy:
        data = self._request("POST", "/auth/easyStartup", {"proxyId": proxy_id})
        return EasyStartProxy(**data)

    # --- Operation Log API ---
    def get_operation_logs(self, page: int = 1, page_size: int = 10) -> OperationLogList:
        data = self._request(
            "GET", "/auth/operationLog/list", params={"page": page, "pageSize": page_size}
        )
        return OperationLogList(
            data=[OperationLog(**log) for log in data["data"]],
            total=data["total"],
            page=data["page"],
            pageSize=data["pageSize"],
            totalPages=data["totalPages"],
        )

    # --- CDK API ---
    def redeem_cdk(self, code: str):
        return self._request("POST", "/auth/cdk/redeem", {"code": code})
