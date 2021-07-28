package Java.se;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * CAS无锁方式的方式实现了对基本数据的一些简单操作的原子化。
 * 线程安全或者线程不安全的AtomicDemo类，这样的类至少有些类的成员变量，实例化其一个实例后，所有线程对这个实例的访问，都是在操作同一份变量，因此需要确保线程安全性
 * 下面的例子展示了一个线程安全的计数和两个线程非安全的计数
 */
public class AtomicDemo {
    // 内部持有3个状态
    private AtomicInteger atomicInteger = new AtomicInteger(0);
    private int i = 0;
    private int unsafe = 0;

    public static void main(String[] args) {
        final AtomicDemo cas = new AtomicDemo();
        final int TOTAL_THREAD_NUM = 10;
        List<Thread> ts = new ArrayList<Thread>();

        long start = System.currentTimeMillis();

        // 实例化TOTAL_THREAD_NUM个线程，每个线程运行时对AtomicDemo的一个实例cas中的三个成员变量atomicInteger，i，unsafe进行递增
        for (int j = 0; j < TOTAL_THREAD_NUM; j++) {
            Thread t = new Thread(new Runnable() {
                @Override
                public void run() {
                    Thread.currentThread().setName(AtomicDemo.class.getName() + Thread.currentThread().getId());
                    for (int i = 0; i < 1000; i++) {
                        cas.count();
                        cas.safeCount();
                        cas.unsafeCount();
                    }
                    System.out.println("Thread name: " + Thread.currentThread().getName());
                }
            });
            ts.add(t);
        }

        System.out.println("total thread number:" + ts.size());

        for (Thread t : ts) {
            t.start();
        }

        // 顺序等待所有线程执行完成
        for (Thread t : ts) {
            try {
                t.join();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println("safe via synchronized: " + cas.i);
        System.out.println("safe: " + cas.atomicInteger.get());
        System.out.println("unsafe: " + cas.unsafe);
        System.out.println(System.currentTimeMillis() - start);
    }

    // 线程安全的计数
    private void safeCount() {
        atomicInteger.incrementAndGet();
        // for (;;) {
        //     int i = atomicInteger.get();
        //     boolean suc = atomicInteger.compareAndSet(i, ++i);
        //     if (suc) {
        //         break; // 自旋，直到成功才退出
        //     }
        // }
    }

    // synchronized实现的安全计数器
    synchronized private void count() {
        i++;
    }

    // 非线程安全的计数器
    private void unsafeCount() {
        unsafe++;
    }
}
