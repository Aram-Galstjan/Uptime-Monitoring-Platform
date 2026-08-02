from pydantic import BaseModel, HttpUrl


class MonitorBase(BaseModel):
    name: str
    url: HttpUrl
    interval_seconds: int = 60


class MonitorCreate(MonitorBase):
    pass


class MonitorResponse(MonitorBase):
    id: int
    is_active: bool = True

    class Config:
        from_attributes = True