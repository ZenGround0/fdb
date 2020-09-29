Utilities for debugging filecoin
Currently just for decoding strings in the state tree.  Will add more things as needs are discovered.

Big int byte decoding (hex string)
```
> fdb decode int 003dd800000000`
67997922230272
```

Bitfield byte decoding (hex string)
```
> fdb decode bf 20177221
[182329]
```

Base64 string decoding
```
> fdb decode int -b64 AD3YAAAAAA==
67997922230272

> fdb decode bf -b64 IBdyIQ==
[182329]

```