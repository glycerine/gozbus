@0xdd278ee77d74da85;

# zbus.capnp : describe times, dates, and basic types
# in capnp schema. capnp can generate bindings for
# golang, python, C++, etc.

using Cxx = import "c++.capnp";
$Cxx.namespace("zbus");

using Go = import "go.capnp"; 
$Go.package("gozbus");
$Go.import("github.com/glycerine/gozbus");


struct Zdate {
  # pack date into 64-bits total

  year  @0   :Int16;
  month @1   :UInt8;
  day   @2   :UInt8;

  # TJD or tjd is the Truncated Julian Day convention
  #   for numbering days sequentially.
  #    (see http://en.wikipedia.org/wiki/Julian_day ) 
  #   TJD was introduced by NASA in 1979. Its epoch 
  #   is 0h May 24, 1968. TJD counts days sequentially 
  #   since that epoch. The longer Julian date requires
  #   many extra under-utilized bits; it is the days since 
  #   12h Jan 1, 4713 BC.  Their relationship: 
  #      TJD == JD − 2440000.5, or
  #      JD  == TJD + 2440000.5; the 0.5 occurs since
  #      the JD counts from noon instead of midnight.
  #
  # TJD makes for convenient comparison and compuation,
  #  and in addition we adapt it to the local timezone,
  #  instead of UTC Greenwich time.
  #
  tjday @3  :UInt32;
}

struct Ztm {
       # milliseconds since midnight of today (whatever today happens to be)
       tmMsecMidnt @0 : UInt32;

       # optional nanosecond precision: nanoseconds since the last second boundary.
       tmNanoSinceLastSec   @1 : UInt32;
}

struct Ztd {
  date @0 : Zdate;
  time @1 : Ztm;
}


struct Z {
  # Z must contain all types, as this is our
  # runtime type identification. It is a thin shim.

  union {
    void              @0: Void; # always first in any union.
    zz                @1: Z;    # any. fyi, this can't be 'z' alone.

    f64               @2: Float64;
    f32               @3: Float32;

    i64               @4: Int64;
    i32               @5: Int32;
    i16               @6: Int16;
    i8                @7: Int8;

    u64               @8:  UInt64;
    u32               @9:  UInt32;
    u16               @10: UInt16;
    u8                @11: UInt8;

    bool              @12: Bool;
    text              @13: Text;
    blob              @14: Data;

    f64vec            @15: List(Float64);
    f32vec            @16: List(Float32);

    i64vec            @17: List(Int64);
    i32vec            @18: List(Int32);
    i16vec            @19: List(Int16);
    i8vec             @20: List(Int8);

    u64vec            @21: List(UInt64);
    u32vec            @22: List(UInt32);
    u16vec            @23: List(UInt16);
    u8vec             @24: List(UInt8);

    zvec              @25: List(Z);
    zvecvec           @26: List(List(Z));

    # time and date
    zdate             @27: Zdate;
    ztm               @28: Ztm;
    ztd               @29: Ztd; # date and time together

  }
}
