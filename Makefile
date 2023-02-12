NAME         := goutils
BIN          := $(CURDIR)/bin
PROJECT_PATH := $(CURDIR)/cmd/$(NAME)

GO           := go
GOARCH       ?= $(shell go env GOARCH)
GOBIN        := $(HOME)/.local/bin
GOOS         ?= $(shell go env GOOS)

ifeq ($(GOOS), windows)
  EXT := .exe
else
  EXT :=
endif

COMPLETION   := $(ZSH_CUSTOM)/plugins/completions/_$(NAME)
TARGET       := $(BIN)/$(NAME)-$(GOOS)-$(GOARCH)$(EXT)

.PHONY: build
build: | $(BIN)
	cd $(PROJECT_PATH) && GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(TARGET)

.PHONY: install
install:
	cd $(PROJECT_PATH) && GOBIN=$(GOBIN) $(GO) install
	$(GOBIN)/$(NAME) completion zsh > $(COMPLETION)

.PHONY: clean
clean:
	$(RM) --recursive $(BIN)
	$(RM) $(COMPLETION)
	$(RM) $(GOBIN)/$(NAME)

$(BIN):
	mkdir -p $(BIN)
