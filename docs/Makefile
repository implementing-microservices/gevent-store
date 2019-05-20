default: docs

.PHONY: docs
docs:
	#- docker build -t loc/gevent-api-docs .
	#- docker run -v ${PWD}:/apidocs loc/gevent-api-docs
	- docker run --name gevent-api-docs-builder -v ${PWD}:/apidocs irakli/redoc-cli:latest && docker rm -f gevent-api-docs-builder

.PHONY: start
start:
	#- docker build -t loc/gevent-api-docs .
	#- docker run -d -p 0.0.0.0:3939:8000/tcp --rm --name gevent-api-docs loc/gevent-api-docs /usr/local/bin/redoc-cli serve -w -ssr -p 8000 api.yml
	- docker run -d --rm --name gevent-api-docs -p 3939:80 -v ${PWD}/api.yml:/usr/share/nginx/html/swagger.yaml -e SPEC_URL=swagger.yaml redocly/redoc:v2.0.0-rc.8-1
	- @echo "server started at: http://0.0.0.0:3939"

.PHONY: stop
stop:
	- docker stop gevent-api-docs