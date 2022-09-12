from hello_api.app import app, repo
from fastapi.testclient import TestClient
from redislite import Redis

from hello_api.repo import RepositoryInterface
from hello_api.redis import RedisRepository
import unittest


class AppIntegrationTest(unittest.TestCase):
    def redis_client(self) -> RepositoryInterface:
        self.fake_redis = Redis()
        self.fake_redis.set("hello:german", "Hallo")
        self.fake_redis.set("hello:english", "Hello")

        return RedisRepository(client=self.fake_redis)

    def setUp(self):
        self.repo = self.redis_client()
        self.client = TestClient(app)
        app.dependency_overrides[repo] = self.redis_client

    def test_english_translation(self):
        response = self.client.get("/translate/hello")
        assert response.status_code == 200
        assert response.json() == {"language": "english", "translation": "Hello"}

    def test_german_translation(self):
        response = self.client.get("/translate/hello?language=GERMAN")
        assert response.status_code == 200
        assert response.json() == {"language": "german", "translation": "Hallo"}
