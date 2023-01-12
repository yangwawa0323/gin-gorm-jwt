generate::
	@echo "make parse dependencies"
	swag init --parseDependency 

# .PHONY: docs
docs::
	echo "make parse dependencies"
	swag init --parseDependency --parseInternal

run::
	go run main.go