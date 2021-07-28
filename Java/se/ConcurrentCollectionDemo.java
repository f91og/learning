package Java.se;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 *  非线程安全的                线程安全的
 * ArrayList	            CopyOnWriteArrayList
 * HashMap	                ConcurrentHashMap
 * HashSet/TreeSet	        CopyOnWriteArraySet
 * ArrayDeque/LinkedList	ArrayBlockingQueue/LinkedBlockingQueue
 * ArrayDeque/LinkedList	LinkedBlockingDeque
 * 使用这些并发集合与使用非线程安全的集合类完全相同
 */
public class ConcurrentCollectionDemo {
    public static void main(String[] args) {
        // ConcurrentHashMap和HashMap的不同：1.线程安全 2.迭代的时候允许修改key
        Map<String, String> map = new ConcurrentHashMap<String, String>();
        
    }
}
