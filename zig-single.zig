const std = @import("std");
const warn = std.debug.warn;

pub fn main() !void {

//    var timer = try std.os.time.Timer.start();

    var direct_allocator = std.heap.DirectAllocator.init(); 
    defer direct_allocator.deinit();

    var arena = std.heap.ArenaAllocator.init(&direct_allocator.allocator); 
    defer arena.deinit(); 

    const allocator = &arena.allocator; 

    const args = try std.os.argsAlloc(allocator);

    const n = try std.fmt.parseInt(u32, args[1], 10);

    var arr = std.ArrayList(u64).init(allocator);
    try arr.resize(n+1);

    var k1: u32 = 1;
    while (k1 <= @floatToInt(u32, @sqrt(f64, @intToFloat(f64, n)))) {
        var k2: u32 = k1;
        while (k2 <= n / k1) {
            const val = if (k1 != k2) k1+k2 else k1;
            const i = k1*k2;
            arr.set(i, arr.at(i) + val);
            k2 += 1;
        }
        k1 += 1;
    }

    warn("{}\n", arr.at(n));
//    warn("ms: {}", timer.lap() / 1000000);
}
