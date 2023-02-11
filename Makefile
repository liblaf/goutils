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
build: $(TARGET)

.PHONY: install
install: $(GOBIN)/$(NAME) completion

.PHONY: completion
completion: $(COMPLETION)

.PHONY: clean
clean:
	$(RM) --recursive $(BIN)
	$(RM) $(COMPLETION)
	$(RM) $(GOBIN)/$(NAME)

$(BIN):
	mkdir -p $(BIN)

$(TARGET): | $(BIN)
	cd cmd/goutils && $(GO) build -o $(TARGET)

$(GOBIN)/$(NAME):
	cd cmd/goutils && GOBIN=$(GOBIN) $(GO) install

$(COMPLETION): $(GOBIN)/$(NAME)
	$< completion zsh > $(COMPLETION)
