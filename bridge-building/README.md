# Multi-Dimentional Bridge building

## Problem

Given an N dimentional orthotope of edge lengths n_1, n_2, ..., n_N. Assume a bridge builder buids randomly by placing hypercubes within the orthotope. Determine when the bridge creates a continuous path between two opposing sides of a single dimention. Create a function that given the N lengths of the orthotope and a random piece location generator, called the "builder", calls the builder the least number of times before a path is created.

## Example

## 1-D

The 1-D version of this problem can be represented as a straight line along an x axis.
Let us pick a length n_1 = 5:

```
0 1 2 3 4
---x----->
o o o o o
```

In this example, a bridge builder may place their first piece, represented by the charectar 'B':

```
0 1 2 3 4
---x----->
o B o o o
```

The side do not have a continous path through the bridge and so we continue building.
The buiding places their second piece:

```
0 1 2 3 4
---x----->
o B o B o
```

The bridge still does not connect both sides. So we allow the building to continue placing 4 more pieces:

```
0 1 2 3 4
---x----->
B B B B B
```

Now the bridge is complete and we stopped the building process.

## 2-D

The 2-D case is more interesting and can be represented by an x and y axes.
Let us pick lengths n_1 = 4 and n_2 = 5 respectively:

```
    0 1 2 3
    --x---->
0 | o o o o
1 | o o o o
2 y o o o o
3 | o o o o
4 | o o o o
  v
```

The numbers have been placed ourside of the orthotope for simpler representation and the 'o's denote the lack of a bridge piece.
In this example, we must wait until the bridge builder has randomly placed enough pieces to connect the left and right sides.
An end state at which we stop the builder may look like this:

```
    0 1 2 3
    --x---->
0 | o o B o
1 | B B o o
2 y o B B o
3 | o o B B
4 | o o o o
  v
```

```
func BridgeComplete() bool {
  // TODO
}

func main() {

  o := 
}
```
