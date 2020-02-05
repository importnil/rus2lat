# rus2lat
Go package which can efficiently transliterate Russian Cyrillic characters to Latin.

rus2lat strictly conforms Yandex transliteration rules. Very fast and accurate.

## Installation & Update
```
go get -u github.com/vadimyer/rus2lat
```

## Usage

There are only 2 methods: one is for raw text and another is for an URL transformation. Both return string values.

```
rus2lat.Raw("Привет, я Вася!") // Privet, ya Vasya!
rus2lat.URL("Привет, я Вася!") // privet-ya-vasya
```

## Benchmarks

##### System specs
- Intel Core i7-6700 (8x4,00GHz)
- 16 GB RAM (4x4 DDR4-2133, quad channel)
- Go 1.6 windows/amd64
- Windows 10 Pro.

##### Results
```
BenchmarkRaw-8   1000000              1451 ns/op             148 B/op          6 allocs/op
BenchmarkURL-8   1000000              1175 ns/op             128 B/op          2 allocs/op
```

Text used: Привет, я Вася!

## Contribution
If you find an issue or just have a proposal, feel free to post it. Pull requests are also welcome.
