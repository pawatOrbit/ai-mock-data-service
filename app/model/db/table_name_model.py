from pydantic import BaseModel

# Define Pydantic model
class GetTableName(BaseModel):
    table_name: str
    table_script: str
