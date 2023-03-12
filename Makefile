dependency:
	go mod tidy

.PHONY: remove_local
remove_local:
	git remote update --prune
	git checkout origin/main
	git for-each-ref --format '%(refname:short)' refs/heads | xargs git branch -D

.PHONY: test
test:
	go test -cover ./...