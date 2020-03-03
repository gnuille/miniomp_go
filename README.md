# miniomp_go
OpenMP like runtime but in go.

* Support for primitive task construct (no dependences, no clauses in general)
* Support for primitive taskloop construct
* Support for taskwait construct

## How do I use this package?
This package uses C preprocessing for conditional compiling features. Right now you can produce to types of builds. A vanila one and an extrae one (go to extrae section for more info).

For running de vanila one you have to download the repo inside the `$GOPATH/src/github.com/teleportex` (see `go env`) and then issue `make install` command for compiling and installing it.

```bash
mkdir -p $GOPATH/src/github.com/teleportex/
git clone https://github.com/teleportex/miniomp_go.git $GOPATH/src/github.com/teleportex/
cd $GOPATH/src/github.com/teleportex/
make install
```
Then you shall be able to import it and initialise the runtime like this:
```go
import "github.com/teleportex/miniomp"

func master() {
   //run tasks
   //do carzy stuff
}

func main(){
  miniomp.Init(master)
}
```
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

## Extrae support (experimental)
Use `make extrae` to run using extrae. To run your own program modify `share/run_extrae.sh` to run your own code!
You will need your own extrae build. You can get yours at [tools.bsc.es](tools.bsc.es).

Right now it supports these events.

| Extrae event code | Description     |
|-------------------|-----------------|
| 8000000           | Task submission |
| 8000001           | Task execution  |
| 8000002           | Taskwait        |
