from fastapi import FastAPI, Depends, HTTPException, status
from sqlalchemy.orm import Session
from .database import engine, Base,  get_db
from .models import User, Monitor, Check
from .schemas import MonitorBase, MonitorCreate, MonitorResponse

# создает таблицы при запуске
Base.metadata.create_all(bind=engine)

app = FastAPI(title="Uptime Monitoring API")


@app.get("/")
def root():
    return {"status": "ok", "service": "api", "message": "container is running"}


@app.get("/health")
def health():
    return {"status": "healthy"}


@app.post("/monitors", response_model = MonitorResponse)
def create_monitors(
    monitor: MonitorCreate,
    db: Session = Depends(get_db)
):
    db_monitor = Monitor(
        name = monitor.name, 
        url = str(monitor.url), 
        interval_seconds = monitor.interval_seconds
    )

    db.add(db_monitor)
    db.commit()
    db.refresh(db_monitor)
    return db_monitor


@app.get("/monitors", response_model=list[MonitorResponse])
def get_all_monitors(
    db: Session = Depends(get_db)
):

    all_monitors = db.query(Monitor).all()
    return all_monitors


@app.get("/monitors/{monitor_id}", response_model = MonitorResponse)
def get_monitor(
    monitor_id: int,
    db: Session = Depends(get_db)
):

    db_monitor = db.query(Monitor).filter(Monitor.id == monitor_id).first()
    if not db_monitor:
            raise HTTPException(
                status_code=status.HTTP_404_NOT_FOUND, 
                detail="not found"
            )

    return db_monitor
    


@app.put("/monitors/{monitor_id}", response_model = MonitorResponse)
def update_monitor(
    monitor_id: int,
    monitor_data: MonitorCreate,
    db: Session = Depends(get_db)
):

    db_monitor = db.query(Monitor).filter(Monitor.id == monitor_id).first()
    if not db_monitor:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND, 
            detail="not found"
        )

    db_monitor.name = monitor_data.name
    db_monitor.url = str(monitor_data.url)
    db_monitor.interval_seconds = monitor_data.interval_seconds

    db.commit()
    db.refresh(db_monitor)
    return db_monitor


@app.delete("/monitors/{monitor_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_monitors(
    monitor_id: int,
    db: Session = Depends(get_db)
):

    db_monitor = db.query(Monitor).filter(Monitor.id == monitor_id).first()
    if not db_monitor:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND, 
            detail="not found"
        )

    db.delete(db_monitor)
    db.commit()
    status_code=status.HTTP_204_NO_CONTENT
    return None