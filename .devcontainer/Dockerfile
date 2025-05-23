# Use the official Microsoft dev container image for Go (1.23)
FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Optional: Install PostgreSQL client (e.g., psql) and other common tools
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
    postgresql-client \
    curl \
    git \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Optional: Install Go tools (replace with actual tools if needed)
# USER vscode
# RUN go install github.com/volatiletech/sqlboiler/v4@latest
# USER root
