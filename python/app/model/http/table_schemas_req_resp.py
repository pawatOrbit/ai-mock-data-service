from pydantic import BaseModel

class GetTableSchemasResponseData(BaseModel):
    table_schemas_names: list[str]


class GetTableSchemasResponse(BaseModel):
    status: int
    data: GetTableSchemasResponseData
