def build_sql_postgres_connection_string(
    host: str, port: str, dbname: str, user: str, password: str
) -> str:
    return f"postgresql://{user}:{password}@{host}:{port}/{dbname}"