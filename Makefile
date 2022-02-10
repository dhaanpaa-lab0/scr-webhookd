PLATFORMS := linux/amd64 windows/amd64 darwin/arm64
APP := webhookd

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

all: clean
	go build -o 'target/webhookd'

clean:
	rm -rfv target/

release: $(PLATFORMS)

$(PLATFORMS): clean
	GOOS=$(os) GOARCH=$(arch) go build -o 'target/$(APP).$(os)-$(arch)'

.PHONY: release $(PLATFORMS)