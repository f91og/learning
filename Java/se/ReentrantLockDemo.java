package Java.se;

import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

// ReentrantLock使用示例
public class ReentrantLockDemo {
    // ......
}

// ReentrantLock是用来代替synchornized这种比较重的锁, 而且可以尝试获取锁。synchornized的话，要么持有要么等待
class Counter{
    private final Lock lock = new ReentrantLock();
    private int count;

    public void add(int n ){
        lock.lock();
        try{
            count += n;
        } finally {
            lock.unlock();
        }
    }

    // 尝试获取锁, 如果等待一段时间后未获得锁，则tryLock返回false，此时程序做一些额外处理，而不是在那里傻等
    public void dec(int n) throws InterruptedException{
        if(lock.tryLock(1, TimeUnit.SECONDS)) {
            try{
                count -= n;
            } finally {
                lock.unlock();
            }
        }
    }

    public int getCount() {
        return count;
    }

}
