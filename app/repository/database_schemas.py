from app.core.pgdb.db import database
from app.model.db.table_name_model import GetTableName

async def get_schema_by_table_name(table_name: GetTableName):
    query = "SELECT table_name, table_script FROM database_schemas WHERE table_name = :table_name"
    value = {"table_name": table_name}

    # Execute the query
    result = await database.fetch_one(query=query, values=value)
    if result is None:
        raise ValueError(f"Table {table_name} not found in the database_schemas table.")
    
    return GetTableName(**result)