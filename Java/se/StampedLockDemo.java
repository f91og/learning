package Java.se;

import java.util.concurrent.locks.StampedLock;

// ReadWriteLock是悲观锁, 即读的时候不允许写.
// 乐观锁StampedLock使用示例, 读的时候允许写锁后入, 读的结果可能不一致, 需要加额外代码判断
// StampedLock是不可重入锁
public class StampedLockDemo {
    public static void main(String[] args) {
        Point p = new Point();
        new Thread(() -> {
            p.distanceFromOrigin();
        }).start();
        new Thread(() -> {
            p.move(1, 1);
        }).start();
    }
}

class Point{
    private final StampedLock stampedLock = new StampedLock();

    private double x;
    private double y;

    public void move(double deltaX, double deltaY) {
        long stamp = stampedLock.writeLock(); // 获取写锁
        try {
            x += deltaX;
            y += deltaY;
        } finally {
            stampedLock.unlockWrite(stamp); // 释放写锁
        }
    }

    public double distanceFromOrigin() {
        long stamp = stampedLock.tryOptimisticRead(); // 获得一个乐观读锁, 并返回一个版本号.乐观锁实际上并不会加锁
        // 注意下面两行代码不是原子操作
        // 假设x,y = (100,200)
        double currentX = x;
        // 此处已读取到x=100，但x,y可能被写线程修改为(300,400)
        double currentY = y;
        // 此处已读取到y，如果没有写入，读取是正确的(100,200)
        // 如果有写入，读取是错误的(100,400)
        // 检查乐观读锁后是否有其他写锁发生, validate()去验证版本号有没有在读期间发生改变. 如果有则采取悲观锁策略. 
        if (!stampedLock.validate(stamp)) { 
            stamp = stampedLock.readLock(); // 获取一个悲观读锁
            try {
                currentX = x;
                currentY = y;
            } finally {
                stampedLock.unlockRead(stamp); // 释放悲观读锁
            }
        }
        return Math.sqrt(currentX * currentX + currentY * currentY);
    }
}

