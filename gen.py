def T(l, v, r):
    return {'left': l, 'value': v, 'right': r}


def g():
    print('g is called')
    h()


def h():
    print('h is called')
    # ここにyieldを書くとprintは呼ばれなくなる


def visit(t):
    if t is not None:
        yield from visit(t['left'])
        g()
        yield t['value']
        yield from visit(t['right'])


e = None
t1 = T(T(T(e, 1, e), 2, T(e, 3, e)), 4, T(e, 5, e))
for x in visit(t1):
    print(x)
