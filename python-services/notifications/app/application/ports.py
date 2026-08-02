from typing import Protocol
from app.domain.entities import NotificationChannel, Alert


class NotificationChannelRepository(Protocol):
    """Контракт для работы с хранилищем каналов уведомлений."""

    def create(self, channel: NotificationChannel) -> NotificationChannel:
        ...

    def get_by_user_id(self, user_id: int) -> list[NotificationChannel]:
        ...

    def delete(self, channel_id: int) -> bool:
        ...


class AlertSender(Protocol):
    """Контракт для отправки уведомлений (Telegram, Email, Slack и т.д.)."""

    def send(self, destination: str, message: str) -> bool:
        ...