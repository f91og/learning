import 'package:flutter/material.dart';

import 'widgets/animation/animate_physics_simulation.dart';
import 'widgets/animation/animate_route_transition.dart';
import 'widgets/design/add_drawer.dart';
import 'widgets/design/display_snackbar.dart';
import 'widgets/first_app.dart';

// 更改runApp里的widget以显示不同demo的演示效果
// main() => runApp(new FirstApp());
// main() => runApp(MaterialApp(home: AnimateRouteTransition()));
// main() => runApp(MaterialApp(home: PhysicsCardDragDemo()));
// main() => runApp(new DrawerDemo());
main() => runApp(new SnackBarDemo());
