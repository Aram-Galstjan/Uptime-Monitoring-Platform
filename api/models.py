from sqlalchemy import Column, Integer, String, Boolean, DateTime, ForeignKey
from sqlalchemy.sql import func
from .database import Base


class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True, index=True)
    email = Column(String, unique=True, index=True, nullable=False)


class Monitor(Base):
    __tablename__ = "monitors"

    name = Column(String, nullable=True)
    id = Column(Integer, primary_key=True, index=True)
    url = Column(String, nullable=False)
    interval_seconds = Column(Integer, default=60)
    is_active = Column(Boolean, default=True)


class Check(Base):
    __tablename__ = "checks"

    id = Column(Integer, primary_key=True, index=True)
    monitor_id = Column(Integer, ForeignKey("monitors.id"))
    status_code = Column(Integer, nullable=True)
    response_time_ms = Column(Integer, nullable=True)
    created_at = Column(DateTime(timezone=True), server_default=func.now())