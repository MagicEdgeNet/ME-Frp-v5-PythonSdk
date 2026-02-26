from typing import Dict, Optional

from .const import BASE_URL, DEFAULT_USER_AGENT


class BaseClient:
    def __init__(
        self,
        token: str = "",
        base_url: Optional[str] = BASE_URL,
        user_agent: Optional[str] = DEFAULT_USER_AGENT,
        timeout: float = 10.0,
        bypass_system_proxy: bool = False,
    ):
        self.token = token
        self.base_url = (base_url or BASE_URL).rstrip("/")
        self.user_agent = user_agent or DEFAULT_USER_AGENT
        self.timeout = timeout
        self.bypass_system_proxy = bypass_system_proxy

    def set_token(self, token: str):
        self.token = token

    def set_base_url(self, url: str):
        self.base_url = url.rstrip("/")

    def set_user_agent(self, user_agent: str):
        self.user_agent = user_agent

    def _get_headers(self) -> Dict[str, str]:
        headers = {"Content-Type": "application/json", "User-Agent": self.user_agent}
        if self.token:
            headers["Authorization"] = f"Bearer {self.token}"
        return headers
