from fastapi import FastAPI, Depends
import hello_api.deps as deps
from hello_api.repo import RepositoryInterface

app = FastAPI()

repo = deps.redis_client


@app.get("/translate/{word}")
def translation(
    word: str, language: str = "english", repo: RepositoryInterface = Depends(repo)
):
    resp = repo.translate(language, word)
    return {"language": language.lower(), "translation": resp}
