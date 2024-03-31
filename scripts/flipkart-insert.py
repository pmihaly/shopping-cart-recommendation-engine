import os
from typing import TYPE_CHECKING

import psycopg2
from boto3.session import Session
from icecream import ic
from mypy_boto3_secretsmanager import SecretsManagerClient

aws_region = os.environ.get("AWS_REGION")
assert aws_region is not None

session = Session(region_name=aws_region)

client = session.client("secretsmanager")

if TYPE_CHECKING:
    assert isinstance(client, SecretsManagerClient)


password_arn = os.environ.get("PGPASSWORD_SECRET_ARN")
assert password_arn is not None

response = client.get_secret_value(SecretId=password_arn)
password = response["SecretString"]

user = os.environ.get("PGUSER")
assert user is not None

host = os.environ.get("PGHOST")
assert host is not None

port = os.environ.get("PGPORT")
assert port is not None

database = os.environ.get("PGDATABASE")
assert database is not None

conn = psycopg2.connect(
    user=user, password=password, host=host, port=port, database=database
)
cur = conn.cursor()

cur.execute("select 1")
ic(cur.fetchone())
