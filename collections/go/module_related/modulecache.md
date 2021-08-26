## Module cache
The module cache is the directory where the go command stores downloaded module files. The module cache is distinct from the build cache, which contains compiled packages and other build artifacts.

The default location of the module cache is $GOPATH/pkg/mod. To use a different location, set the GOMODCACHE environment variable.

The module cache has no maximum size, and the go command does not remove its contents automatically.

The cache may be shared by multiple Go projects developed on the same machine. The go command will use the same cache regardless of the location of the main module. Multiple instances of the go command may safely access the same module cache at the same time.

The go command creates module source files and directories in the cache with read-only permissions to prevent accidental changes to modules after they’re downloaded. This has the unfortunate side-effect of making the cache difficult to delete with commands like rm -rf. The cache may instead be deleted with go clean -modcache. Alternatively, when the -modcacherw flag is used, the go command will create new directories with read-write permissions. This increases the risk of editors, tests, and other programs modifying files in the module cache. The go mod verify command may be used to detect modifications to dependencies of the main module. It scans the extracted contents of each module dependency and confirms they match the expected hash in go.sum.

The table below explains the purpose of most files in the module cache. Some transient files (lock files, temporary directories) are omitted. For each path, $module is a module path, and $version is a version. Paths ending with slashes (/) are directories. Capital letters in module paths and versions are escaped using exclamation points (Azure is escaped as !azure) to avoid conflicts on case-insensitive file systems.