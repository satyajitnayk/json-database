# API Reference

### `New(dir string, options *Options) (*Driver, error)`

Creates a new database instance. If the specified directory doesn't exist, it will be created.

- Parameters:

  - `dir`: The directory where the database will be stored.
  - `options`: Options for configuring the database, including a custom logger.

- Returns:
  - `*Driver`: The database driver instance.
  - `error`: Any error that occurred during database creation.

### `Write(collection string, resource string, v interface{}) error`

Writes a JSON-encoded record to the specified collection.

- Parameters:

  - `collection`: The name of the collection.
  - `resource`: The name of the resource (record).
  - `v`: The data to be written, typically a struct.

- Returns:
  - `error`: Any error that occurred during the write operation.

### `Read(collection string, resource string, v interface{}) error`

Reads a JSON-encoded record from the specified collection.

- Parameters:

  - `collection`: The name of the collection.
  - `resource`: The name of the resource (record).
  - `v`: A pointer to the variable where the data will be unmarshaled.

- Returns:
  - `error`: Any error that occurred during the read operation.

### `ReadAll(collection string) ([]string, error)`

Reads all records from the specified collection.

- Parameters:

  - `collection`: The name of the collection.

- Returns:
  - `[]string`: A slice of JSON-encoded records.
  - `error`: Any error that occurred during the read operation.

### `Delete(collection string, resource string) error`

Deletes a record from the specified collection.

- Parameters:

  - `collection`: The name of the collection.
  - `resource`: The name of the resource (record).

- Returns:
  - `error`: Any error that occurred during the delete operation.

## Example

Refer to the provided `main` function in the code for a comprehensive example demonstrating the usage of the database API with a collection of users.

## Notes

- The API uses mutexes to handle concurrent access to collections, ensuring data integrity.
- Logging is implemented using the `lumber` package, and logs are output to the console.
