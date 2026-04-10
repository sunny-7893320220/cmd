from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine

engine = create_engine("postgresql://postgres:pass123@localhost:5432/postgres")

sessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

