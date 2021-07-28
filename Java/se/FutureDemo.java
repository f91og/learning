package Java.se;

import java.util.concurrent.Callable;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

/**
 * Future 和 CompletableFuture的用法
 * Runnable接口的问题是没有返回值, 如果任务需要返回一个结果那么只能保存到啥啥可以共享的变量中, 比较麻烦
 * Java标准库还提供了一个Callable接口，和Runnable接口比，它多了一个返回值
 * ExecutorService.submit()返回了一个Future类型，一个Future类型的实例代表一个未来能获取结果的对象
 */
public class FutureDemo {
    public static void main(String[] args) throws InterruptedException, ExecutionException {
        ExecutorService executor = Executors.newFixedThreadPool(4); 
        // 定义任务:
        Callable<String> task = () -> {
            // ...do sth...
            return "Callable task finished";
        };
        // 提交任务并获得Future:
        Future<String> future = executor.submit(task);
        // 从Future获取异步执行返回的结果:
        // 在调用get()时如果异步任务已经完成，就直接获得结果。如果还没有完成，那么get()会阻塞(主线程会因为这个方法阻塞)，直到任务完成后才返回结果
        String result1 = future.get();  
        System.out.println(result1); 

        // 从Java 8开始引入了CompletableFuture，它针对Future做了改进，可以传入回调对象，当异步任务完成或者发生异常时，自动调用回调对象的回调方法
        // 创建异步执行任务, 多个CompletableFuture可以串行执行:
        // 第一个任务:
        CompletableFuture<String> cfQuery = CompletableFuture.supplyAsync(() -> {
            return queryCode("中国石油");
        });
        // 第二个任务，注意这里是从上一个CompletableFuture cfQuery的thenApplyAsync方法来创建新的CompletableFuture，已达到链式执行的目的
        CompletableFuture<Double> cfFetch = cfQuery.thenApplyAsync((code) -> {
            return fetchPrice(code);
        });
        // 如果执行成功:
        cfFetch.thenAccept((result2) -> {
            System.out.println("price: " + result2);
        });
        // 如果执行异常:
        // cfFetch.exceptionally((e) -> {
        //     e.printStackTrace();
        //     return null;
        // });

        // 除了anyOf()可以实现“任意个CompletableFuture只要一个成功”，allOf()可以实现“所有CompletableFuture都必须成功”，这些组合操作可以实现非常复杂的异步流程控制

        // 主线程不要立刻结束，否则CompletableFuture默认使用的线程池会立刻关闭:
        Thread.sleep(200);
    }

    // 获取股票代码
    static String queryCode(String name) {
        try {
            Thread.sleep(100);
        } catch (InterruptedException e) {
        }
        return "601857";
    }

    // 获取股票价格
    static Double fetchPrice(String code) {
        try {
            Thread.sleep(100);
        } catch (InterruptedException e) {
        }
        return 5 + Math.random() * 20;
    }
}

