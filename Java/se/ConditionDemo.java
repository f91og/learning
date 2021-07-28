package Java.se;

import java.util.LinkedList;
import java.util.Queue;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

// synchronized可以配合wait和notify实现线程在条件不满足时等待，条件满足时唤醒。
// 使用ReentrantLock + Condition实现同样的功能
public class ConditionDemo{
    public static void main(String[] args) {
        
    }
}

class TaskQueue{
    private final Lock lock = new ReentrantLock();
    // 引用的Condition对象必须从Lock实例的newCondition()返回，这样才能获得一个绑定了Lock实例的Condition实例
    private final Condition condition = lock.newCondition();
    private Queue<String> queue = new LinkedList<>();

    public void addTask(String s) {
        lock.lock();
        try{
            queue.add(s);
            // 利用condition的signalAll方法来唤醒所有wait的线程，作用等同于notifyAll
            condition.signalAll();
        }finally{
            lock.unlock();
        }
    }

    // 
    public String getTask() throws InterruptedException {
        lock.lock();
        try{
            while (queue.isEmpty()) {
                // 等同于wait()
                condition.await();
            }
            return queue.remove();
        }finally{
            lock.unlock();
        }
    }
    
}