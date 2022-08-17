const std = @import("std");
const count = 100000000;

pub fn main() anyerror!void {
    var n:i64 = 1;
    while(n <= count) {
      var result: bool = isPrimeZig(n);
      if (result) {
        std.log.info("{} is prime", .{n});
      } else {
        std.log.info("{} is not prime", .{n});
      }
      n += 1;
    }
}

pub fn isPrimeZig(num: i64) bool {
  if (num == 1) return false;
  if (num == 2) return true;
  if (@rem(num, 2) == 0) return false;
  var root: f64 = @sqrt(@intToFloat(f64, num));
  var end: i64 = @floatToInt(i64, (root + 1));
  var p: i64 = 3;
  while(p <= end) {
    if (@rem(num, p) == 0) return false;
    p += 2;
  }
  return true;
}
