import os
from typing import TYPE_CHECKING

import psycopg2
import structlog
from boto3.session import Session
from mypy_boto3_secretsmanager import SecretsManagerClient

logger = structlog.stdlib.get_logger()

aws_region = os.environ["AWS_REGION"]

session = Session(region_name=aws_region)
logger.info("Got boto3 session")

client = session.client("secretsmanager")

if TYPE_CHECKING:
    assert isinstance(client, SecretsManagerClient)

password_arn = os.environ["PGPASSWORD_SECRET_ARN"]

assert password_arn is not None

response = client.get_secret_value(SecretId=password_arn)
password = response["SecretString"]

logger.info("Got secret value from AWS Secrets Manager")

user = os.environ["PGUSER"]
host = os.environ["PGHOST"]
port = os.environ["PGPORT"]
database = os.environ["PGDATABASE"]

logger.info(
    "Got environment variables", user=user, host=host, port=port
)

conn = psycopg2.connect(
    user=user, password=password, host=host, port=port
)
cur = conn.cursor()
logger.info("Connected to PostgreSQL")

conn.autocommit = True

cur.execute(f"drop database if exists {database}")
logger.info("Deleted database", database=database)

conn = psycopg2.connect(
    user=user, password=password, host=host, port=port, database=database
)
cur = conn.cursor()

sql_files = os.listdir("/postgres")

logger.debug("Executing all sql files in /postgres", sql_files=sql_files)
for sql_file in sql_files:
    with open(f"/postgres/{sql_file}") as f:
        cur.execute(f.read())
        logger.info("Executed SQL file", sql_file=sql_file)

conn.commit()
