from datetime import datetime
from pydantic import BaseModel, ConfigDict


class NotificationChannel(BaseModel):
    channel_type: str
    destination: str
    is_active: bool=True


class NotificationChannelCreate(NotificationChannel):
    user_id: int


class NotificationChannelResponse(NotificationChannel):
    id: int
    user_id: int
    created_at: datetime

    model_config = ConfigDict(from_attributes=True)


class AlertBase(BaseModel):
    monitor_id: int
    error_messege: str
    status_code: int | None=None
    status: str


class AlertCreate(AlertBase):
    pass 


class AlertResponse(AlertBase):
    id: int
    sent_at: datetime

    model_config = ConfigDict(from_attributes=True)