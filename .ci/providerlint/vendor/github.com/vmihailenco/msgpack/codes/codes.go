package codestype Code bytevar (PosFixedNumHigh Code = 0x7fNegFixedNumLow  Code = 0xe0Nil Code = 0xc0False Code = 0xc2True  Code = 0xc3Float  Code = 0xcaDouble Code = 0xcbUint8  Code = 0xccUint16 Code = 0xcdUint32 Code = 0xceUint64 Code = 0xcfInt8  Code = 0xd0Int16 Code = 0xd1Int32 Code = 0xd2Int64 Code = 0xd3FixedStrLow  Code = 0xa0FixedStrHigh Code = 0xbfFixedStrMask Code = 0x1fStr8         Code = 0xd9Str16        Code = 0xdaStr32        Code = 0xdbBin8  Code = 0xc4Bin16 Code = 0xc5Bin32 Code = 0xc6FixedArrayLow  Code = 0x90FixedArrayHigh Code = 0x9fFixedArrayMask Code = 0xfArray16        Code = 0xdcArray32        Code = 0xddFixedMapLow  Code = 0x80FixedMapHigh Code = 0x8fFixedMapMask Code = 0xfMap16        Code = 0xdeMap32        Code = 0xdfFixExt1  Code = 0xd4FixExt2  Code = 0xd5FixExt4  Code = 0xd6FixExt8  Code = 0xd7FixExt16 Code = 0xd8Ext8     Code = 0xc7Ext16    Code = 0xc8Ext32    Code = 0xc9)
ixedNum(c Code) bool {return c <= PosFixedNumHigh || c >= NegFixedNumLow}
ixedMap(c Code) bool {return c >= FixedMapLow && c <= FixedMapHigh}
ixedArray(c Code) bool {return c >= FixedArrayLow && c <= FixedArrayHigh}
ixedString(c Code) bool {return c >= FixedStrLow && c <= FixedStrHigh}
tring(c Code) bool {return IsFixedString(c) || c == Str8 || c == Str16 || c == Str32}
in(c Code) bool {return c == Bin8 || c == Bin16 || c == Bin32}
ixedExt(c Code) bool {return c >= FixExt1 && c <= FixExt16}
xt(c Code) bool {return IsFixedExt(c) || c == Ext8 || c == Ext16 || c == Ext32}