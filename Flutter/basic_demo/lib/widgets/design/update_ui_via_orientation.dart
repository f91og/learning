// update the display of an app when the user rotates the screen from portrait mode to landscape mode.
// OrientationBuilder widget and GridView 

import 'package:flutter/material.dart';

class UpdateUIWhenOrientationChanged extends StatelessWidget {
  const UpdateUIWhenOrientationChanged({ Key? key }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    const appTitle = 'Orientation Demo';
    
    return const MaterialApp(
      title: appTitle,
      home: OrientationList(
        title: appTitle,
      ),
    );
  }
}

class OrientationList extends StatelessWidget {
  const OrientationList({Key? key, required String this.title}) : super(key: key);
  final String title;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(title)),
      body: OrientationBuilder(
         builder: (context, orientation) {
           return GridView.count(
            // Create a grid with 2 columns in portrait mode, or 3 columns in
            // landscape mode.
            crossAxisCount: orientation == Orientation.portrait ? 2 : 3,
            children: List.generate(100, (index) {
              return Center(
                child: Text(
                  'Item $index',
                  style: Theme.of(context).textTheme.headline1,
                ),
              );
            })
          );
         }
      ),
    );
  }
}
