
NODE           = node
DOCKER         = docker
CONTAINER_NAME = guake-cl

ESLINT         = ./node_modules/.bin/eslint
ESLINT_FLAGS   = --config config/eslint.json





install: snap
	cd snapcraft && snap install guake-cl* && cd ..
#	sudo cp guake-cl /etc/bash_completion.d/guake-cl

snap:
	cd snapcraft && snapcraft clean && snapcraft snap && cd ..

docker-build:
	$(DOCKER) build --tag=$(CONTAINER_NAME) .

docker-cleanbuild:
	$(DOCKER) build --no-cache=true --tag=$(CONTAINER_NAME) .

docker-run:
	$(DOCKER) run $(CONTAINER_NAME)

eslint:
	$(ESLINT) $(ESLINT_FLAGS) ./src

screenshot-images:
	bash scripts/capture-screenshots.sh

assemble-images:
	bash scripts/assemble-images.sh
