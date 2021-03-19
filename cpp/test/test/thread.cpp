#include <atomic>
#include <chrono>
#include <condition_variable>
#include <future>
#include <iostream>
#include <mutex>
#include <queue>
#include <thread>
#include <filesystem>

using namespace std;

void test1()
{
    std::thread t([]() {
        std::cout << "hello world." << std::endl;
    });
    t.join();
}

// ------
//int v = 1;
//void critical_section(int change_v)
//{
//    static std::mutex mtx;
//    std::lock_guard<std::mutex> lock(mtx);
//    // 执行竞争操作
//    v = change_v;
//    // 离开此作用域后 mtx 会被释放
//}
//void test2()
//{
//    for (int i = 1; i < 1000; i++)
//    {
//        std::thread tt(critical_section, i);
//        tt.join();
//    }
//    std::cout << v << std::endl;
//}

// -----------------------
int v = 1;
void critical_section(int change_v)
{
    static std::mutex mtx;
    std::unique_lock<std::mutex> lock(mtx);
    // 执行竞争操作
    v = change_v;
    std::cout << v << std::endl;
    // 将锁进行释放
    lock.unlock();
    // 在此期间，任何人都可以抢夺 v 的持有权
    // 开始另一组竞争操作，再次加锁
    lock.lock();
    v += 1;
    std::cout << v << std::endl;
}
void test3()
{
    std::thread t1(critical_section, 2), t2(critical_section, 3);
    t1.join();
    t2.join();
}

// ---------------------------------------
void test4()
{
    // 将一个返回值为7的 lambda 表达式封装到 task 中
    // std::packaged_task 的模板参数为要封装函数的类型
    std::packaged_task<int()> task([]() { return 7; });
    // 获得 task 的期物
    std::future<int> result = task.get_future(); // 在一个线程中执行 task
    std::thread(std::move(task)).detach();
    std::cout << "waiting...";
    result.wait(); // 在此设置屏障，阻塞到期物的完成
    // 输出执行结果
    std::cout << "done!" << std::endl
              << "future result is " << result.get() << std::endl;
}

// ----------------------------------------------
int test5()
{
    std::queue<int> produced_nums;
    std::mutex mtx;
    std::condition_variable cv;
    bool notified = false; // 通知信号
    // 生产者
    auto producer = [&]() {
        for (int i = 0;; i++)
        {
            std::this_thread::sleep_for(std::chrono::milliseconds(900));
            std::unique_lock<std::mutex> lock(mtx);
            std::cout << "producing " << i << std::endl;
            produced_nums.push(i);
            notified = true;
            cv.notify_all(); // 此处也可以使用 notify_one
        }
    };
    // 消费者
    auto consumer = [&]() {
        while (true)
        {
            std::unique_lock<std::mutex> lock(mtx);
            while (!notified)
            { // 避免虚假唤醒
                cv.wait(lock);
            }
            // 短暂取消锁，使得生产者有机会在消费者消费空前继续生产
            lock.unlock();
            std::this_thread::sleep_for(std::chrono::milliseconds(1000)); // 消费者慢于生产者
            lock.lock();
            while (!produced_nums.empty())
            {
                std::cout << "consuming " << produced_nums.front() << std::endl;
                produced_nums.pop();
            }
            notified = false;
        }
    };
    // 分别在不同的线程中运行
    std::thread p(producer);
    std::thread cs[2];
    for (int i = 0; i < 2; ++i)
    {
        cs[i] = std::thread(consumer);
    }
    p.join();
    for (int i = 0; i < 2; ++i)
    {
        cs[i].join();
    }
    return 0;
}

// --------------------------------------------
void test6()
{
    int a = 0;
    int flag = 0;
    std::thread t1([&]() {
        while (flag != 1)
            ;
        int b = a;
        std::cout << "b = " << b << std::endl;
    });
    std::thread t2([&]() {
        a = 5;
        flag = 1;
    });
    t1.join();
    t2.join();
}

// atomic
std::atomic<int> count = {0};
void test7()
{
    std::thread t1([]() {
        count.fetch_add(1);
    });
    std::thread t2([]() {
        count++;    // 等价于 fetch_add
        count += 1; // 等价于 fetch_add
    });
    t1.join();
    t2.join();
    std::cout << count << std::endl;
}

int main(int argc, char **argv)
{
    //test1();
    //test2();
    //test3();
    //test4();
    //test5();
    //test6();
    test7();

    return 0;
}