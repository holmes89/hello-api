from hello_api.repo import RepositoryInterface
from hello_api.redis import RedisRepository


def redis_client() -> RepositoryInterface:
    return RedisRepository()
