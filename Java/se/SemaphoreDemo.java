package Java.se;

import java.util.concurrent.Semaphore;

// 控制访问特定资源的线程数目，可用于流量控制，限制最大的并发访问数
// 实现多个线程同时执行，但是限制同时执行的线程数量为 2 个。
public class SemaphoreDemo {
    static class TaskThread extends Thread {
        Semaphore semaphore;

        public TaskThread(Semaphore semaphore) {
            this.semaphore = semaphore;
        }

        @Override
        public void run() {
            try {
                // 获取信号量许可，若没有可用则阻塞
                semaphore.acquire();
                System.out.println(getName() + " acquire");
                Thread.sleep(1000);
                // 释放获得的信号量
                semaphore.release();
                System.out.println(getName() + " release ");
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    public static void main(String[] args) {
        int threadNum = 5;
        Semaphore semaphore = new Semaphore(2);
        for (int i = 0; i < threadNum; i++) {
            new TaskThread(semaphore).start();
        }
    }
}
