# models

This package is used to transform a image from a color model to another.

The main difference of this package with the "convert" package is that the
strategies of this package converts from a image.Image to a [][]color.Color,
while the strategies from "convert" converts from [][]color.Color to image.Image

You might wonder why these strategies weren't implemented in "convert" package.
This is to avoid circular imports, as the "configs" package would have to
import "convert" package, which already imports "configs" package, so, it would
be impossible.
