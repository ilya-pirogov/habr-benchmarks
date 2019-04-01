#include <vector>
#include <cmath>
#include <iostream>
#include <thread>
#include <atomic>

int main(int argc, char ** argv) {
  if(argc < 2) {
    std::cerr << "Usage: " << argv[0] << " maxN" << std::endl;
    return 1;
  }

  int64_t n = std::stoi(argv[1]);
  int64_t nthreads = std::stoi(argv[2]);

  std::vector<std::atomic<int64_t>> arr(n + 1);

  auto work = [&](size_t nthreads, size_t tnum, int64_t n) {
    for(int64_t k1 = tnum; k1 <= int64_t(std::sqrt(n)); k1 += nthreads) {
      for(int64_t k2 = k1; k2 <= n / k1; ++k2)
        arr[k1 * k2] += (k1 != k2 ? k1 + k2 : k1);
    }
  };

  auto run = [&](size_t nthreads) {
    std::vector<std::thread> out;
    for(size_t tnum = 1; tnum <= nthreads; ++tnum)
      out.emplace_back(work, nthreads, tnum, n);
    return out;
  };

  for(auto & x: run(nthreads)) x.join();

  std::cout << arr[n] << std::endl;
}
