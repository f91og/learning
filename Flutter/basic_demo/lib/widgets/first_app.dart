import 'package:flutter/material.dart';
import 'package:english_words/english_words.dart';

class FirstApp extends StatelessWidget {
  // 实际上画面渲染的是下面build方法里返回的Widget
  @override
  Widget build(BuildContext context) {
    // final wordPair = new WordPair.random();
    return new MaterialApp(
        debugShowCheckedModeBanner: false,
        title: "Welcome to Flutter",
        theme: new ThemeData(primaryColor: Colors.white),
        // 調用new()构造方法来创建ui组件实例
        // 一个页面的层次结构是 MaterialApp -> [title这样的和MaterialApp里画面显示无关的属性，MaterialApp里具体显示什么（用Scaffold包裹）]
        // home: new Scaffold(
        //   appBar: new AppBar(
        //     title: new Text("Flutter Demo"),
        //   ),
        //   body: new Center(
        //     // child: new Text(wordPair.asPascalCase)
        //     child: new RandomWords(), // 利用有状态的Widget来包裹动态变化的数据（即状态）
        //   ),
        // ),
        home: new RandomWords());
  }
}

class RandomWords extends StatefulWidget {
  @override
  _RandomWordsState createState() =>
      _RandomWordsState(); // 在Dart语言中使用下划线前缀标识符，会强制其变成私有的
}

// 该类持有RandomWords widget的状态，有状态的widget的ui界面指定是在它的状态里指定的
// 感觉这里的边界分离有点不合理，为什么把状态和ui界面封装在一起？估计是为了方便状态和ui的相互通信?
class _RandomWordsState extends State<RandomWords> {
  final _suggestions = <WordPair>[];
  final _biggerFont = const TextStyle(fontSize: 18.0);
  final _saved = new Set<WordPair>();

  void _pushSaved() {
    Navigator.of(context).push(
      new MaterialPageRoute(
        builder: (context) {
          final tiles = _saved.map(
            (pair) {
              return new ListTile(
                title: new Text(
                  pair.asPascalCase,
                  style: _biggerFont,
                ),
              );
            }
          );
          final divided = ListTile.divideTiles(
            context: context,
            tiles: tiles
          ).toList();

          return new Scaffold(
            appBar: new AppBar(title: new Text('Saved Suggestions'),),
            body: new ListView(children: divided,),
          );
        }
      )
    );
  }

  @override
  Widget build(BuildContext context) {
    // final wordPair = new WordPair.random();
    // return new Text(wordPair.asPascalCase);
    return new Scaffold(
      appBar: new AppBar(
        title: new Text("Startup Name Generator"),
        actions: <Widget>[
          new IconButton(onPressed: _pushSaved, icon: new Icon(Icons.list))
        ],
      ),
      body: _buildSuggestions(),
    );
  }

  Widget _buildSuggestions() {
    return new ListView.builder(
      padding: const EdgeInsets.all(16.0),
      // 对于每个建议的单词对都会调用一次itemBuilder，然后将单词对添加到ListTile行中
      itemBuilder: (context, i) {
        // 在奇数行，添加一个分割线widget，来分隔相邻的词对
        if (i.isOdd) return new Divider();

        // 语法 "i ~/ 2" 表示i除以2，但返回值是整形（向下取整），比如i为：1, 2, 3, 4, 5时，结果为0, 1, 1, 2, 2， 这可以计算出ListView中减去分隔线后的实际单词对数量
        final index = i ~/ 2;
        // 如果是建议列表中最后一个单词对，接着再生成10个单词对，然后添加到建议列表
        if (index >= _suggestions.length) {
          _suggestions.addAll(generateWordPairs().take(10));
        }
        return _buildRow(_suggestions[index]);
      },
    );
  }

  Widget _buildRow(WordPair pair) {
    final alreadySaved = _saved.contains(pair);
    return new ListTile(
      title: new Text(
        pair.asPascalCase,
        style: _biggerFont,
      ),
      trailing: new Icon(
        alreadySaved ? Icons.favorite : Icons.favorite_border,
        color: alreadySaved ? Colors.red : null,  // 这里的参数传递和Dart语言特性有关，第一个参数没有指定参数名
      ),
      onTap: () {
        // 调用setState()通知框架状态已经改变
        // 在Flutter的响应式风格的框架中，调用setState() 会为State对象触发build()方法，从而导致对UI的更新
        setState(() {
          if (alreadySaved) {
            _saved.remove(pair);
          } else {
            _saved.add(pair);
          }
        });
      },
    );
  }
}
