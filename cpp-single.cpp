#include <vector>
#include <string>
#include <cmath>
#include <iostream>

int main(int argc, char **argv)
{
    if (argc != 2)
    {
        std::cerr << "Usage: " << argv[0] << " maxN" << std::endl;
        return 1;
    }
    int64_t n = std::stoi(argv[1]);

    std::vector<int64_t> arr;
    arr.resize(n + 1);

    for (int64_t k1 = 1; k1 <= static_cast<int64_t>(std::sqrt(n)); ++k1)
    {
        for (int64_t k2 = k1; k2 <= n / k1; ++k2)
        {
            auto val = k1 != k2 ? k1 + k2 : k1;
            arr[k1 * k2] += val;
        }
    }

    std::cout << arr.back() << std::endl;
}
