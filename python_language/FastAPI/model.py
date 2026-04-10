from pydantic import BaseModel

class products(BaseModel):
    id: int
    name: str
    price: float
    description: str
    quantity: int
