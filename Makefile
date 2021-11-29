clean:
	rm -rf build/*

win:
	echo "Building for Windows"
	CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CGO_LDFLAGS="-static-libgcc -static -lpthread" go build -o build/pong-win-v$(shell cat version).exe .

linux:
	echo "Building for Linux"
	GOARCH=amd64 CGO_ENABLED=1 go build -o build/pong-linux-v$(shell cat version) .

all: win linux

beta:
	mv build/pong-win-v-v$(shell cat version).exe build/pong-win-v-v$(shell cat version)-beta.exe
	mv build/pong-linux-v$(shell cat version) build/pong-linux-v$(shell cat version)-beta

release:
	mv build/pong-win-v-v$(shell cat version).exe build/pong-win-v-v$(shell cat version)-release.exe
	mv build/pong-linux-v$(shell cat version) build/pong-linux-v$(shell cat version)-release

push-release-win: clean win
	butler push build/pong-win-v$(shell cat version).exe bendimester23/bendi-pong:win-release

push-release-linux: clean linux
	butler push build/pong-linux-v$(shell cat version) bendimester23/bendi-pong:linux-release

push-beta-win: clean win
	butler push build/pong-win-v$(shell cat version).exe bendimester23/bendi-pong:win-beta

push-beta-linux: clean linux
	butler push build/pong-linux-v$(shell cat version) bendimester23/bendi-pong:linux-beta

push-release-all: push-release-win push-release-linux
	echo "Pushing Release-$(shell cat version)" | lolcat
	git tag -a v$(shell cat version) -m "Release v$(shell cat version)"

push-beta-all: push-beta-win push-beta-linux
	echo "Pushing Beta-$(shell cat version)" | lolcat
	git tag -a v$(shell cat version) -m "Beta v$(shell cat version)"
