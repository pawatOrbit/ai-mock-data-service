from fastapi import APIRouter
from app.model.http.table_schemas_req_resp import GetTableSchemasResponse
from app.service.table_schemas_service import get_table_schema_list_service

router = APIRouter(
    prefix="/table_schemas",
    tags=["Table Schemas"],
)

@router.post("/get_table_schemas_list", response_model=GetTableSchemasResponse)
async def get_table_schemas_list() -> GetTableSchemasResponse:
    """
    Get a list of table schemas.
    """
    return await get_table_schema_list_service()