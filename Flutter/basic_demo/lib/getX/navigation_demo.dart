import 'package:flutter/material.dart';
import 'package:get/get.dart';

// 使用命名路由
class GetXNavigationDemo extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      title: "GetX",
      initialRoute: "/",
      defaultTransition: Transition.zoom,
      getPages: [         
        GetPage(name: "/", page: () => GetXNavigationDemo()),
        GetPage(name: "/home", page: () => Home()),
        GetPage(name: "/my", page: () => My(), transition: Transition.rightToLeft)
      ],
      home: NavigationForNamedExample(),
    );
  }
}

class NavigationForNamedExample extends StatelessWidget {
  GlobalKey<NavigatorState> _navKey = GlobalKey();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("GetX NavigationForNamed"),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            ElevatedButton(
              onPressed: () async {
                Get.toNamed("/my");
              },
              child: Text("跳转到首页"))
          ],
        ),
      ),
    );
  }
}


