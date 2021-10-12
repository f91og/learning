// 这里写业务逻辑类，把业务逻辑和视图逻辑分开
import 'package:get/get.dart';

class Controller extends GetxController {
  var count = 0.obs; // 如果希望状态被Getx管理则需要在后面加入 .obs
  increment() => count++;
}
