
projects-gortener-up:
	$(MAKE) -C projects/gortener up

projects-gortener-down:
	$(MAKE) -C projects/gortener down

projects-gortener-migrate-status:
	$(MAKE) -C projects/gortener migrate-status

projects-gortener-migrate-up:
	$(MAKE) -C projects/gortener migrate-up

projects-gortener-database-connect:
	$(MAKE) -C projects/gortener database-connect

projects-gortener-database-dump:
	$(MAKE) -C projects/gortener database-dump

projects-gortener-check-dependencies:
	$(MAKE) -C projects/gortener check-dependencies

projects-gortener-start:
	$(MAKE) -C projects/gortener start

projects-gortener-start-server:
	$(MAKE) -C projects/gortener start-server