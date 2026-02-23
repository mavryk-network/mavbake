#!/bin/sh

TMP_NAME="./$(head -c 24 /dev/urandom | base64 | tr -dc 'a-zA-Z0-9' | head -c 32)"
PRERELEASE=false
if [ "$1" = "--prerelease" ]; then
	PRERELEASE=true
fi

if which curl >/dev/null; then
	if curl --help 2>&1 | grep "--progress-bar" >/dev/null 2>&1; then
		PROGRESS="--progress-bar"
	fi

	set -- curl -L $PROGRESS -o "$TMP_NAME"
	if [ "$PRERELEASE" = true ]; then
		LATEST=$(curl -sL https://api.github.com/repos/mavryk-network/mavbake/releases | grep tag_name | sed 's/  "tag_name": "//g' | sed 's/",//g' | head -n 1 | tr -d '[:space:]')
	else
		LATEST=$(curl -sL https://api.github.com/repos/mavryk-network/mavbake/releases/latest | grep tag_name | sed 's/  "tag_name": "//g' | sed 's/",//g' | tr -d '[:space:]')
	fi
else
	if wget --help 2>&1 | grep "--show-progress" >/dev/null 2>&1; then
		PROGRESS="--show-progress"
	fi
	set -- wget -q $PROGRESS -O "$TMP_NAME"
	if [ "$PRERELEASE" = true ]; then
		LATEST=$(wget -qO- https://api.github.com/repos/mavryk-network/mavbake/releases | grep tag_name | sed 's/  "tag_name": "//g' | sed 's/",//g' | head -n 1 | tr -d '[:space:]')
	else
		LATEST=$(wget -qO- https://api.github.com/repos/mavryk-network/mavbake/releases/latest | grep tag_name | sed 's/  "tag_name": "//g' | sed 's/",//g' | tr -d '[:space:]')
	fi
fi

if mavbake version 2>/dev/null | grep "$LATEST" >/dev/null 2>&1; then
	echo "latest mavbake already available"
	exit 0
fi

PLATFORM=$(uname -m)
UNAME=$(uname -s | tr '[:upper:]' '[:lower:]')
OS=linux
if [ "$UNAME" = "darwin" ]; then
	mkdir -p /usr/local/bin
	OS=macos
fi

if [ "$PLATFORM" = "x86_64" ]; then
	PLATFORM="amd64"
elif [ "$PLATFORM" = "aarch64" ] || [ "$PLATFORM" = "arm64" ]; then
	PLATFORM="arm64"
else
	echo "unsupported platform: $PLATFORM" 1>&2
	exit 1
fi

BIN="mavbake"
rm -f "/usr/local/bin/$BIN"
rm -f "/usr/bin/$BIN"
rm -f "/bin/$BIN"
rm -f "/usr/local/sbin/$BIN"
rm -f "/usr/sbin/$BIN"
rm -f "/sbin/$BIN"
# check destination folder
if [ -w "/usr/local/bin" ]; then
    DESTINATION="/usr/local/bin/$BIN"
elif [ -w "/usr/local/sbin" ]; then
    DESTINATION="/usr/local/sbin/$BIN"
elif [ -w "/usr/bin" ]; then
    DESTINATION="/usr/bin/$BIN"
elif [ -w "/usr/sbin" ]; then
    DESTINATION="/usr/sbin/$BIN"
elif [ -w "/bin" ]; then
    DESTINATION="/bin/$BIN"
elif [ -w "/sbin" ]; then
    DESTINATION="/sbin/$BIN"
else
    echo "No writable system binary directory found, installing locally."
    DESTINATION="./$BIN"
fi

if [ "$PRERELEASE" = true ]; then
	echo "downloading latest mavbake prerelease for $PLATFORM..."
else
	echo "downloading mavbake-$OS-$PLATFORM $LATEST..."
fi

if "$@" "https://github.com/mavryk-network/mavbake/releases/download/$LATEST/mavbake-$OS-$PLATFORM" &&
	mv "$TMP_NAME" "$DESTINATION" &&
	chmod +x "$DESTINATION"; then
	if [ "$1" = "--prerelease" ]; then
		echo "latest mavbake prerelease for $PLATFORM successfully installed"
	else
		echo "mavbake $LATEST for $PLATFORM successfully installed"
	fi
	hash -r 2>/dev/null || true
else
	echo "mavbake installation failed!" 1>&2
	exit 1
fi
