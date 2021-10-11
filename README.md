# world.go

```
⣿⣿⣿⣿⣿⣿⣿⡿⡛⠟⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⠿⠨⡀⠄⠄⡘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⠿⢁⠼⠊⣱⡃⠄⠈⠹⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⡿⠛⡧⠁⡴⣦⣔⣶⣄⢠⠄⠄⠹⣿⣿⣿⣿⣿⣿⣿⣤⠭⠏⠙⢿⣿⣿⣿⣿⣿
⣿⡧⠠⠠⢠⣾⣾⣟⠝⠉⠉⠻⡒⡂⠄⠙⠻⣿⣿⣿⣿⣿⡪⠘⠄⠉⡄⢹⣿⣿⣿⣿
⣿⠃⠁⢐⣷⠉⠿⠐⠑⠠⠠⠄⣈⣿⣄⣱⣠⢻⣿⣿⣿⣿⣯⠷⠈⠉⢀⣾⣿⣿⣿⣿
⣿⣴⠤⣬⣭⣴⠂⠇⡔⠚⠍⠄⠄⠁⠘⢿⣷⢈⣿⣿⣿⣿⡧⠂⣠⠄⠸⡜⡿⣿⣿⣿
⣿⣇⠄⡙⣿⣷⣭⣷⠃⣠⠄⠄⡄⠄⠄⠄⢻⣿⣿⣿⣿⣿⣧⣁⣿⡄⠼⡿⣦⣬⣰⣿
⣿⣷⣥⣴⣿⣿⣿⣿⠷⠲⠄⢠⠄⡆⠄⠄⠄⡨⢿⣿⣿⣿⣿⣿⣎⠐⠄⠈⣙⣩⣿⣿
⣿⣿⣿⣿⣿⣿⢟⠕⠁⠈⢠⢃⢸⣿⣿⣶⡘⠑⠄⠸⣿⣿⣿⣿⣿⣦⡀⡉⢿⣧⣿⣿
⣿⣿⣿⣿⡿⠋⠄⠄⢀⠄⠐⢩⣿⣿⣿⣿⣦⡀⠄⠄⠉⠿⣿⣿⣿⣿⣿⣷⣨⣿⣿⣿
⣿⣿⣿⡟⠄⠄⠄⠄⠄⠋⢀⣼⣿⣿⣿⣿⣿⣿⣿⣶⣦⣀⢟⣻⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⡆⠆⠄⠠⡀⡀⠄⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⡿⡅⠄⠄⢀⡰⠂⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
```

Experimental CGo-based wrapper for the [WORLD] vocoder.

**NOTE**: Some functions may have not been implemented yet (refer to the table below):

| pyworld                  | world.go   |
|--------------------------|------------|
| dio                      | Dio        |
| harvest                  | Harvest    |
| stonemask                | StoneMask  |
| get_cheaptrick_fft_size  | -          |
| get_cheaptrick_f0_floor  | -          |
| cheaptrick               | CheapTrick |
| d4c                      | D4C        |
| synthesize               | Synthesize |
| get_num_aperiodicities   | -          |
| code_aperiodicity        | -          |
| decode_aperiodicity      | -          |
| code_spectral_envelope   | -          |
| decode_spectral_envelope | -          |
| wav2world                | WavToWorld |

# Usage

The API of world.go is similar to that of [pyworld]:

```go
import "https://github.com/ongyx/world.go"

// where x is the wavfile and fs is the sample rate.
var x []float64
var fs int

f0, sp, ap := world.WavToWorld(x, fs)
```

`sp` and `ap` are represented by a custom `Matrix` type.
To access the underlying `[][]float64` slice, use `mat.S`.

Almost all algorithms accept a `Options` struct in their args.
Most of the time, you can use the defaults:

```go
f0, tpos := world.Dio(x, fs, world.DefaultOptions)
```

But if you need to customize the defaults, just copy them:

```go
o := *world.DefaultOptions

// change defaults...

f0, tpos := world.Dio(x, fs, &o)
```

# Building/Installing

Clone this repo and install:

```
git clone --recurse-submodules https://github.com/ongyx/world.go
go install
```

For Linux and Windows, the `libworld.a` static library has been precompiled under the `lib/` folder.

If you want to compile `libworld.a` from source, do `make cross-compile` which will compile both the Linux and Windows versions.
This assumes that you are using a Linux system with both the native gcc compiler and the MSYS2 cross-compiler.

# Todo

- documentation
- tests

[WORLD]: https://github.com/mmorise/World
[pyworld]: https://github.com/JeremyCCHsu/Python-Wrapper-for-World-Vocoder
