BINARY := marginalia
CMD := ./cmd/marginalia
WEB_DIR := internal/server/web

GO_BUILD_FLAGS := -trimpath -ldflags='-s -w'
export CGO_ENABLED := 0

.PHONY: build frontend backend run clean

## build: build the frontend then the backend into a single static binary
build: frontend backend

## frontend: install deps and build the embedded web assets (web/dist)
frontend:
	cd $(WEB_DIR) && pnpm install --frozen-lockfile && pnpm build

## backend: build the Go binary (requires web/dist to exist for go:embed)
backend:
	go build $(GO_BUILD_FLAGS) -o $(BINARY) $(CMD)

## run: run the server from source
run:
	go run $(CMD)

## clean: remove build artifacts
clean:
	rm -f $(BINARY)
	rm -rf $(WEB_DIR)/dist
