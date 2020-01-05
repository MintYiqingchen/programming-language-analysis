### Style Reflection (Sixteen.go)
Constraints:

- The problem is decomposed using some form of abstraction (procedures, functions, objects, etc.)

- The abstractions have access to information about themselves, although they cannot modify that information


Possible names:

- Introspective
- Navel-gazing

### Style Plugin (Nineteen.go)
Constraints:

- The problem is decomposed using some form of abstraction
  (procedures, functions, objects, etc.)

- All or some of those abstractions are physically encapsulated into
  their own, usually pre-compiled, packages. Main program and each of
  the packages are compiled independently. These packages are loaded
  dynamically by the main program, usually in the beginning (but not
  necessarily).

- Main program uses functions/objects from the dynamically-loaded
  packages, without knowing which exact implementations will be
  used. New implementations can be used without having to adapt or
  recompile the main program.

- External specification of which packages to load. This can be done
  by a configuration file, path conventions, user input or other
  mechanisms for external specification of code to be linked at run
  time.

Possible names:

- No commitment
- Plugins
- Dependency injection
### Run
16.1
```Bash
go build ./Sixteen.go
./Sixteen ../pride-and-prejudice.txt
```

19.1
compile plugins:
```Bash
go build -buildmode=plugin -o plugin/ExtractWords1.so plugin-src/ExtractWords1.go
go build -buildmode=plugin -o plugin/ExtractWords2.so plugin-src/ExtractWords2.go
go build -buildmode=plugin -o plugin/Frequency1.so plugin-src/Frequency1.go
go build -buildmode=plugin -o plugin/Frequency2.so plugin-src/Frequency2.go
```
NineteenConfig.json is the configuration file
compile and run nineteen:
```Bash
go build ./Nineteen.go
./Nineteen ../pride-and-prejudice.txt
```