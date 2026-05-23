build:
	go build ./cmd/agentguard-sensor

test:
	go test ./...

install-linux:
	bash install/linux/install.sh

uninstall-linux:
	bash install/linux/uninstall.sh

install-macos:
	bash install/macos/install.sh

uninstall-macos:
	bash install/macos/uninstall.sh

package:
	tar -czf agentguard-sensor-installers.tgz install/
