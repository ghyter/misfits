FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Install dependencies for Fyne
RUN apt-get update && \
    apt-get install -y \
        libgl1-mesa-dev \
        libxcursor-dev \
        libxrandr-dev \
        libxinerama-dev \
        libxi-dev \
        xorg-dev \ 
        libgtk-3-dev && \
    rm -rf /var/lib/apt/lists/*

# Set up environment variables for Fyne cross-compiling (optional)
ENV FYNE_CROSS_PLATFORM linux/amd64

# Set up a writable GOPATH
ENV GOPATH=/home/vscode/go
RUN mkdir -p $GOPATH && \
    chown -R vscode:vscode $GOPATH

# Use non-root user (vscode)
USER vscode

# Install the Fyne CLI
RUN go install fyne.io/fyne/v2/cmd/fyne@latest