# Image processing

This project is intended to offer a set of basic features on image processing. It is a tool for image processing.

## Features 🌟

- [X] Flip X
- [X] Flip Y
- [-] Rotate 
  - [ ] 90 degrees
  - [ ] 180 degrees
  - [X] 270 degrees (kind of...)
- [X] Grayscale
- [-] Resize
  - [X] Nearest neighbor (low quality)
  - [ ] Bilinear interpolation
  - [ ] Bicubic interpolation
- [ ] Blur
- [ ] Sharpen
- [ ] Other...

## Dev requirements 🔎

It was developed on golang 1.20.4, but it is likely to work on golang 1.16+ for development. 

## Installing

For the build version you just need to install the [executable](https://github.com/DavidEsdrs/image-processing/releases
).

❗ Are you encountering the "iproc: command not found" error or something similar? This happens because Windows isn't updating the PATH environment variables. You need add a PATH variable with the path that you downloaded the executable. Click [here](https://helpdeskgeek.com/windows-10/add-windows-path-environment-variable/) to learn how to update them manually.

## How to use :books:

After installing the build in your machine. You can perform one or more operations (listed below) onto the images. To perform the operation, you need to run the main package with one or more arguments, separated by space.
The arguments -i (input) and -o (output) are mandatory

```sh
iproc -i [input] -o [output] [...args]
```

> **note**: The effect is applied for all the images in the image folder

Each argument perform a given effect:

- Flip Y:

```sh
iproc -i [input] -o [output] -fy
```

- Flip X:

```sh
iproc -i [input] -o [output] -fx
```

- Rotate (actually, it is a transpose, a rotation 270 degrees + flip in Y axis, it will be changed!!!)

```sh
iproc -i [input] -o [output] -t
```

- Resize (nearest neighbor):

```sh
iproc -i [input] -o [output] -nn [factor]
```

**Note**: The factor of resize must be > 0. Note that the algorithm applied is the `nearest neighbor`, which is known to give pixelated results

Examplo:

```sh
# half of the actual size
iproc -i [input] -o [output] -nn .5
```

- Grayscale:

```sh
iproc -i [input] -o [output] -gs
```

> **note**: If you prefer use the development version, you just need to clone this repository and change `iproc` for `go run main.go`.

> **More will be added soon** 😄

## Considerations ⚠️

As the project progresses, it will get closer to being a tool and more effects will be added.

## Examples ⭐

Apply grayscale filter, flip in Y axis and resize it to half its size

input:
```sh
iproc -i ./images/almoço.png -o ./assets/almoço.png -gs -fy -nn .5
```

before:

![lunch before effects](./images/almoço.png)

after:

![lunch after effects](./assets/almoço.png)