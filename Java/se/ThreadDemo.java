package Java.se;

// Thread使用实例
public class ThreadDemo {
    public static void main(String[] args) throws InterruptedException {
        // 创建新线程方式1：继承Thread类，覆写run()方法来定义线程
        Thread t1 = new MyThread1();
        // t1.run(); // run()直接执行代码，不会生成新线程去运行

        // 设置守护进程的方式，start前标记为守护进程即可。所有非守护线程都执行完毕后，虚拟机退出
        // 守护线程不能持有任何需要关闭的资源，例如打开文件等，因为虚拟机退出时，守护线程没有任何机会来关闭文件，这会导致数据丢失
        // t.setDaemon(true); 
        t1.start();

        // 创建新线程方式2：创建Thread实例时，传入一个Runnable实例
        Thread t2 = new Thread(new MyRunable());
        t2.start();

        // 利用lambda和匿名函数按照方式2来创建新线程
        Thread t3 = new Thread(() -> {
            System.out.println("start new thread3!");
        });
        t3.start();;

        // 通过join()让一个线程等待另一个线程直到其运行结束。注意是被等待的那个线程调用join()方法
        // join(long)的重载方法也可以指定一个等待时间，超过等待时间后就不再继续等待
        // 同一个Thread不能重复调用start方法, 否则会报java.lang.IllegalThreadStateException
        // System.out.println("start new thread2!");
        // t2.start();
        // t2.join();  // 让main线程等待t2线程运行结束
        // System.out.println("start new thread2!");

        // Thread.setPriority(int n) // 1~10, 默认值5。虽然可以设置线程调度的优先级，但优先级高的线程并不一定就是先执行

        // 通过interrupt()中断线程，和join一样是要中断的线程使用自己的interrupt()。这里的动作方向逻辑有点违背直觉
        // interrupt仅向目标线程发出中断请求，具体实现还是看目标线程中的代码实现
        // 如果下面的代码是先join再interrupt则线程t4永远不会结束，因为t4.join()之后main线程就一直等t4结束，执行不到t4.interrupt()
        Thread t4 = new MyThread4();
        System.out.println("start new thread4");
        t4.start();
        Thread.sleep(1); // 让当前线程暂停1毫秒
        t4.interrupt(); // 终端线程t4
        t4.join();  // 等待t4结束
        System.out.println("tread4 end");

        // 如果main线程中断另一线程t5，此时t5又在等待另一个线程HelloThread结束时，那么t5中的hello.join()则会抛出InterruptedException。
        // mian -中断-> t5 -等待-> hello, 此时t5里面的hello.join和hello线程都会抛出InterruptedException
        Thread t5 = new MyThread5();
        System.out.println("start new thread5");
        t5.start();
        Thread.sleep(1000);
        t5.interrupt();
        t5.join();
        System.out.println("thread5 end");

        // 利用标志位来中断线程，这个变量是线程间共享，所以要保证可见性一致的问题
        HelloThread2 t6 = new HelloThread2();
        t6.start();
        Thread.sleep(1);
        t6.running = false; // main线程刷新了running的值后t6线程能够马上看到
    }

}

// 创建新线程方式1：继承Thread类，覆写run()方法来定义线程
class MyThread1 extends Thread {
    @Override
    public void run(){
        System.out.println("start new thread1!");
    }
}

// 创建新线程方式2：创建Thread实例时，传入一个Runnable实例
class MyRunable implements Runnable {
    @Override
    public void run(){
        System.out.println("start new thread2!");
    }
}

// 目标线程需要反复检测自身状态是否是interrupted状态，如果是，就立刻结束运行
class MyThread4 extends Thread {
    @Override
    public void run() {
        int n = 0;
        while(! isInterrupted()) {  // 这个isInterrupted()是通过继承Thread得到的
            n++;
            System.out.println(n + " hello!");
        }
    }
}

class MyThread5 extends Thread {
    @Override
    public void run() {
        Thread hello = new HelloThread();
        hello.start();
        try {
            hello.join();
        } catch(InterruptedException e){
            System.out.println("interrupted!");
        }
        hello.interrupt();
    }
}

class HelloThread extends Thread {
    @Override
    public void run() {
        int n = 0;
        while(!isInterrupted()){
            n++;
            System.out.println(n + "hello!");
            try {
                Thread.sleep(100);
            } catch (InterruptedException e){
                break;
            }
        }
    }
}

class HelloThread2 extends Thread {
    public volatile boolean running = true;    // 利用volatile保证每个线程看到的变量都是内存中的最新值
    public void run() {
        int n = 0;
        while(running) {
            n ++;
            System.out.println(n + "hello!");
        }
        System.out.println("HelloThread2 end");
    }
}