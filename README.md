
Visit app at http://localhost:80

## Build

`docker build -t flask-vuejs-docker .`

## Deploy

`docker run -p 80:8080 flask-vuejs-docker`

## Running the tests

`pipenv install --dev`

`pipenv run python -m pytest`