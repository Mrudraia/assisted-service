lint-swagger:
	spectral lint swagger.yaml

generate-go-client:
	rm -rf client/ models/*.go
	goswagger generate client --template=stratoscale -f swagger.yaml

generate-go-server:
	rm -rf restapi/
	goswagger generate server --template=stratoscale -f swagger.yaml

generate-python-client:
	mkdir -p ${BUILD_FOLDER}
	./hack/generate.sh generate_python_client

generate-mocks:
	find . -name 'mock_*.go' -type f -not -path './vendor/*' -delete
	go generate -v $(shell go list ./... | grep -v 'assisted-service/models\|assisted-service/client\|assisted-service/restapi')

generate-migration:
	go run ./tools/migration_generator/migration_generator.go -name=${MIGRATION_NAME}

generate-keys:
	mkdir -p ${BUILD_FOLDER}
	(cd ./tools && go run auth_keys_generator.go -keys-dir=${BUILD_FOLDER})

generate-events:
	rm -rf internal/common/events
	mkdir -p internal/common/events
	tools/generate_events.py ./docs/events.yaml internal/common/events/events.go

--remove-dashes-and-dots:
	find models/ -name "*.go" -maxdepth 1 -exec sed -i 's/Dash//g;s/Dot//g' {} \;

validate-swagger-file:
	./hack/generate.sh validate_swagger_file

generate-configuration:
	./hack/generate.sh generate_configuration

generate-manifests:
	mkdir -p ${BUILD_FOLDER}
	./hack/generate.sh generate_manifests

generate-crds:
	./hack/generate.sh generate_crds

generate-bundle:
	mkdir -p ${BUILD_FOLDER}
	./hack/generate.sh generate_bundle

generate-from-swagger: lint-swagger generate-go-client generate-go-server validate-swagger-file --remove-dashes-and-dots

generate-vendor:
	go mod vendor
	cd api && go mod vendor
	cd models && go mod vendor

generate: generate-vendor generate-from-swagger generate-events generate-mocks generate-configuration generate-bundle
