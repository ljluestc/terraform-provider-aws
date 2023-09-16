package msgpcodevar (PosFixedNumHigh byte = 0x7fNegFixedNumLow  byte = 0xe0Nil byte = 0xc0False byte = 0xc2True  byte = 0xc3Float  byte = 0xcaDouble byte = 0xcbUint8  byte = 0xccUint16 byte = 0xcdUint32 byte = 0xceUint64 byte = 0xcfInt8  byte = 0xd0Int16 byte = 0xd1Int32 byte = 0xd2Int64 byte = 0xd3FixedStrLow  byte = 0xa0FixedStrHigh byte = 0xbfFixedStrMask byte = 0x1fStr8         byte = 0xd9Str16        byte = 0xdaStr32        byte = 0xdbBin8  byte = 0xc4Bin16 byte = 0xc5Bin32 byte = 0xc6FixedArrayLow  byte = 0x90FixedArrayHigh byte = 0x9fFixedArrayMask byte = 0xfArray16        byte = 0xdcArray32        byte = 0xddFixedMapLow  byte = 0x80FixedMapHigh byte = 0x8fFixedMapMask byte = 0xfMap16        byte = 0xdeMap32        byte = 0xdfFixExt1  byte = 0xd4FixExt2  byte = 0xd5FixExt4  byte = 0xd6FixExt8  byte = 0xd7FixExt16 byte = 0xd8Ext8     byte = 0xc7Ext16    byte = 0xc8Ext32    byte = 0xc9)
ixedNum(c byte) bool {return c <= PosFixedNumHigh || c >= NegFixedNumLow}
ixedMap(c byte) bool {return c >= FixedMapLow && c <= FixedMapHigh}
ixedArray(c byte) bool {return c >= FixedArrayLow && c <= FixedArrayHigh}
ixedString(c byte) bool {return c >= FixedStrLow && c <= FixedStrHigh}
tring(c byte) bool {return IsFixedString(c) || c == Str8 || c == Str16 || c == Str32}
in(c byte) bool {return c == Bin8 || c == Bin16 || c == Bin32}
ixedExt(c byte) bool {return c >= FixExt1 && c <= FixExt16}
xt(c byte) bool {return IsFixedExt(c) || c == Ext8 || c == Ext16 || c == Ext32}