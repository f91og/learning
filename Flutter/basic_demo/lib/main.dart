import 'package:basic_demo/widgets/navigation/animate_across_screen.dart';
import 'package:basic_demo/widgets/navigation/named_routes.dart';
import 'package:flutter/material.dart';

import 'getX/demo.dart';
import 'getX/navigation_demo.dart';
import 'widgets/animation/animate_physics_simulation.dart';
import 'widgets/animation/animate_route_transition.dart';
import 'widgets/animation/fade_widget_in_out.dart';
import 'widgets/design/add_drawer.dart';
import 'widgets/design/display_snackbar.dart';
import 'widgets/design/update_ui_via_orientation.dart';
import 'widgets/design/work_with_tabs.dart';
import 'widgets/first_app.dart';
import 'widgets/forms/validation.dart';
import 'package:get/get.dart';

// 更改runApp里的widget以显示不同demo的演示效果
// main() => runApp(new FirstApp());
// main() => runApp(MaterialApp(home: AnimateRouteTransition()));
// main() => runApp(MaterialApp(home: PhysicsCardDragDemo()));
// main() => runApp(new DrawerDemo());
// main() => runApp(new SnackBarDemo());
// main() => runApp(new FadeWidgetDemo());
// main() => runApp(new UpdateUIWhenOrientationChanged());
// main() => runApp(new TabBarDemo());
// main() => runApp(new ValidationFormDemo());
// main() => runApp(new HeroApp());
// main() => runApp(new GetMaterialApp(home: GetXDemoHome()));
main() => runApp(GetXNavigationDemo());