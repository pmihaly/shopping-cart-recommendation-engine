import asyncio
import os

import httpx
import structlog

logger = structlog.stdlib.get_logger()

logger.info("Started db initianization")

runtime_api = os.environ["AWS_LAMBDA_RUNTIME_API"]


res = httpx.get(f"http://{runtime_api}/2018-06-01/runtime/invocation/next")
request_id = res.headers["Lambda-Runtime-Aws-Request-Id"]

logger.info("Got request id", request_id=request_id)

import flipkart_products_to_db

logger.info("Products initialized")

import generate_carts

asyncio.run(generate_carts.main())

logger.info("Carts initialized")

httpx.post(
    f"http://{runtime_api}/2018-06-01/runtime/invocation/{request_id}/response",
)
