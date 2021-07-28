package Java.se;

// 线程同步：synchronized加锁使用实例
public class SynchronizedDemo {
    public static void main(String[] args) throws InterruptedException {
        // 线程不安全实例
        var add = new AddThread();
        var dec = new DecThread();
        add.start();
        dec.start();
        add.join();
        dec.join();
        // 不要以为add是递增了1000次，dev是递减了1000次，结果就一定是0. 因为递增和递减不是原子操作，不能保证递增的时候就一定加了1，递减的时候也是同理
        System.out.println(Counter.count); 

        // 利用synchronized加锁来确保线程安全，synchronized加的锁是可重入锁
        var add2 = new AddThread2();
        var dec2 = new DecThread2();
        add2.start();
        dec2.start();
        add2.join();
        dec2.join();
        System.out.println(Counter2.count);

        // synchronized(this)使用示例，用类的当前实例作为锁对象，这样两个线程操作同一对象时，后来的就要等先来的释放锁
        var c1 = new Counter3();    // Counter3类的多个实例，可以并发运行，这种类被称为“线程安全的”
        var c2 = new Counter3();
        new Thread(() -> {
            c1.add(1);
        }).start();;
        new Thread(() -> {
            c1.dec(1);
        }).start();
        new Thread(() -> {
            c2.add(1);
        }).start();;
        new Thread(() -> {
            c2.dec(1);
        }).start();;

    }
}

class Counter {
    public static int count = 0;    
}

class AddThread extends Thread{
    public void run() {
        for (int i=0;i<1000;i++) {Counter.count += 1;}
    }
}

class DecThread extends Thread{
    public void run() {
        for (int i=0;i<1000;i++) {Counter.count -= 1;}
    }
}

class Counter2 {
    public static final Object lock = new Object();
    public static int count = 0;
}


// synchronized(lockObject){}，选择一个共享实例作为锁。
class AddThread2 extends Thread{
    public void run() {
        for (int i=0;i<1000;i++) {
            synchronized(Counter2.lock) {
                Counter2.count += 1;
            }
        }
    }
}

class DecThread2 extends Thread{
    public void run() {
        for (int i=0;i<1000;i++) {
            synchronized(Counter2.lock) {
                Counter2.count -= 1;
            }
        }
    }
}

// 把synchronized逻辑封装起来，而不是让线程自己选择锁对象
// 对于java这种通过类来操纵类的模型，改变另个一个类中的成员状态要么在自己类中操作另一个类的成员，要么将对类的成员操作和成员一起封装
class Counter3 {
    private int count = 0;

    public void add(int n) {
        synchronized(this){
            count += n;
        }
    }

    // 用synchronized修饰的方法就是同步方法，它表示整个方法都必须用this实例加锁。下面的add和上面的等价
    // public synchronized void add(int n) { // 锁住this
    //     count += n;
    // } // 解锁

    public void dec(int n){
        synchronized(this){
            count -= n;
        }
    }

    public int get() {
        return count;
    }
}

// 对类的静态方法使用synchronized相当于使用类的class（即Counter.class）作为锁对象




