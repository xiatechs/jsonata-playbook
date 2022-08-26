# jsonata playbook for go-jsonata 1.5.4

```
to run the web-app locally just run 'make start' to boot up docker container
and then go to https://127.0.0.1:8050 on browser
```

[official documentation](https://docs.jsonata.org/overview.html)

[open source library](https://github.com/blues/jsonata-go)

![](jsonatalogo.PNG) 

this is a series of practical examples of go-jsonata. 
the official documentation is helpful, but some of us 
want to see it being used with 'real' examples. 

the purpose of this playbook is to provide a series of real 
examples to demonstrate the power of jsonata in event based
data transformation work.

this will contain a series of examples of jsonata
for you to refer to when you start building connectors
using the more advanced top-level jsonata language.

jsonata will be a key aspect of ensuring that data is 
normalised & transformed either for key/value storage
or for emitting data to flink or endpoints.

1) [variables](variables.md)

2) [json navigation - non array](navigation-nonarray.md)

3) [json navigation - array](navigation-array.md)

4) [conditionals](conditionals.md)

5) [functions](functions.md)

6) [map](map.md)

7) [filter](filter.md)

8) [reduce](reduce.md)

9) [composition](composition.md)

10) [errors](errors.md)

11) [transformation](transformation.md)

12) [state](state.md)

13) [time](time.md)

14) [stateless watermarking](watermark.md)

if you want to add more examples please feel free to raise a PR
- Tom
