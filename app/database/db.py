import databases
from app.utils.db_path_builder import build_sql_postgres_connection_string
from app.config import userNamedb, password, host, port, database_name

connectionString = build_sql_postgres_connection_string(
    host=host,
    port=port,
    dbname=database_name,
    user=userNamedb,
    password=password,
)

database = databases.Database(connectionString)