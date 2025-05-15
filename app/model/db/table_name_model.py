from pydantic import BaseModel

# Define Pydantic model
class GetSchemaByTableNameModel(BaseModel):
    table_name: str
    table_script: str

class GetTableSchemaListModel(BaseModel):
    table_name: str
