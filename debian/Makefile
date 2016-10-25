pkg = petfeeder
version_base = 0.0.1

ifndef BUILD_NUMBER
version = $(version_base)-manualbuild
else
version = $(version_base)-$(BUILD_NUMBER)
endif

all: clean
	@sed -i "/^Version:/c\Version: $(version)" $(pkg)/DEBIAN/control
	@(cd ../ && make)
	@mkdir -p $(pkg)/usr/bin
	@mkdir -p $(pkg)/usr/sbin
	@mkdir -p $(pkg)/etc/petfeeder
	@(install -m 0755 ../cmd/manualfeed/manualfeed $(pkg)/usr/bin/)
	@(install -m 0755 ../petfeeder $(pkg)/usr/sbin/)
	@(install -m 0644 ../config.yaml $(pkg)/etc/petfeeder/)
	@find $(pkg) -name DEBIAN -prune -o -type f -exec md5sum {} \; > $(pkg)/DEBIAN/md5sums
	@dpkg-deb -b $(pkg) $(pkg)-$(version).deb
clean:
	@rm -rf $(pkg)/usr
	@rm -rf $(pkg)/etc/petfeeder
	@rm -f $(pkg)/DEBIAN/md5sums
	@rm -f $(pkg)-$(version).deb
