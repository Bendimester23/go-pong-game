clean:
	rm -rf build/*

win:
	echo "Building for Windows"
	CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CGO_LDFLAGS="-static-libgcc -static -lpthread" go build -o build/pong-win.exe .

linux:
	echo "Building for Linux"
	GOARCH=amd64 CGO_ENABLED=1 go build -o build/pong-linux .

all: win linux

push-release-win: clean win
	butler push build/pong-win.exe bendimester23/bendi-pong:win-release-v$(cat version)

push-release-linux: clean linux
	butler push build/pong-linux bendimester23/bendi-pong:linux-release-v$(cat version)

push-beta-win: clean win
	butler push build/pong-win.exe bendimester23/bendi-pong:win-beta-v$(cat version)

push-beta-linux: clean linux
	butler push build/pong-linux bendimester23/bendi-pong:linux-beta-v$(cat version)

push-release-all: push-release-win push-release-linux
	echo "Pushing Release-$(cat version)" | lolcat
	git tag -a v$(cat version) -m "Release v$(cat version)"

push-beta-all: push-beta-win push-beta-linux
	echo "Pushing Beta-$(cat version)" | lolcat
	git tag -a v$(cat version) -m "Beta v$(cat version)"
