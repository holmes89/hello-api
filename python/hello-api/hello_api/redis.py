import redis
import os
from hello_api.repo import RepositoryInterface

class RedisRepository(RepositoryInterface):

    host: str = os.environ.get('DB_HOST', 'localhost')
    port: str = os.environ.get('DB_PORT', '6379')
    default_language: str = os.environ.get('DEFAULT_LANGUAGE', 'english')

    def __init__(self) -> None:
        self.client = redis.Redis(host=self.host, port=self.port)

    def translate(self, language: str, word: str) -> str:
        """translates word into given language"""
        lang = language.lower() if language is not None else self.default_language
        key = f'{word.lower()}:{lang}'
        return self.client.get(key)