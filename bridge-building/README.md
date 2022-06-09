# Multi-Dimentional Bridge building

## Problem

Given an N dimentional orthotope of edge lengths n<sub>1</sub>, n<sub>2</sub>, ..., n<sub>N</sub>. Assume a bridge builder buids randomly by placing hypercubes within the orthotope. Determine when a continuous path exists along the 1st dimention from 0 to n<sub>1</sub>-1.

## Example

### 1D

The 1D version of this problem can be represented as a straight line along an x axis.
Let us pick a length n<sub>1</sub> = 5 where the 'o's denote the lack of a bridge piece.:

```
0 1 2 3 4
---x----->
o o o o o
```

A bridge builder may place their first piece, represented by the charectar 'B', randomly at 1:

```
0 1 2 3 4
---x----->
o B o o o
```

The sides do not have a continous path through the bridge and so we continue building.
The buider places their second piece at 3:

```
0 1 2 3 4
---x----->
o B o B o
```

The bridge still does not connect both sides. So we allow the builder to continue placing 3 more pieces:

```
0 1 2 3 4
---x----->
B B B B B
```

Now the bridge is complete and we stop the building process.

### 2-D

The 2-D case is more interesting and can be represented by an x and y axis.
Let us pick lengths n<sub>1</sub> = 4 and n<sub>2</sub> = 5 respectively:

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

## Helper Structs and Functions

You are given some starting code to work with. It includes:

- An `Orthotope` object struct that represents the space the brindge pieaces can be built in.
- A `Orthotope.Built(loc..)` function that returns whether or not there is a bridge piece at loc.

## Requirment

Implement the `Orthotope.BridgeComplete` function such that it passes the existing test cases.
