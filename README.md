
## generate model script
```shell
goctl model mysql datasource -url="root:localpwd@tcp(localhost:3306)/gormzero" -table="users"  -dir="services/model" -cache=true --style=goZero --home ./template


goctl model mysql datasource -url="root:@tcp(localhost:3306)/data_flow" -table="maker"  -dir="services/model" -cache=true --style=go_zero --home ./template

goctl model mysql datasource -url="root:@tcp(localhost:3306)/data_flow" -table="maker_smart_address"  -dir="services/model" -cache=true --style=go_zero --home ./template

goctl model mysql datasource -url="root:@tcp(localhost:3306)/data_flow" -table="handle_trade_hash"  -dir="services/model" -cache=true --style=go_zero --home ./template

goctl model mysql datasource -url="root:@tcp(localhost:3306)/data_flow" -table="native_token"  -dir="services/model" -cache=true --style=go_zero --home ./template


goctl model mysql datasource -url="root:@tcp(localhost:3306)/data_flow" -table="pair_data_5m"  -dir="services/model" -cache=true --style=go_zero --home ./template


goctl model pg datasource -url="postgres://metaofo:123456@localhost:5432/pg?sslmode=disable" -table="evm_data"  -dir="services/model" -cache=true --style=go_zero --home ./template
```