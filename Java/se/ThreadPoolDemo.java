package Java.se;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**      
 * Java标准库提供了ExecutorService接口表示线程池. Executors类中封装了这个接口的三种具体实现
 * FixedThreadPool: 固定大小的线程池
 * CachedThreadPool: 线程数根据任务动态调整的线程池
 * SingleThreadExecutor: 仅单线程执行的线程池
 * ScheduledThreadPool: 需要定期反复执行的任务
 */
public class ThreadPoolDemo {
    public static void main(String[] args) {         
        // 创建固定大小的线程池:
        ExecutorService es = Executors.newFixedThreadPool(3);
        // 先执行线程0,1,2, 结束后线程池有空闲时才执行3,4,5
        for(int i=0; i<6; i++){
            es.submit(new Task("" + i));
        }
        // 关闭线程池
        es.shutdown();


}

class Task implements Runnable{
    private final String name;

    public Task(String name){
        this.name = name;
    }

    @Override
    public void run() {
        System.out.println("start task" + name);
        try{
            Thread.sleep(1000);
        }catch(InterruptedException e){

        }
        System.out.println("end task:" + name);
    }
}


