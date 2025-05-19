from starlette.middleware.base import BaseHTTPMiddleware
from starlette.requests import Request
from starlette.responses import Response
from datetime import datetime
import logging

logger = logging.getLogger("middleware_logger")

class LoggingMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        # Read request body (must cache as it's stream)
        body = await request.body()
        body_text = body.decode('utf-8') if body else ""

        # get request headers
        headers = request.headers

        # Client IP
        client_ip = request.client.host

        timeBegin = datetime.now()

        # Call the actual route
        response: Response = await call_next(request)

        timeEnd = datetime.now()
        timeDiff = timeEnd - timeBegin
        timeDiffSeconds = timeDiff.microseconds

        # Append response status to the log
        logger.info(
            f"[http][internal] {request.method} {response.status_code} {request.url.path} {timeDiffSeconds}ms - "
            f"Client IP: {client_ip} -"
            f"{headers} - "
            f"Request Body: {body_text} - "
        )

        return response
