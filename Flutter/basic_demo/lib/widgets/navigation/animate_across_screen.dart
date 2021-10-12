// guide users through an app as they navigate from screen to screen
// 通过一个窗口到另一个窗口的切换来引导用户对app的使用，有点类似才打开app的时候一张张图片切换的教学
// 这个例子在网络图片显示那块有点问题
// 尝试去掉Hero widget后的显示效果是啥样的
import 'package:flutter/material.dart';

class HeroApp extends StatelessWidget {
  const HeroApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      title: 'Transition Demo',
      home: MainScreen(),
    );
  }
}

// 1. Create two screens showing the same image
class MainScreen extends StatelessWidget {
  const MainScreen({Key? key}) : super(key: key);
  
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Main Screen'),
      ),
      body: GestureDetector(
        onTap: () {
          Navigator.push(context, MaterialPageRoute(builder: (context) {
            return const DetailScreen();
          }));
        },
        child: Hero(
          tag: 'imageHero',
          child: Image.network(
            'https://picsum.photos/250?image=9',
          ),
        ), 
      ),
    );
  }
}

class DetailScreen extends StatelessWidget {
  // 如果没有这里的const定义则上面在Navigator代码块中使用的时候则无法在DetailScreen()前面加const关键字
  // 这么做的原因是什么❓需要搞清楚里面的机制 
  const DetailScreen({Key? key}) : super(key: key); 

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: GestureDetector(
        onTap: () {
          Navigator.pop(context);
        },
        child: Center(
          child: Hero(
            tag: 'imageHero',
            child: Image.network(
              'https://picsum.photos/250?image=9',
            ),
          ),
        ),
      ),
    );
  } 
}

// 2. Add a Hero widget to the first screen