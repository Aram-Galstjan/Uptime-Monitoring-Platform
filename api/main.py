from fastapi import FastAPI
from .database import engine, Base
from .models import User, Monitor, Check

# создает таблицы при запуске
Base.metadata.create_all(bind=engine)

app = FastAPI(title="Uptime Monitoring API")

@app.get("/")
def home():
    return {"status": "ok", "message": "работает"}