import 'package:flutter/material.dart';
import 'package:flutter/physics.dart';

// 1. Set up an animation controller
// 2. Move the widget using gestures
// 3. Animate the widget
// 4. Calculate the velocity to simulate a springing motion

class PhysicsCardDragDemo extends StatelessWidget {
  const PhysicsCardDragDemo({ Key? key }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: const DraggableCard(
        child: FlutterLogo(size: 128,),
      ),
    );
  }
}

/// A draggable card that moves back to [Alignment.center] when it's released.
class DraggableCard extends StatefulWidget {
  const DraggableCard({required this.child, Key? key }) : super(key: key);
  
  final Widget child;

  @override
  _DraggableCardState createState() => _DraggableCardState();
}

// Extending SingleTickerProviderStateMixin allows the state object to be a TickerProvider for the AnimationController.
// TickerProvider: An interface implemented by classes that can vend Ticker objects.
// Tickers can be used by any object that wants to be notified whenever a frame triggers
class _DraggableCardState extends State<DraggableCard> with SingleTickerProviderStateMixin{
  // AnimationController是一个特殊的Animation对象，在屏幕刷新的每一帧，就会生成一个新的值
  late AnimationController _controller;
  late Animation<Alignment> _animation;
  
  /// The alignment of the card as it is dragged or being animated. While the card is being dragged, this value is set to the values computed
  /// in the GestureDetector onPanUpdate callback. If the animation is running, this value is set to the value of the [_animation].
  Alignment _dragAlignment = Alignment.center;

  /// Calculates and runs a [SpringSimulation].
  void _runAnimation(Offset pixelsPerSecond, Size size) {
    _animation = _controller.drive(
      AlignmentTween(
        begin: _dragAlignment,
        end: Alignment.center,
      )
    );
    // Calculate the velocity relative to the unit interval, [0,1], used by the animation controller.
    final unitsPerSecondX = pixelsPerSecond.dx / size.width;
    final unitsPerSecondY = pixelsPerSecond.dy / size.height;
    final unitsPerSecond = Offset(unitsPerSecondX, unitsPerSecondY);
    final unitVelocity = unitsPerSecond.distance;

    const spring = SpringDescription(mass: 30, stiffness: 1, damping: 1);
    final simulation = SpringSimulation(spring, 0, 1, -unitVelocity);

    _controller.animateWith(simulation);   
    _controller.reset();
    _controller.forward();  // 动画的启动
  }

  // 下面两个生命周期方法initState()和dispose()在State和SingleTickerProviderStateMixin都有，那么super是谁？
  // 因为with 修饰的会覆盖 extends 中修饰的同名方法，所以这里的super.initState()因该是SingleTickerProviderStateMixin的initState()?
  @override
  void initState(){
    super.initState();
    // _DraggableCardState要混入SingleTickerProviderStateMixin的原因是
    // 当创建一个AnimationController时，需要传递一个vsync参数，存在vsync时会防止屏幕外动画（动画的UI不在当前屏幕时）消耗不必要的资源
    _controller = AnimationController(vsync: this, duration: const Duration(seconds: 1));
    // 添加帧监听器，在每一帧都会被调用。帧监听器中最常见的行为是改变状态后调用setState()来触发UI重建
    _controller.addListener(() { 
      setState(() { 
        _dragAlignment = _animation.value;  // 通过Animation对象的value属性获取动画的当前状态值，这里的_animation.value是屏幕刷新的每一帧AnimationController生成的
      });
    });
  }
  @override
  void dispose(){
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    var size = MediaQuery.of(context).size; // query current media
    return GestureDetector(
      //手指按下时会触发此回调
      onPanDown: (details) {
        _controller.stop();
      },  
      //手指滑动时会触发此回调
      onPanUpdate: (details) {
        setState(() { // 这个方法会触发UI重新渲染
          _dragAlignment += Alignment(
            details.delta.dx / (size.width / 2),
            details.delta.dy / (size.height / 2),
          );
        });
      },
      onPanEnd: (details){
        _runAnimation(details.velocity.pixelsPerSecond, size);
      },
      child: Align(
        alignment: _dragAlignment,
        child: Card(  // child: the widget below this widget in the tree
          child: widget.child,  // // 这里的widget是这个状态的当前widget实例，这里是flutter框架内自动调用了get方法
        )  
      ),
    );
  }
}