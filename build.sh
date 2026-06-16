
sleep 0.1

go run ./cmd/examples/cors/preflight & dlv debug --init=dlv.init ./cmd/api \
-- -port 4000 -limiter-enabled false
#-- -cors-trusted-origins "http://localhost:9000" 