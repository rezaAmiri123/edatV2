module github.com/rezaAmiri123/edatV2

go 1.20

replace (
	github.com/jackc/pgconn latest => github.com/jackc/pgconn v1.13.0
	github.com/jackc/pgx/pgtype latest => github.com/jackc/pgx/pgtype v1.12.0
)

require (
	github.com/google/uuid v1.6.0
	github.com/jackc/pgconn v1.14.3
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa
	github.com/stackus/errors v0.1.5
	github.com/stretchr/testify v1.8.4
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.38.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
