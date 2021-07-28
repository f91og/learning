package Java.se;

// Java的Object类包含了三个final方法，允许线程就资源的锁定状态进行通信。这三个方法分别是：wait(),notify(),notifyAll()
public class WaitAndNotifyDemo {
    public static void main(String[] args) throws InterruptedException {
        WaitAndNotify waitAndNotify = new WaitAndNotify();

        // 对象::实例名的方式
        Runnable taskWait = waitAndNotify::threadWait;
        new Thread(taskWait).start();
    
        Thread.sleep(3000);
    
        // 这里的一个疑问是即使另一个线程使用threadNotify方法，但是此时它拿不到this锁，应该执行不了才对
        /**
         * 关键就在于wait()方法的执行机制非常复杂。首先，它不是一个普通的Java方法，而是定义在Object类的一个native方法，也就是由JVM的C代码实现的。
         * 其次，必须在synchronized块中才能调用wait()方法，因为wait()方法调用时，会释放线程获得的锁，wait()方法返回后，线程又会重新试图获得锁
         */
        Runnable taskNotify = waitAndNotify::threadNotify;
        new Thread(taskNotify).start();   
    }
}

// 多线程安全问题指的是同时多个线程操作同一段代码，如果是不同代码那压根没有啥线程安全问题
// 所以多个线程的协作是多个线程操作同一段代码时，如何让这些线程相互通信
class WaitAndNotify{

    // wait()方法的作用是将当前运行的线程挂起（即让其进入阻塞状态），直到notify或notifyAll方法来唤醒线程
    // 调用wait的方法的当前线程应具有对象监视器，要用synchronized加以修饰，否则会抛出IllegalMonitorStateException
    synchronized void threadWait() {
        System.out.println("wait");
        try{
            // 释放this锁
            this.wait();
            // 等待结束后重新获得this锁，已唤醒的线程还需要重新获得锁后才能继续执行
        } catch (InterruptedException e){
            e.printStackTrace();
        }
        System.out.println("unlock");
    }

    // notify()を呼び出すとそのオブジェクトのウェイトセットにあるスレッドが1つ再開します。
    // 複数のスレッドが待機している場合は、そのうちのどれか1つが再開します
    synchronized void threadNotify() {
        this.notify();  // 只唤醒一个等待的线程，有一定的随机性
        // this.notifyAll();    // 将等待的线程全部唤醒
        System.out.println("notified!");
    }
}