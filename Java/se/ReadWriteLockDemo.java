package Java.se;

import java.util.Arrays;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;

// ReadWriteLock使用示例，不需要互斥时使用，比如允许多个线程同时读的情形
// 保证：只允许一个线程写入（其他线程既不能写入也不能读取），没有写入时，多个线程允许同时读（提高性能）。适合读多写少的场景
public class ReadWriteLockDemo {
    public static void main(String[] args) {
        Counter4 c1 = new Counter4();

        /** 
         * 不要搞混淆了
         * 多线程操作同一个实例的代码时，对于实例的共享代码的线程安全的保证的实现是类自己实现的，不是通过线程来实现
         * 即类的原本逻辑和保证线程安全的逻辑封装在一起，和调用它的线程没啥关系。这样看来线程好像就是个单纯的新线程启动工具而已
         * */ 
        // 写线程
        new Thread(() -> {
            c1.inc(1);
        }).start();
        // 读线程1
        new Thread(() -> {
            System.out.println(c1.get());
        }).start();
        // 读线程2
        new Thread(() -> {
            System.out.println(c1.get());
        }).start();;
    }
}

class Counter4 {
    private final ReadWriteLock rwlock = new ReentrantReadWriteLock();
    private final Lock rlock = rwlock.readLock();   // 获取读锁
    private final Lock wlock = rwlock.writeLock();  // 获得写锁
    private int[] counts = new int[10];

    // 修改的时候加写锁，ReadWriteLock的写锁可以确保一个线程写的时候其它线程不能写也不能读
    public void inc(int index) {
        wlock.lock(); // 加写锁
        try {
            counts[index] += 1;
        } finally {
            wlock.unlock(); // 释放写锁
        }
    }

    // 读取的时候加读锁，ReadWriteLock的读锁可以确保一个线程在读的时候，其他前程只能读，不能写
    public int[] get() {
        rlock.lock(); // 加读锁
        try {
            return Arrays.copyOf(counts, counts.length);
        } finally {
            rlock.unlock(); // 释放读锁
        }
    }
}
