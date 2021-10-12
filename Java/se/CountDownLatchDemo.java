package Java.se;

import java.util.concurrent.CountDownLatch;

// CountDownLatch相当于一个门栓，一开始关闭，然后计时或者计数，为0之后所有等待这个门栓的线程都可以通过。
// 一次性工具，打开后不可以再关闭。 参考：https://www.jianshu.com/p/962bc7225ce8
// 可以实现：1.一个线程等待其他线程的任务完成之后才继续执行自己的任务 2.同时启动多个线程
public class CountDownLatchDemo {
    static class TaskThread extends Thread {
        CountDownLatch latch;

        public TaskThread(CountDownLatch latch) {
            this.latch = latch;
        }

        @Override
        public void run() {
            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            } finally {
                System.out.println(getName() + " Task is Done");
                latch.countDown();
            }
        }
    }

    public static void main(String[] args) throws InterruptedException {
        int threadNum = 10;
        // 构造方法里传入需要等待的线程数量
        CountDownLatch latch = new CountDownLatch(threadNum);
        
        for(int i = 0; i < threadNum; i++) {
            TaskThread task = new TaskThread(latch);
            task.start();
        }
        
        System.out.println("Task Start!");
        // 调用此方法的线程会被阻塞，直到 CountDownLatch 的 count 为 0，这里让main线程一直等待其余线程执行完毕
        latch.await();
        System.out.println("All Task is Done!");
    }
}

