NAME       := goutils
BIN        := bin
COMPLETION := $(ZSH_CUSTOM)/plugins/completions/_$(NAME)
GO         := go
GOBIN      := $(HOME)/.local/bin

.PHONY: build
build: $(BIN)/$(NAME)

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

$(BIN)/$(NAME): | $(BIN)
	go build -o $(BIN) ./...

$(GOBIN)/$(NAME):
	GOBIN=$(GOBIN) $(GO) install ./...

$(COMPLETION): $(BIN)/$(NAME)
	$< completion zsh > $(COMPLETION)
