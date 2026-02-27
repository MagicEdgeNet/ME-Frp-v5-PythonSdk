import asyncio
from typing import Any, Dict, List, Optional

import aiohttp

from .base import BaseClient
from .exceptions import APIError, AuthError, NetworkError
from .models import *


class AsyncMEFrpClient(BaseClient):
    def __init__(
        self,
        token: str = "",
        base_url: Optional[str] = None,
        user_agent: Optional[str] = None,
        timeout: float = 10.0,
        bypass_system_proxy: bool = False,
        session: Optional[aiohttp.ClientSession] = None,
    ):
        super().__init__(token, base_url, user_agent, timeout, bypass_system_proxy)
        self._session = session
        self._own_session = False

    async def _get_session(self) -> aiohttp.ClientSession:
        if self._session is None or self._session.closed:
            self._session = aiohttp.ClientSession(trust_env=not self.bypass_system_proxy)
            self._own_session = True
        return self._session

    async def close(self):
        if self._own_session and self._session and not self._session.closed:
            await self._session.close()

    async def __aenter__(self):
        return self

    async def __aexit__(self, exc_type, exc_val, exc_tb):
        await self.close()

    async def _request(
        self, method: str, path: str, json_data: Any = None, params: Optional[Dict[str, Any]] = None
    ) -> Any:
        url = f"{self.base_url}{path}"
        session = await self._get_session()

        try:
            async with session.request(
                method,
                url,
                json=json_data,
                params=params,
                headers=self._get_headers(),
                timeout=aiohttp.ClientTimeout(total=self.timeout),
            ) as resp:
                if resp.status == 401:
                    raise AuthError("Unauthorized: invalid token", 401)

                try:
                    data = await resp.json()
                except Exception as e:
                    text = await resp.text()
                    raise APIError(f"Failed to decode response: {text}", resp.status) from e

                if data.get("code") != 200:
                    raise APIError(data.get("message", "Unknown error"), data.get("code"))

                return data.get("data")
        except aiohttp.ClientError as e:
            raise NetworkError(f"Network request failed: {str(e)}") from e
        except asyncio.TimeoutError as e:
            raise NetworkError("Request timed out") from e

    # --- Public API ---
    async def get_register_email_code(self, email: str, captcha_token: str):
        return await self._request(
            "POST", "/public/register/emailCode", {"email": email, "captchaToken": captcha_token}
        )

    async def register(self, username, email, email_code, password):
        return await self._request(
            "POST",
            "/public/register",
            {"username": username, "email": email, "emailCode": email_code, "password": password},
        )

    async def login(self, username, password, captcha_token):
        data = await self._request(
            "POST",
            "/public/login",
            {"username": username, "password": password, "captchaToken": captcha_token},
        )
        if data and "token" in data:
            self.set_token(data["token"])
        return data

    async def forgot_password(self, username: str, email: str):
        return await self._request("POST", "/public/iforgot", {"username": username, "email": email})

    async def get_statistics(self) -> Statistics:
        data = await self._request("GET", "/public/statistics")
        return Statistics(**data)

    async def get_popup_notice(self) -> str:
        return await self._request("GET", "/auth/popupNotice")

    async def get_notice(self) -> List[str]:
        return await self._request("GET", "/auth/notice")

    async def get_system_status(self) -> SystemStatus:
        data = await self._request("GET", "/auth/system/status")
        return SystemStatus(**data)

    # --- User API ---
    async def get_user_info(self) -> UserInfo:
        data = await self._request("GET", "/auth/user/info")
        return UserInfo(**data)

    async def sign(self, captcha_token: str):
        return await self._request("POST", "/auth/user/sign", {"captchaToken": captcha_token})

    async def get_user_groups(self) -> List[UserGroup]:
        data = await self._request("GET", "/auth/user/groups")
        return [UserGroup(**g) for g in data]

    async def get_frp_token(self) -> str:
        data = await self._request("GET", "/auth/user/frpToken")
        return data.get("frpToken") if isinstance(data, dict) else data

    async def reset_token(self, captcha_token: str):
        return await self._request("POST", "/auth/user/tokenReset", {"captchaToken": captcha_token})

    async def reset_password(self, old_password: str, new_password: str):
        return await self._request(
            "POST",
            "/auth/user/passwordReset",
            {"oldPassword": old_password, "newPassword": new_password},
        )

    # --- Proxy API ---
    async def get_proxy_list(self) -> ProxyListResponse:
        data = await self._request("GET", "/auth/proxy/list")
        proxies = [Proxy(**p) for p in data["proxies"]]
        nodes = [NodeConnection(**n) for n in data["nodes"]]
        return ProxyListResponse(proxies=proxies, nodes=nodes)

    async def create_proxy(self, req_data: Dict):
        return await self._request("POST", "/auth/proxy/create", req_data)

    async def delete_proxy(self, proxy_id: int):
        return await self._request("POST", "/auth/proxy/delete", {"proxyId": proxy_id})

    async def update_proxy(self, req_data: Dict):
        return await self._request("POST", "/auth/proxy/update", req_data)

    async def kick_proxy(self, proxy_id: int):
        return await self._request("POST", "/auth/proxy/kick", {"proxyId": proxy_id})

    async def toggle_proxy(self, proxy_id: int):
        return await self._request("POST", "/auth/proxy/toggle", {"proxyId": proxy_id})

    async def get_proxy_config(self, proxy_id: int, format: str = "ini") -> ProxyConfigResponse:
        data = await self._request(
            "POST", "/auth/proxy/config", {"proxyId": proxy_id, "format": format}
        )
        return ProxyConfigResponse(**data)

    # --- Node API ---
    async def get_node_list(self) -> List[Node]:
        data = await self._request("GET", "/auth/node/list")
        return [Node(**n) for n in data]

    async def get_node_status(self, node_id: int) -> NodeStatus:
        data = await self._request("GET", "/auth/node/status", params={"nodeId": node_id})
        return NodeStatus(**data)

    async def get_node_name_list(self) -> List[NodeNameListItem]:
        data = await self._request("GET", "/auth/node/nameList")
        return [NodeNameListItem(**n) for n in data]

    async def get_node_token(self, node_id: int) -> str:
        data = await self._request("POST", "/auth/node/secret", {"nodeId": node_id})
        return data.get("token") if isinstance(data, dict) else data

    async def get_free_port(self, node_id: int, proxy_type: str):
        return await self._request(
            "GET", "/auth/node/freePort", params={"nodeId": node_id, "proxyType": proxy_type}
        )

    # --- Ads API ---
    async def get_user_ads(self) -> List[Ads]:
        data = await self._request("GET", "/auth/ads/manage")
        return [Ads(**a) for a in data]

    async def get_user_ad_credits(self) -> List[AdCredit]:
        data = await self._request("GET", "/auth/ads/credits")
        return [AdCredit(**c) for c in data]

    # --- Order API ---
    async def get_orders(self, page: int = 1, page_size: int = 10) -> List[Order]:
        data = await self._request(
            "GET", "/auth/orders", params={"page": page, "pageSize": page_size}
        )
        return [Order(**o) for o in data]

    # --- Node Donate API ---
    async def get_user_node_donates(self) -> List[NodeDonate]:
        data = await self._request("GET", "/auth/node/donate/list")
        return [NodeDonate(**d) for d in data]

    async def apply_node_donate(self, req_data: Dict):
        return await self._request("POST", "/auth/node/donate", req_data)

    # --- Store API ---
    async def get_store_products(self) -> List[StoreItem]:
        data = await self._request("GET", "/public/store/products")
        return [StoreItem(**p) for p in data]

    # --- Easy Startup API ---
    async def get_easy_startup_config(self, proxy_id: int) -> EasyStartProxy:
        data = await self._request("POST", "/auth/easyStartup", {"proxyId": proxy_id})
        return EasyStartProxy(**data)

    # --- Operation Log API ---
    async def get_operation_logs(self, page: int = 1, page_size: int = 10) -> OperationLogList:
        data = await self._request(
            "GET", "/auth/operationLog/list", params={"page": page, "pageSize": page_size}
        )
        return OperationLogList(
            data=[OperationLog(**log) for log in data["data"]],
            total=data["total"],
            page=data["page"],
            pageSize=data["pageSize"],
            totalPages=data["totalPages"],
        )

    async def get_operation_log_stats(self) -> OperationLogStats:
        data = await self._request("GET", "/auth/operationLog/stats")
        return OperationLogStats(**data)

    async def get_create_proxy_data(self) -> CreateProxyDataResponse:
        data = await self._request("GET", "/auth/createProxyData")
        return CreateProxyDataResponse(**data)

    # --- CDK API ---
    async def redeem_cdk(self, code: str):
        return await self._request("POST", "/auth/cdk/redeem", {"code": code})
