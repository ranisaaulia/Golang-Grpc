#!/bin/bash
path=""
cmd=""

# Parse command-line options and arguments
while [[ $# -gt 0 ]]; do
  case "$1" in
    --path=*)
      path="${1#*=}"
      shift
      ;;
    --cmd=*)
      cmd="${1#*=}"
      shift
      ;;
    *)
      echo "Invalid argument: $1"
      exit 1
      ;;
  esac
done

if [ -z "$cmd" ]; then
  echo "Error: --cmd arguments must be filled."
  exit 1
fi

if [ "$cmd" == "sqlc" ]; then
  if [ -z "$path" ]; then
    echo "Error: --path arguments must be filled."
    exit 1
  fi
  # Check if directory exist
  if [ ! -d "$path" ]; then
    echo "Error: Directory path '$path' does not exist."
    exit 1
  fi
  # Check if schema.sql is exist
  if [ ! -f "$path/sqlc/schema.sql" ]; then
    echo "Error: File path '$path/sqlc/schema.sql' does not exist."
    exit 1
  fi
  # Check if queries.sql is exist
  if [ ! -f "$path/sqlc/queries.sql" ]; then
    echo "Error: File path '$path/sqlc/queries.sql' does not exist."
    exit 1
  fi
  # Check if sqlc.yaml exists
  if [ ! -f "$path/sqlc/sqlc.yaml" ]; then
    last_segment=$(basename "$path")_sql
    # If sqlc.yaml doesn't exist, create it using the provided arguments
    cat <<EOF > "$path/sqlc/sqlc.yaml"
version: "1"
project:
    id: ""
packages: [
    {
        name: "$last_segment",
        path: ".",
        queries: "queries.sql",
        schema: "schema.sql",
        engine: "postgresql",
        emit_prepared_queries: true,
        emit_interface: false,
        emit_exact_table_names: false,
        emit_empty_slices: false,
        emit_exported_queries: false,
        emit_json_tags: true,
        emit_result_struct_pointers: false,
        emit_params_struct_pointers: false,
        emit_methods_with_db_argument: false,
        emit_enum_valid_method: false,
        emit_all_enum_values: false,
        json_tags_case_style: "pascal",
        output_db_file_name: "db.go",
        output_models_file_name: "models.go",
    }
]
EOF
    echo "sqlc.yaml file created, generating the sqlc"
    sqlc generate -f $path/sqlc/sqlc.yaml
    echo "sqlc generated on $path/sqlc"
  elif [ -f "$path/sqlc/sqlc.yaml" ]; then
    echo "generating the sqlc"
    sqlc generate -f $path/sqlc/sqlc.yaml
    echo "sqlc generated on $path/sqlc"
  else
    echo "command not exist"
    exit 1
  fi
elif [ "$cmd" == "proto" ]; then
  if [ -z "$path" ]; then
    echo "Error: --path arguments must be filled."
    exit 1
  fi
  # Check if directory exist
  if [ ! -d "$path/proto" ]; then
    echo "Creating directory $path/proto"
    mkdir $path/proto
  fi
  echo "Executing proto command..."
  protoc \
    --proto_path=. \
    --go-grpc_out . \
    --go-grpc_opt paths=source_relative \
    --micro_out=. \
    --micro_opt paths=source_relative \
    --go_out=:. \
    --go_opt paths=source_relative \
    --validate_out="lang=go:." \
    --validate_opt paths=source_relative $path/proto/*.proto
  echo "Proto generated at $path/proto"
elif [ "$cmd" == "init" ]; then
  go get -u google.golang.org/protobuf/proto
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest
	go get -d github.com/envoyproxy/protoc-gen-validate
	go install github.com/envoyproxy/protoc-gen-validate
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get go-micro.dev/v4/registry/cache@v4.4.0
  docker run -d \
    --name postgres-container \
    -e POSTGRES_USER=user \
    -e POSTGRES_PASSWORD=P@ssw0rd \
    -e POSTGRES_DB=go_grpc \
    -p 6543:5432 \
    postgres:11.1-alpine

else
  echo "Error: Unknown command '$cmd'. Supported commands are 'sqlc', 'init' and 'proto'."
  exit 1
fi




