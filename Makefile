.PHONY: lint format test

# Thin wrappers so the standard commands run from the repo root.
# The scripts in scripts/ stay the source of truth; CI calls these targets.

lint:
	./scripts/lint

format:
	./scripts/format

test:
	./scripts/test
