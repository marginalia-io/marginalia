BINARY := marginalia
CMD := ./cmd/marginalia
WEB_DIR := frontend
EMBED_DIST := internal/server/embed/dist
WGO := github.com/bokwoon95/wgo@latest

GO_BUILD_FLAGS := -trimpath -ldflags='-s -w'
export CGO_ENABLED := 0

.PHONY: build frontend backend run dev dev-frontend dev-backend clean

## build: build the frontend then the backend into a single static binary
build: frontend backend

## frontend: install deps and build the embedded web assets (-> $(EMBED_DIST))
frontend:
	cd $(WEB_DIR) && pnpm install --frozen-lockfile && pnpm build

## backend: build the Go binary (requires $(EMBED_DIST) to exist for go:embed)
backend:
	go build $(GO_BUILD_FLAGS) -o $(BINARY) $(CMD)

## run: run the server from source
run:
	go run $(CMD)

## dev: run the frontend (Vite HMR) and backend (live-reload) together
dev:
	@$(MAKE) -j2 dev-frontend dev-backend

## dev-frontend: Vite dev server with HMR (proxies /api to the backend)
dev-frontend:
	cd $(WEB_DIR) && pnpm dev

## dev-backend: live-reloading backend; -tags dev skips the frontend embed
dev-backend:
	go run $(WGO) run -tags dev $(CMD)

## clean: remove build artifacts
clean:
	rm -f $(BINARY)
	rm -rf $(EMBED_DIST)
