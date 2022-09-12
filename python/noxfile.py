import nox


nox.options.sessions = "lint", "tests"
locations = "hello_api", "tests", "noxfile.py"


@nox.session
def tests(session):
    session.run("poetry", "install", external=True)
    session.run("pytest")


@nox.session
def lint(session):
    args = session.posargs or locations
    session.install("flake8", "flake8-black")
    session.run("flake8", *args)

@nox.session
def black(session):
    args = session.posargs or locations
    session.install("black")
    session.run("black", *args)
