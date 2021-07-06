package Java.se;

// java动态代理代码示例
// 从接口直接创建能运行的实例
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

public class DynamicProxyDemo {
    public static void main(String[] args) {
        InvocationHandler handler = new InvocationHandler(){

            @Override
            public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
                // TODO Auto-generated method stub
                System.out.println(method);
                if (method.getName().equals("morning")) {
                    System.out.println("Good morning, " + args[0]);
                }
                return null;
            }
        };

        // 通过newProxyInstance为接口创建直接能运行的实例，得告诉newProxyInstance需要创建的实例的接口是啥，怎么实例化，以及具体运行什么
        // newProxyInstance的第一个参数：接口的类装载器。类装载器负责从Java字符文件将字符流读入内存，并构造Class类对象
        // 因为java中可以实现多接口，所以第二个参数是个数组。第一个参数传入接口的类加载器
        // 第三个参数来定制接口实例具体干什么事情
        /**
         * handler里的invoke方法，拿到Hello.class中的各项属性，第一个参数proxy表示要代理的interface的具体实例，
         * 第二个参数method表示要代理的这个实例的具体的要代理的方法，第三个参数args表示要传入method的参数。
         * 所以具体的执行过程就是 handler -> invoke -> proxy和method的处理+自己加的处理逻辑
         * 利用动态代理不光可以为接口直接生成实例，也可以在不修改类的代码的同时生成功能更强的实例
         */
        Hello hello = (Hello) Proxy.newProxyInstance(
            Hello.class.getClassLoader(), // 传入ClassLoader
            new Class[] {Hello.class},    // 传入要实现的接口
            handler                       // 传入处理调用方法的InvocationHandler
        );

        hello.morning("Bob");
    }
}

// 上面的动态代理转化为静态代理就是
// public class HelloDynamicProxy implements Hello {
//     InvocationHandler handler;
//     public HelloDynamicProxy(InvocationHandler handler) {
//         this.handler = handler;
//     }
//     public void morning(String name) {
//         handler.invoke(
//            this,
//            Hello.class.getMethod("morning", String.class),
//            new Object[] { name });
//     }
// }

interface Hello {
    void morning(String name);
}