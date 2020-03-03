# miniomp_go
OpenMP like runtime but in go.

* Support for primitive task construct (no dependences, no clauses in general)
* Support for primitive taskloop construct
* Support for taskwait construct

## How do I run a task?

You have to encapsulate your task in a function, in the example named function, and your arguments in an variable, in the example arguments.

```go
miniomp.NewTask( function, arguments )
```

The function has to be declared with the following interface:
```go
func function(arguments []interface{}) interface{} {
 //accesing to parameters
  a := arguments[0].(int)
  b := arguments[1].(string)
  //computing
  fmt.Println(b+" "+a)
  return nil
}
```
And the arguments that the task will use have to be declared with the following interface:
```go
arguments := []interface{}{ 42, "YAY"}
```

## Extrae support
Use `make extrae` to run using extrae. To run your own program modify `share/run_extrae.sh` to run your own code!

| Extrae event code | Description     |
|-------------------|-----------------|
| 8000000           | Task submission |
| 8000001           | Task execution  |
| 8000002           | Taskwait        |
