VERSION=$(shell git describe --tags --abbrev=0)

EXT :=
ifeq ($(OS),Windows_NT)
	EXT := .exe
endif

.PHONY: docs debug coverage

all: docs tfformat compile checksum clean

test:
	go test -timeout 30s github.com/terraform-routeros/terraform-provider-routeros

docs:
	go generate ./...
	# !!! GNU Sed
	find docs -type f -exec sed -i -E '/^.*__[[:alpha:]_]+__/d' {} \;

tfformat:
	terraform fmt -recursive examples/

debug:
	go generate routeros/provider.go
	go build -gcflags="all=-N -l" -o terraform-provider-routeros_${VERSION}$(EXT) main.go

compile:
	mkdir -p pkg
	echo "Removing previously built packages"
	rm -rf pkg/*
	go generate routeros/provider.go
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_linux_arm.zip terraform-provider-routeros_${VERSION}
	
	GOOS=linux GOARCH=arm64 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_linux_arm64.zip terraform-provider-routeros_${VERSION}

	GOOS=linux GOARCH=386 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_linux_386.zip terraform-provider-routeros_${VERSION}

	GOOS=linux GOARCH=amd64 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_linux_amd64.zip terraform-provider-routeros_${VERSION}

	GOOS=windows GOARCH=amd64 go build -o terraform-provider-routeros_${VERSION}.exe main.go
	zip pkg/terraform-provider-routeros_${VERSION}_windows_amd64.zip terraform-provider-routeros_${VERSION}.exe

	GOOS=windows GOARCH=386 go build -o terraform-provider-routeros_${VERSION}.exe main.go
	zip pkg/terraform-provider-routeros_${VERSION}_windows_386.zip terraform-provider-routeros_${VERSION}.exe

	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_darwin_amd64.zip terraform-provider-routeros_${VERSION}

	GOOS=darwin GOARCH=arm64 go build -o terraform-provider-routeros_${VERSION} main.go
	zip pkg/terraform-provider-routeros_${VERSION}_darwin_arm64.zip terraform-provider-routeros_${VERSION}

checksum:
	cd pkg && sha256sum *.zip > terraform-provider-routeros_${VERSION}_SHA256SUMS

clean:
	rm terraform-provider-routeros_${VERSION}
# Dump every router in ROS_COVERAGE_HOSTS (space separated https:// URLs,
# credentials from ROS_USERNAME/ROS_PASSWORD) and write the coverage table.
coverage:
	@test -n "$(ROS_COVERAGE_HOSTS)" || { echo "set ROS_COVERAGE_HOSTS='https://router1 https://router2'"; exit 1; }
	cd tools/ros-coverage && files="" && \
	for h in $(ROS_COVERAGE_HOSTS); do f=$$(go run . dump -url $$h) || exit 1; files="$$files $$f"; done && \
	v=$$(echo $$files | sed -E 's/.*schema-([0-9.]+)-.*/\1/') && \
	go run . compare $$files > coverage-$$v.md && echo "tools/ros-coverage/coverage-$$v.md"
