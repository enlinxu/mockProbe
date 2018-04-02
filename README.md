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

# Write your own Probe

### 1. define the SupplyChain
The [default supply chain](https://github.com/songbinliu/mockProbe/blob/1f3976c29bd0342b4dff239c275df92573bb6f1c/pkg/registration/supply_chain_factory.go#L41) is the one used in [virtualCluster](https://github.com/songbinliu/virtualCluster).  


### 2. implement the discovery method
The [default discovery](https://github.com/songbinliu/mockProbe/blob/1f3976c29bd0342b4dff239c275df92573bb6f1c/pkg/discovery/discovery_client.go#L74) returns an empty list. 

### 3. implement the action Handler
The [default action handler] does nothing but log the action information.
