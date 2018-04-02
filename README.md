# mockProbe
a skeleton probe for turbonomic


# Test it

### Compile
```bash
glide update --strip-vendor
make build
```

### connect to Turbonomic server
```bash
# edit conf/turbo.json to set the right address and credentials

./_output/mockProbe -v=3
```

