
sleep 0.1

go run ./cmd/examples/cors/simple --addr=":9001" & dlv debug --init=dlv.init ./cmd/api -- -cors-trusted-origins "http://localhost:9000 http://localhost:9001"

#go run ./cmd/api -cors-trusted-origins="http://localhost:9000 http://localhost:9001"

