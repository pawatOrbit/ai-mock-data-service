from app.core.pgdb.db import database
from app.model.db.table_name_model import GetSchemaByTableNameModel, GetTableSchemaListModel

async def get_schema_by_table_name(table_name: GetSchemaByTableNameModel):
    query = "SELECT table_name, table_script FROM database_schemas WHERE table_name = :table_name"
    value = {"table_name": table_name}

    # Execute the query
    result = await database.fetch_one(query=query, values=value)
    if result is None:
        raise ValueError(f"Table {table_name} not found in the database_schemas table.")
    
    return GetSchemaByTableNameModel(**result)

async def get_table_schema_list()-> list[GetTableSchemaListModel]:
    query = "SELECT table_name FROM database_schemas"

    # Execute the query
    result = await database.fetch_all(query=query)
    if result is None:
        raise ValueError("No table names found in the database_schemas table.")
    
    return [GetTableSchemaListModel(**row) for row in result]