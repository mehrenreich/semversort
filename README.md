# semversort

Without:

```
$ git tag
v2.10.0
v2.11.0
v2.12.0
v2.12.1
v2.12.2
v2.13.0
v2.14.0
v2.14.1
v2.14.2
v2.15.0
v2.15.1
v2.2.1
v2.2.2
v2.3.0
v2.4.0
v2.4.1
v2.5.0
v2.6.0
v2.6.1
v2.6.2
v2.6.3
v2.6.4
v2.6.5
v2.6.6
v2.6.7rc1
v2.6.8
v2.6.9
v2.7.0
v2.8.0
v2.8.2
v2.8.3
v2.8.4
v2.9.0
```

With:

```
$ git tag | semversort
v2.2.1
v2.2.2
v2.3.0
v2.4.0
v2.4.1
v2.5.0
v2.6.0
v2.6.1
v2.6.2
v2.6.3
v2.6.4
v2.6.5
v2.6.6
v2.6.7rc1
v2.6.8
v2.6.9
v2.7.0
v2.8.0
v2.8.2
v2.8.3
v2.8.4
v2.9.0
v2.10.0
v2.11.0
v2.12.0
v2.12.1
v2.12.2
v2.13.0
v2.14.0
v2.14.1
v2.14.2
v2.15.0
v2.15.1
```
