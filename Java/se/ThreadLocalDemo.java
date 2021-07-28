package Java.se;

// Java标准库提供了一个特殊的ThreadLocal，它可以在一个线程中传递同一个对象. 这是因为在一个线程中会存在对某个对象的横跨若干方法的调用
// 实现原理是每个线程都持有实例对象的一个副本，这样可以保证线程之间互不干扰
public class ThreadLocalDemo {
    public static void main(String[] args) {
        
    }


}

// 通过AutoCloseable接口配合try (resource) {...}结构，让编译器自动为我们关闭
class UserContext implements AutoCloseable {

    // ThreadLocal实例通常总是以静态字段初始化如下
    static final ThreadLocal<String> ctx = new ThreadLocal<>();

    public UserContext(String user) {
        // 通过设置一个User实例关联到ThreadLocal中，在移除之前，所有方法都可以随时获取到该User实例
        ctx.set(user);
    }

    public static String currentUser() {
        return ctx.get();
    }

    @Override
    public void close() {
        ctx.remove();
    }
}
