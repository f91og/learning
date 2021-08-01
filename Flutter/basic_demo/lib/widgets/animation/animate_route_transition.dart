import 'package:flutter/material.dart';

// Animate a page route transition
// 1. Set up a PageRouteBuilder, 设置带过度动画的路由
// 2. Create a Tween，创造一个过度效果
// 3. Use an AnimatedWidget，利用animation.drive(过度效果)来动画演示
// 4. Use a CurveTween，使用CurveTween变化效果

// 下面这个例子演示了新页面从低到上的弹入和从上到低的弹出的动画效果, 并使用了Curves.ease这个动画效果
class AnimateRouteTransition extends StatelessWidget {
  const AnimateRouteTransition({ Key? key }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Center(
        child: ElevatedButton(
          onPressed: (){
            Navigator.of(context).push(_createRoute());
          },
          child: const Text("Go"),
        ),
      ),
    );
  }
}

//  PageRouteBuilder has two callbacks, one to build the content of the route (pageBuilder), and one to build the route’s transition (transitionsBuilder)
Route _createRoute(){
  return PageRouteBuilder(
    pageBuilder: (context, animation, secondaryAnimation) => const Page2(),
    transitionsBuilder: (context, animation, secondaryAnimation, child) {
      const begin = Offset(0.0, 1.0);
      const end = Offset.zero;
      const curve = Curves.ease;
      // Tween是一个无状态(stateless)对象，需要begin和end值。Tween的唯一职责就是定义从输入范围到输出范围的映射。输入范围通常为0.0到1.0，但这不是必须的
      final tween = Tween(begin: begin, end: end).chain(CurveTween(curve: curve));
      return SlideTransition(
        position: animation.drive(tween),
        child: child
      );
    }
    );
}

class Page2 extends StatelessWidget {
  const Page2({ Key? key }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: const Center(
        child: Text('Page2'),
      ),
    );
  }
}