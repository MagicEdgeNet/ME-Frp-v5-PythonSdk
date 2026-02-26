class MEFrpError(Exception):
    """Base exception for MEFrp SDK"""

    def __init__(self, message: str, code: int = None):
        super().__init__(message)
        self.message = message
        self.code = code


class APIError(MEFrpError):
    """Raised when API returns a non-200 code"""

    pass


class AuthError(MEFrpError):
    """Raised when authentication fails"""

    pass


class NetworkError(MEFrpError):
    """Raised when network request fails"""

    pass
