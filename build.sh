
sleep 0.1

go run ./cmd/examples/cors/preflight & dlv debug --init=dlv.init ./cmd/api \
-- -cors-trusted-origins "http://localhost:9000" 