SHELL = /bin/bash

.PHONY: release
release: ## example: make release V=0.0.0
	@read -p "Press enter to confirm and push to origin ..."
	git tag v$(V)
	git push origin v$(V)
