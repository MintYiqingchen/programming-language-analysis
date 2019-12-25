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