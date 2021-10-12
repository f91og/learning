import 'package:basic_demo/getX/demo_controller.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

// 初识GetX状态管理，路由，sanckbar，Dialog和BottomSheet
class GetXDemoHome extends StatelessWidget {
  const GetXDemoHome({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // 使用Get.put()实例化类，使其对当下的所有子路由可用
    final Controller c = Get.put(Controller());

    return Scaffold(
      // 使用Obx(()=>每当改变计数时，就更新Text()
      appBar: AppBar(
        title: Obx(() => Text("Clicks: ${c.count}")),
      ),

      // 用一个简单的Get.to()即可代替Navigator.push
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            ElevatedButton(
              child: Text("Go to other"),
              onPressed: () => Get.to(Other()),
            ),
            ElevatedButton(
              onPressed: () {
                Get.snackbar("GetX snackbar title",
                    "GetX snackbar content"); // GetX sanckbar
              },
              child: Text("show GetX snackbar"),
            ),
            ElevatedButton(
              onPressed: () => Get.defaultDialog(),
              child: Text("Show Getx Dialog"),
            ),
            ElevatedButton(
              onPressed: () {
                Get.bottomSheet(Container(
                  child: Wrap(
                    children: [
                      ListTile(
                        leading: Icon(Icons.wb_sunny_outlined),
                        title: Text("白天模式"),
                        onTap: () {
                          Get.changeTheme(ThemeData.light());
                        },
                      ),
                      ListTile(
                        leading: Icon(Icons.wb_sunny),
                        title: Text("黑夜模式"),
                        onTap: () {
                          Get.changeTheme(ThemeData.dark());
                        }
                      ),
                    ],
                  ),
                ));
              },
              child: Text("Show GetX Bottom Sheet"),
            )
          ],
        ),
      ),
      floatingActionButton:
          FloatingActionButton(child: Icon(Icons.add), onPressed: c.increment),
    );
  }
}

class Other extends StatelessWidget {
  // 可以让Get找到一个正在被其他页面使用的Controller，并将它返回
  final Controller c = Get.find();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Text("Clicks: ${c.count}"),
      ),
    );
  }
}
