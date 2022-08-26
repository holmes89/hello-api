from fastapi import FastAPI, Depends
import os
import hello_api.deps as deps
from hello_api.repo import RepositoryInterface

app = FastAPI()

@app.get("/translate/{word}")
def translation(word: str, language: str = 'english', repo: RepositoryInterface = Depends(deps.redis_client)):
    resp = repo.translate(language, word)
    return {"language": language.lower(), "translation": resp}