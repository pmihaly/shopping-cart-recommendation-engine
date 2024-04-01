import os
from typing import TYPE_CHECKING

import httpx
import psycopg2
import structlog
from boto3.session import Session
from mypy_boto3_secretsmanager import SecretsManagerClient

logger = structlog.stdlib.get_logger()

res = httpx.get(
    f"http://{os.environ['AWS_LAMBDA_RUNTIME_API']}/2018-06-01/runtime/invocation/next"
)
request_id = res.headers["Lambda-Runtime-Aws-Request-Id"]
logger.info("Started execution", request_id=request_id)

aws_region = os.environ.get("AWS_REGION")
assert aws_region is not None

session = Session(region_name=aws_region)
logger.info("Got boto3 session")

client = session.client("secretsmanager")

if TYPE_CHECKING:
    assert isinstance(client, SecretsManagerClient)

password_arn = os.environ.get("PGPASSWORD_SECRET_ARN")

assert password_arn is not None

response = client.get_secret_value(SecretId=password_arn)
password = response["SecretString"]

logger.info("Got secret value from AWS Secrets Manager")

user = os.environ.get("PGUSER")
assert user is not None

host = os.environ.get("PGHOST")
assert host is not None

port = os.environ.get("PGPORT")
assert port is not None

database = os.environ.get("PGDATABASE")
assert database is not None

logger.info(
    "Got environment variables", user=user, host=host, port=port, database=database
)

conn = psycopg2.connect(
    user=user, password=password, host=host, port=port, database=database
)
cur = conn.cursor()
logger.info("Connected to PostgreSQL")

sql_files = os.listdir("/postgres")

logger.debug("Executing all sql files in /postgres", sql_files=sql_files)
for sql_file in sql_files:
    with open(f"/postgres/{sql_file}") as f:
        cur.execute(f.read())
        logger.info("Executed SQL file", sql_file=sql_file)

conn.commit()

httpx.post(
    f"http://{os.environ['AWS_LAMBDA_RUNTIME_API']}/2018-06-01/runtime/invocation/{request_id}/response",
)
