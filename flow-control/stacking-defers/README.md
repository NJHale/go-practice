# Stacking Defers

Deferred function calls are pushed onto a stack. When a function returns, its
deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this [blog  post](https://blog.golang.org/defer-panic-and-recover).

> **Source:** https://tour.golang.org/flowcontrol/13
