from dataclasses import dataclass
from datetime import datetime


@dataclass
class NotificationChannel:
    id: int | None
    user_id: int
    channel_type: str  # "telegram", "email", "slack"
    destination: str   # chat_id или email
    is_active: bool = True
    created_at: datetime | None = None


@dataclass
class Alert:
    id: int | None
    monitor_id: int
    error_message: str
    status: str        # "sent", "failed", "pending"
    status_code: int | None = None
    sent_at: datetime | None = None