from sqlalchemy import Column, Integer, String, Boolean, DateTime, ForeignKey
from sqlalchemy.sql import func
from .database import Base


class NotificationChannel(Base):
    __tablename__ = "notification_channels"

    id = Column(Integer, primary_key=True, index=True)
    user_id = Column(Integer, index=True, nullable=False)
    channel_type =  Column(String, nullable=False)  # Telegram/email
    destination = Column(String, nullable=False)  # chat_id/email
    is_active = Column(Boolean, default=True)
    created_at = Column(DateTime, server_default=func.now())



class Alert(Base):
    __tablename__ = "alerts"

    id = Column(Integer, primary_key=True, index=True)
    monitor_id = Column(Integer, ForeignKey("monitors.id"))
    error_message = Column(String, nullable=False)
    status_code = Column(Integer, nullable=True)
    sent_at = Column(DateTime, server_default=func.now())
    status = Column(String, nullable=False)
