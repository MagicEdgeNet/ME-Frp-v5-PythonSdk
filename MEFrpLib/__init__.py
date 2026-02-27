try:
    from .client_async import AsyncMEFrpClient
except ImportError:
    AsyncMEFrpClient = None  # type: ignore

try:
    from .client_sync import MEFrpClient
except ImportError:
    MEFrpClient = None  # type: ignore

from .exceptions import APIError, AuthError, MEFrpError, NetworkError
from .models import *

__version__ = "3.2.0"
__all__ = [
    "MEFrpClient",
    "AsyncMEFrpClient",
    "MEFrpError",
    "APIError",
    "AuthError",
    "NetworkError",
]
