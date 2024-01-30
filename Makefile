
USER_GH=eyedeekay
VERSION=0.33.8
packagename=edgar

echo: fmt
	@echo "type make version to do release $(VERSION)"

version:
	github-release release -s $(GITHUB_TOKEN) -u $(USER_GH) -r $(packagename) -t v$(VERSION) -d "version $(VERSION)"; sleep 3s

upload:
	CGO_ENABLED=0 go build -o $(packagename)
	github-release upload -s $(GITHUB_TOKEN) -u $(USER_GH) -r $(packagename) -t v$(VERSION) -f "$(packagename)" -l "`sha256sum $(packagename)`" -n "$(packagename)"

del:
	github-release delete -s $(GITHUB_TOKEN) -u $(USER_GH) -r $(packagename) -t v$(VERSION)

tar:
	tar --exclude .git \
		--exclude .go \
		--exclude bin \
		--exclude examples \
		-cJvf ../$(packagename)_$(VERSION).orig.tar.xz .

link:
	rm -f ../goSam
	ln -sf . ../goSam

fmt:
	find . -name '*.go' -exec gofmt -w -s {} \;