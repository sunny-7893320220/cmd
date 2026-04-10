from itertools import product

from fastapi import Depends, FastAPI
from Databases_modle import Base
import Databases_modle
from model import products
from db import sessionLocal
from db import sessionLocal,engine
Base.metadata.create_all(bind=engine)
app = FastAPI()



def get_db():
    db = sessionLocal()
    try:
        yield db
    finally:
        db.close()

@app.get("/")
def read_root():
   return {"message": "Welcome to FastAPI"}

product = [
    products(id=1,name="PC",price=10,description="laptops",quantity=1),
    products(id=2,name="Videos",price=10,description="PC_for_laptops",quantity=2),
    products(id=3,name="sounds",price=10,description="This_is_the_sound_system",quantity=1)
    #$ docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=pass123 postgres

]
def inti_db():
    session = sessionLocal()
    count = session.query(Databases_modle.products).count()
    if count == 0:

        for i in product:
            session.add(Databases_modle.products(**i.model_dump()))
    session.commit()
inti_db()

@app.get("/products/")

def read_product(db: sessionLocal = Depends(get_db)):
   DB_results = db.query(Databases_modle.products).all()
   return DB_results

@app.get("/products/{id}")
def read_product(id: int, db: sessionLocal = Depends(get_db)):
    DB_result = db.query(Databases_modle.products).filter(Databases_modle.products.id == id).first()
    return DB_result


@app.post("/products/")
def create_product(products: products, db: sessionLocal = Depends(get_db)):
    new_product = db.add(Databases_modle.products(**products.model_dump()))
    db.commit()
    return products

@app.put("/products/{id}")
def update_product(id: int, updated_product: products, db: sessionLocal = Depends(get_db)):
    product_to_update = db.query(Databases_modle.products).filter(Databases_modle.products.id == id).first()
    if product_to_update:
        product_to_update.name = updated_product.name
        product_to_update.price = updated_product.price
        product_to_update.description = updated_product.description
        product_to_update.quantity = updated_product.quantity
        db.commit()
        return {"message": "Product updated successfully"}
    return {"message": "Product not found"}


@app.delete("/products/{id}")
def delete_product(id: int, db: sessionLocal = Depends(get_db)):
    product_to_delete = db.query(Databases_modle.products).filter(Databases_modle.products.id == id).first()
    if product_to_delete:
        db.delete(product_to_delete)
        db.commit()
        return {"message": "Product deleted successfully"}
    return {"message": "Product not found"} 