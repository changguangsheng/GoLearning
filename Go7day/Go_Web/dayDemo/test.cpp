#include <iostream>
#include <cstdlib>
#include <ctime>

int main() {
    float input;
    std::cout << "请输入一个浮点数：";
    std::cin >> input;

    // 设置随机种子
    std::srand(std::time(0));

    // 生成随机数
    float num1 = (std::rand() % 100) * input / 100.0f;
    float num2 = input - num1;

    std::cout << "生成的两个数为：" << num1 << " 和 " << num2 << std::endl;

    return 0;
}
